package vimeo

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// client configured to use test server
	client = NewClient(nil, nil)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

// teardown closes the test HTTP server.
func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

type values map[string]string

func testFormURLValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Add(k, v)
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func TestNewClient(t *testing.T) {
	c := NewClient(nil, nil)
	if baseURL := c.BaseURL.String(); baseURL != defaultBaseURL {
		t.Errorf("NewClient BaseURL is %v, want %v", baseURL, defaultBaseURL)
	}

	if userAgent := c.UserAgent; userAgent != defaultUserAgent {
		t.Errorf("NewClient UserAgent is %+v, want %+v", userAgent, defaultUserAgent)
	}

	if client := c.Client(); client != http.DefaultClient {
		t.Errorf("NewClient Client is %+v, want %+v", client, http.DefaultClient)
	}

	testClient := new(http.Client)
	c = NewClient(testClient, nil)
	if client := c.Client(); client != testClient {
		t.Errorf("NewClient Client is %+v, want %+v", client, testClient)
	}
}

func TestNewRequest(t *testing.T) {
	c := NewClient(nil, nil)

	type T struct {
		Field string
	}

	testURL := defaultBaseURL + "test"
	testBody := &T{Field: "Value"}
	testBodyAsStr := "{\"Field\":\"Value\"}\n"

	req, _ := c.NewRequest("GET", "/test", testBody)

	if url := req.URL.String(); url != testURL {
		t.Errorf("NewRequest URL is %v, want %v", url, testURL)
	}

	body, _ := ioutil.ReadAll(req.Body)
	if body := string(body); body != testBodyAsStr {
		t.Errorf("NewRequest Body is %v, want %v", body, testBodyAsStr)
	}

	if headerUA := req.Header.Get("User-Agent"); headerUA != c.UserAgent {
		t.Errorf("NewRequest header User-Agent is %v, want %v", headerUA, c.UserAgent)
	}

	if headerAccept := req.Header.Get("Accept"); headerAccept != mediaTypeVersion {
		t.Errorf("NewRequest header Accept is %v, want %v", headerAccept, mediaTypeVersion)
	}
}

func TestNewRequest_badURL(t *testing.T) {
	c := NewClient(nil, nil)
	_, err := c.NewRequest("GET", ":", nil)

	if err == nil {
		t.Errorf("Expected error to be returned")
	}

	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func TestNewRequest_emptyBody(t *testing.T) {
	c := NewClient(nil, nil)
	req, err := c.NewRequest("GET", "/", nil)

	if err != nil {
		t.Errorf("NewRequest returned unexpected error %v", err)
	}

	if req.Body != nil {
		t.Fatalf("Constructed request contains a non-nil Body")
	}
}

func TestNewRequest_invalidJSON(t *testing.T) {
	c := NewClient(nil, nil)

	type T struct {
		F map[interface{}]interface{}
	}

	_, err := c.NewRequest("GET", "/", &T{})

	if err == nil {
		t.Errorf("NewRequest expected error")
	}

	if err, ok := err.(*json.UnsupportedTypeError); !ok {
		t.Errorf("Expected a JSON error; got %#v.", err)
	}
}

func TestNewRequest_emptyUserAgent(t *testing.T) {
	c := NewClient(nil, nil)

	c.UserAgent = ""

	req, err := c.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("NewRequest returned unexpected error: %v", err)
	}

	if _, ok := req.Header["User-Agent"]; ok {
		t.Fatal("Constructed request contains unexpected User-Agent header")
	}
}

func TestDo(t *testing.T) {
	setup()
	defer teardown()

	type T struct {
		F string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method is %v, want %v", r.Method, m)
		}
		fmt.Fprint(w, `{"F":"v"}`)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	body := new(T)

	client.Do(req, body)

	want := &T{"v"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body is %v, want %v", body, want)
	}
}

func TestDo_httpError(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad request", http.StatusBadRequest)
	})

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, nil)

	if err == nil {
		t.Error("Expected HTTP error.")
	}
}

func TestDo_noContent(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	var body json.RawMessage

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, &body)

	if err != nil {
		t.Fatalf("Do returned unexpected error: %v", err)
	}
}

func TestDo_ioWriter(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method is %v, want %v", r.Method, m)
		}
		fmt.Fprint(w, `{"F":"v"}`)
	})

	var b bytes.Buffer
	body := bufio.NewWriter(&b)

	req, _ := client.NewRequest("GET", "/", nil)
	_, err := client.Do(req, body)

	if err != nil {
		t.Errorf("Do returned unexpected error: %v", err)
	}
}

func TestPagination_GetPage(t *testing.T) {
	p := pagination{Page: 1}
	if page := p.GetPage(); page != 1 {
		t.Errorf("pagination GetPage is %v, want %v", page, 1)
	}
}

