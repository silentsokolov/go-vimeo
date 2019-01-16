package vimeo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	libraryVersion   = "2.2.3"
	defaultBaseURL   = "https://api.vimeo.com/"
	defaultUserAgent = "go-vimeo/" + libraryVersion

	mediaTypeVersion = "application/vnd.vimeo.*+json;version=3.4"

	headerRateLimit     = "X-RateLimit-Limit"
	headerRateRemaining = "X-RateLimit-Remaining"
	headerRateReset     = "X-RateLimit-Reset"
)

// Client manages communication with Vimeo API.
type Client struct {
	client *http.Client

	BaseURL *url.URL

	UserAgent string

	// Config
	Config *Config

	// Services used for communicating with the API
	Categories      *CategoriesService
	Channels        *ChannelsService
	ContentRatings  *ContentRatingsService
	CreativeCommons *CreativeCommonsService
	Groups          *GroupsService
	Languages       *LanguagesService
	Tags            *TagsService
	Videos          *VideosService
	Users           *UsersService
}

type service struct {
	client *Client
}

// NewClient returns a new Vimeo API client. If a nil httpClient is
// provided, http.DefaultClient will be used. To use API methods which require
// authentication, provide an http.Client that will perform the authentication
// for you (such as that provided by the golang.org/x/oauth2 library).
func NewClient(httpClient *http.Client, config *Config) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if config == nil {
		config = DefaultConfig()
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: defaultUserAgent}
	c.Config = config
	c.Categories = &CategoriesService{client: c}
	c.Channels = &ChannelsService{client: c}
	c.ContentRatings = &ContentRatingsService{client: c}
	c.CreativeCommons = &CreativeCommonsService{client: c}
	c.Groups = &GroupsService{client: c}
	c.Languages = &LanguagesService{client: c}
	c.Tags = &TagsService{client: c}
	c.Videos = &VideosService{client: c}
	c.Users = &UsersService{client: c}
	return c
}

// Client returns the HTTP client configured for this client.
func (c *Client) Client() *http.Client {
	return c.client
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", mediaTypeVersion)

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		return response, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			if err != nil {
				return nil, err
			}
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}

	return response, err
}

type paginator interface {
	GetPage() int
	GetTotal() int
	GetPaging() (string, string, string, string)
}

type paging struct {
	Next  string `json:"next,omitempty"`
	Prev  string `json:"previous,omitempty"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
}

type pagination struct {
	Total  int    `json:"total,omitempty"`
	Page   int    `json:"page,omitempty"`
	Paging paging `json:"paging,omitempty"`
}

// GetPage returns the current page number.
func (p pagination) GetPage() int {
	return p.Page
}

// GetTotal returns the total number of pages.
func (p pagination) GetTotal() int {
	return p.Total
}

// GetPaging returns the data pagination presented as relative references.
// In the following procedure: next, previous, first, last page.
func (p pagination) GetPaging() (string, string, string, string) {
	return p.Paging.Next, p.Paging.Prev, p.Paging.First, p.Paging.Last
}

// Response is a Vimeo response. This wraps the standard http.Response.
// Provides access pagination links.
type Response struct {
	*http.Response
	// Pagination
	Page       int
	TotalPages int
	Total      int
	NextPage   string
	PrevPage   string
	FirstPage  string
	LastPage   string
}

func (r *Response) setPaging(p paginator) {
	r.Page = p.GetPage()
	r.TotalPages = p.GetTotal() // Deprecated
	r.Total = p.GetTotal()
	r.NextPage, r.PrevPage, r.FirstPage, r.LastPage = p.GetPaging()
}

// ErrorResponse is a Vimeo error response. This wraps the standard http.Response.
// Provides access error message returned Vimeo.
type ErrorResponse struct {
	Response *http.Response
	Message  string `json:"error"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Message)
}

// Rate represents the rate limit for the current client.
type Rate struct {
	Limit     int
	Remaining int
	Reset     time.Time
}

// RateLimitError occurs when API response with a rate limit remaining value of 0.
type RateLimitError struct {
	Rate     Rate
	Response *http.Response
	Message  string
}

