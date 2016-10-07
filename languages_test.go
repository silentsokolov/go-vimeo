package vimeo

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestLanguagesService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/languages", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	opt := &ListLanguageOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}

	languages, _, err := client.Languages.List(opt)
	if err != nil {
		t.Errorf("Languages.List returned unexpected error: %v", err)
	}

	want := []*Language{{Name: "Test"}}
	if !reflect.DeepEqual(languages, want) {
		t.Errorf("Languages.List returned %+v, want %+v", languages, want)
	}
}