func TestPagination_GetTotal(t *testing.T) {
	p := pagination{Total: 1}
	if total := p.GetTotal(); total != 1 {
		t.Errorf("pagination GetTotal is %v, want %v", total, 1)
	}
}

func TestPagination_GetPaging(t *testing.T) {
	p := pagination{
		Paging: paging{
			Next:  "/page=3",
			Prev:  "/page=1",
			First: "/page=1",
			Last:  "/page=10",
		}}

	next, prev, first, last := p.GetPaging()

	if next != p.Paging.Next {
		t.Errorf("pagination GetPaging next is %v, want %v", next, p.Paging.Next)
	}

	if prev != p.Paging.Prev {
		t.Errorf("pagination GetPaging prev is %v, want %v", prev, p.Paging.Prev)
	}

	if first != p.Paging.First {
		t.Errorf("pagination GetPaging first is %v, want %v", first, p.Paging.First)
	}

	if last != p.Paging.Last {
		t.Errorf("pagination GetPaging last is %v, want %v", last, p.Paging.Last)
	}
}

func TestResponse_setPaging(t *testing.T) {
	p := pagination{
		Page:  1,
		Total: 10,
		Paging: paging{
			Next:  "/page=3",
			Prev:  "/page=1",
			First: "/page=1",
			Last:  "/page=10",
		},
	}
	resp := Response{}
	resp.setPaging(p)

	if resp.Page != p.Page {
		t.Errorf("Response OptPage is %v, want %v", resp.Page, p.Page)
	}

	if resp.Total != p.Total {
		t.Errorf("Response Total is %v, want %v", resp.Total, p.Total)
	}

	if resp.NextPage != p.Paging.Next {
		t.Errorf("Response NextPage is %v, want %v", resp.NextPage, p.Paging.Next)
	}

	if resp.PrevPage != p.Paging.Prev {
		t.Errorf("Response PrevPage is %v, want %v", resp.PrevPage, p.Paging.Prev)
	}

	if resp.FirstPage != p.Paging.First {
		t.Errorf("Response FirstPage is %v, want %v", resp.FirstPage, p.Paging.First)
	}

	if resp.LastPage != p.Paging.Last {
		t.Errorf("Response LastPage is %v, want %v", resp.LastPage, p.Paging.Last)
	}
}

func TestErrorResponse_Error(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{Method: "GET"},
		StatusCode: http.StatusBadRequest,
	}

	errResponse := ErrorResponse{Response: res}

	if errResponse.Error() == "" {
		t.Errorf("ErrorResponse.Error returned empty error.")
	}
}

func TestSanitizeURL(t *testing.T) {
	URLWithoutSecret, _ := url.Parse("/method?a=b")
	wantURLWithoutSecret, _ := url.Parse("/method?a=b")
	if u := sanitizeURL(URLWithoutSecret); !reflect.DeepEqual(URLWithoutSecret, wantURLWithoutSecret) {
		t.Errorf("sanitizeURL url is %v, want %v", u, wantURLWithoutSecret)
	}

	URLWithSecret, _ := url.Parse("/method?a=b&client_secret=SECRET")
	wantURLWithSecret, _ := url.Parse("/method?a=b&client_secret=REDACTED")
	if u := sanitizeURL(URLWithSecret); !reflect.DeepEqual(URLWithSecret, wantURLWithSecret) {
		t.Errorf("sanitizeURL url is %v, want %v", u, wantURLWithSecret)
	}
}

func TestCheckError_statusOK(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusOK,
	}

	err := CheckResponse(res)

	if err != nil {
		t.Fatalf("CheckResponse returned unexpected error: %v", err)
	}
}

func TestCheckError_statusFail(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusBadRequest,
		Body:       ioutil.NopCloser(strings.NewReader(`{"error": "Invalid type for field [field]"}`)),
	}

	wantError := &ErrorResponse{
		Response: res,
		Message:  "Invalid type for field [field]",
	}

	err := CheckResponse(res).(*ErrorResponse)

	if err == nil {
		t.Error("Expected error response.")
	}

	if !reflect.DeepEqual(err, wantError) {
		t.Errorf("CheckResponse unexpected err %v, want %v", err, wantError)
	}
}