func (r *RateLimitError) Error() string {
	return fmt.Sprintf("%v %v: %d %v Reset in %v.",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Message, r.Rate.Reset)
}

// parseRate parses the rate related headers.
func parseRate(r *http.Response) Rate {
	var rate Rate

	if reset := r.Header.Get(headerRateReset); reset != "" {
		t, err := time.Parse(time.RFC3339, reset)
		if err == nil {
			rate.Reset = t
		}
	}
	if limit := r.Header.Get(headerRateLimit); limit != "" {
		rate.Limit, _ = strconv.Atoi(limit)
	}
	if remaining := r.Header.Get(headerRateRemaining); remaining != "" {
		rate.Remaining, _ = strconv.Atoi(remaining)
	}

	return rate
}

func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("client_secret")) > 0 {
		params.Set("client_secret", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// CheckResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.
func CheckResponse(r *http.Response) error {
	if code := r.StatusCode; 200 <= code && code <= 299 || code == 308 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)

	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	if r.StatusCode == http.StatusTooManyRequests && r.Header.Get(headerRateRemaining) == "0" {
		return &RateLimitError{
			Rate:     parseRate(r),
			Response: errorResponse.Response,
			Message:  errorResponse.Message,
		}
	}

	return errorResponse
}

// CallOption is an optional argument to an API call.
// A CallOption is something that configures an API call in a way that is not specific to that API: page, filter and etc
type CallOption interface {
	Get() (key, value string)
}

// OptPage is an optional argument to an API call
type OptPage int

// Get return key/value for make query
func (o OptPage) Get() (string, string) {
	return "page", fmt.Sprint(o)
}

// OptPerPage is an optional argument to an API call
type OptPerPage int

// Get return key/value for make query
func (o OptPerPage) Get() (string, string) {
	return "per_page", fmt.Sprint(o)
}

// OptSort is an optional argument to an API call
type OptSort string

// Get key/value for make query
func (o OptSort) Get() (string, string) {
	return "sort", fmt.Sprint(o)
}

// OptDirection is an optional argument to an API call
// All sortable resources accept the direction parameter which must be either asc or desc.
type OptDirection string

// Get key/value for make query
func (o OptDirection) Get() (string, string) {
	return "direction", fmt.Sprint(o)
}

// OptFilter is an optional argument to an API call
type OptFilter string

// Get key/value for make query
func (o OptFilter) Get() (string, string) {
	return "filter", fmt.Sprint(o)
}

// OptFilterEmbeddable is an optional argument to an API call
type OptFilterEmbeddable string

// Get key/value for make query
func (o OptFilterEmbeddable) Get() (string, string) {
	return "filter_embeddable", fmt.Sprint(o)
}

// OptFilterPlayable is an optional argument to an API call
type OptFilterPlayable string

// Get key/value for make query
func (o OptFilterPlayable) Get() (string, string) {
	return "filter_playable", fmt.Sprint(o)
}

// OptQuery is an optional argument to an API call. Search query.
type OptQuery string

// Get key/value for make query
func (o OptQuery) Get() (string, string) {
	return "query", fmt.Sprint(o)
}

// OptFilterContentRating is an optional argument to an API call
// Content filter is a specific type of resource filter, available on all video resources.
// Any videos that do not match one of the provided ratings will be excluded from the list of videos.
// Valid ratings include: language/drugs/violence/nudity/safe/unrated
type OptFilterContentRating []string

// Get key/value for make query
func (o OptFilterContentRating) Get() (string, string) {
	return "filter_content_rating", strings.Join(o, ",")
}

// OptFields is an optional argument to an API call.
// With a simple parameter you can reduce the size of the responses,
// and dramatically increase the performance of your API requests.
type OptFields []string

// Get key/value for make query
func (o OptFields) Get() (string, string) {
	return "fields", strings.Join(o, ",")
}

// OptWeakSearch is an option argument to an API call
// to allow usage of legacy search on the vimeo backend to find private videos
type OptWeakSearch bool

// Get return key/value for make query
func (o OptWeakSearch) Get() (string, string) {
	return "weak_search", fmt.Sprint(o)
}

func addOptions(s string, opts ...CallOption) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs := u.Query()
	for _, o := range opts {
		qs.Set(o.Get())
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
