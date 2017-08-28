package vimeo

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTagsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tags/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	tag, _, err := client.Tags.Get("1")
	if err != nil {
		t.Errorf("Tags.Get returned unexpected error: %v", err)
	}

	want := &Tag{Name: "Test"}
	if !reflect.DeepEqual(tag, want) {
		t.Errorf("Tags.Get returned %+v, want %+v", tag, want)
	}
}

func TestTagsService_ListVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/tags/1/videos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	opt := &ListVideoOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}
	videos, _, err := client.Tags.ListVideo("1", opt)
	if err != nil {
		t.Errorf("Tags.ListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Tags.ListVideo returned %+v, want %+v", videos, want)
	}
}