func TestCheckError_rateLimit(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusTooManyRequests,
		Header:     http.Header{},
		Body:       ioutil.NopCloser(strings.NewReader(`{"error": "Invalid type for field [field]"}`)),
	}

	res.Header.Set(headerRateLimit, "1000")
	res.Header.Set(headerRateRemaining, "0")
	res.Header.Set(headerRateReset, "2017-09-14T09:47:00+00:00")

	err := CheckResponse(res)
	rateLimitErr, ok := err.(*RateLimitError)
	if !ok {
		t.Fatalf("Expected a *RateLimitError error; got %#v.", err)
	}
	if got, want := rateLimitErr.Rate.Limit, 1000; got != want {
		t.Errorf("rateLimitErr rate limit = %v, want %v", got, want)
	}
	if got, want := rateLimitErr.Rate.Remaining, 0; got != want {
		t.Errorf("rateLimitErr rate remaining = %v, want %v", got, want)
	}
	reset := time.Date(2017, 9, 14, 9, 47, 0, 0, time.UTC)
	if rateLimitErr.Rate.Reset.UTC() != reset {
		t.Errorf("rateLimitErr rate reset = %v, want %v", rateLimitErr.Rate.Reset.UTC(), reset)
	}
}

func TestAddOptions(t *testing.T) {
	opURL, err := addOptions("api", OptPage(2), OptFilter("feature"))
	if err != nil {
		t.Errorf("addOptions returned unexpected error: %v", err)
	}

	if opURL != "api?filter=feature&page=2" {
		t.Errorf("addOptions returned url: %v, get %v", opURL, "api?filter=feature&page=2")
	}
}

func TestPageOption(t *testing.T) {
	opt := OptPage(10)
	k, v := opt.Get()

	if k != "page" {
		t.Errorf("OptPage returned key: %v, get %v", k, "page")
	}

	if v != "10" {
		t.Errorf("OptPage returned value: %v, get %v", v, "10")
	}
}

func TestPerPageOption(t *testing.T) {
	opt := OptPerPage(10)
	k, v := opt.Get()

	if k != "per_page" {
		t.Errorf("OptPerPage returned key: %v, get %v", k, "per_page")
	}

	if v != "10" {
		t.Errorf("OptPerPage returned value: %v, get %v", v, "10")
	}
}

func TestSortOption(t *testing.T) {
	opt := OptSort("name")
	k, v := opt.Get()

	if k != "sort" {
		t.Errorf("OptSort returned key: %v, get %v", k, "sort")
	}

	if v != "name" {
		t.Errorf("OptSort returned value: %v, get %v", v, "name")
	}
}

func TestDirectionOption(t *testing.T) {
	opt := OptDirection("name")
	k, v := opt.Get()

	if k != "direction" {
		t.Errorf("OptDirection returned key: %v, get %v", k, "sort")
	}

	if v != "name" {
		t.Errorf("OptDirection returned value: %v, get %v", v, "name")
	}
}

func TestFilterOption(t *testing.T) {
	opt := OptFilter("name")
	k, v := opt.Get()

	if k != "filter" {
		t.Errorf("OptFilter returned key: %v, get %v", k, "filter")
	}

	if v != "name" {
		t.Errorf("OptFilter returned value: %v, get %v", v, "name")
	}
}

func TestQueryOption(t *testing.T) {
	opt := OptQuery("name")
	k, v := opt.Get()

	if k != "query" {
		t.Errorf("OptQuery returned key: %v, get %v", k, "query")
	}

	if v != "name" {
		t.Errorf("OptQuery returned value: %v, get %v", v, "name")
	}
}

func TestFilterPlayableOption(t *testing.T) {
	opt := OptFilterPlayable("name")
	k, v := opt.Get()

	if k != "filter_playable" {
		t.Errorf("OptFilterPlayable returned key: %v, get %v", k, "filter_playable")
	}

	if v != "name" {
		t.Errorf("OptFilterPlayable returned value: %v, get %v", v, "name")
	}
}

func TestFilterEmbeddableOption(t *testing.T) {
	opt := OptFilterEmbeddable("name")
	k, v := opt.Get()

	if k != "filter_embeddable" {
		t.Errorf("OptFilterEmbeddable returned key: %v, get %v", k, "filter_embeddable")
	}

	if v != "name" {
		t.Errorf("OptFilterEmbeddable returned value: %v, get %v", v, "name")
	}
}

func TestFilterContentRatingOption(t *testing.T) {
	opt := OptFilterContentRating([]string{"a", "b", "c"})
	k, v := opt.Get()

	if k != "filter_content_rating" {
		t.Errorf("OptFilterContentRating returned key: %v, get %v", k, "filter_content_rating")
	}

	if v != "a,b,c" {
		t.Errorf("OptFilterContentRating returned value: %v, get %v", v, "a,b,c")
	}
}

func TestFieldsOption(t *testing.T) {
	opt := OptFields([]string{"a", "b", "c"})
	k, v := opt.Get()

	if k != "fields" {
		t.Errorf("OptFields returned key: %v, get %v", k, "fields")
	}

	if v != "a,b,c" {
		t.Errorf("OptFields returned value: %v, get %v", v, "a,b,c")
	}
}
