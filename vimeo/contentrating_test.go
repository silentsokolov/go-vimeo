package vimeo

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestContentRatingsService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/contentratings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormURLValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	ratings, _, err := client.ContentRatings.List(Page(1), PerPage(2))
	if err != nil {
		t.Errorf("ContentRatings.List returned unexpected error: %v", err)
	}

	want := []*ContentRating{{Name: "Test"}}
	if !reflect.DeepEqual(ratings, want) {
		t.Errorf("ContentRatings.List returned %+v, want %+v", ratings, want)
	}
}
