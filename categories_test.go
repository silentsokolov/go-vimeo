package vimeo

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCategoriesService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	opt := &ListCategoryOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}
	categories, _, err := client.Categories.List(opt)
	if err != nil {
		t.Errorf("Categories.List returned unexpected error: %v", err)
	}

	want := []*Category{{Name: "Test"}}
	if !reflect.DeepEqual(categories, want) {
		t.Errorf("Categories.List returned %+v, want %+v", categories, want)
	}
}

func TestCategoriesService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/categories/cat", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	category, _, err := client.Categories.Get("cat")
	if err != nil {
		t.Errorf("Categories.Get returned unexpected error: %v", err)
	}

	want := &Category{Name: "Test"}
	if !reflect.DeepEqual(category, want) {
		t.Errorf("Categories.Get returned %+v, want %+v", category, want)
	}
}

func TestCategoriesService_ListChannel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/categories/cat/channels", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	opt := &ListChannelOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}
	channels, _, err := client.Categories.ListChannel("cat", opt)
	if err != nil {
		t.Errorf("Categories.ListChannel returned unexpected error: %v", err)
	}

	want := []*Channel{{Name: "Test"}}
	if !reflect.DeepEqual(channels, want) {
		t.Errorf("Categories.ListChannel returned %+v, want %+v", channels, want)
	}
}

func TestCategoriesService_ListGroup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/categories/cat/groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	opt := &ListGroupOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}
	groups, _, err := client.Categories.ListGroup("cat", opt)
	if err != nil {
		t.Errorf("Categories.ListGroup returned unexpected error: %v", err)
	}

	want := []*Group{{Name: "Test"}}
	if !reflect.DeepEqual(groups, want) {
		t.Errorf("Categories.ListGroup returned %+v, want %+v", groups, want)
	}
}

func TestCategoriesService_ListVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/categories/cat/videos", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Categories.ListVideo("cat", opt)
	if err != nil {
		t.Errorf("Categories.ListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Categories.ListVideo returned %+v, want %+v", videos, want)
	}
}

func TestVideo_GetID(t *testing.T) {
	v := &Video{Name: "Test", URI: "/videos/1"}

	if id := v.GetID(); id != 1 {
		t.Errorf("Video.GetID returned %+v, want %+v", id, 1)
	}
}

func TestCategoriesService_GetVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/categories/cat/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Categories.GetVideo("cat", 1)
	if err != nil {
		t.Errorf("Categories.GetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Categories.GetVideo returned %+v, want %+v", video, want)
	}
}
