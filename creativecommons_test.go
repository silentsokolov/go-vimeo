package vimeo

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCreativeCommonsService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/creativecommons", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	opt := &ListCreativeCommonOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}

	commons, _, err := client.CreativeCommons.List(opt)
	if err != nil {
		t.Errorf("CreativeCommons.List returned unexpected error: %v", err)
	}

	want := []*CreativeCommon{{Name: "Test"}}
	if !reflect.DeepEqual(commons, want) {
		t.Errorf("CreativeCommons.List returned %+v, want %+v", commons, want)
	}
}
