package vimeo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGroup_GetID(t *testing.T) {
	v := &Group{Name: "Test", URI: "/groups/1"}

	if id := v.GetID(); id != "1" {
		t.Errorf("Group.GetID returned %+v, want %+v", id, "1")
	}
}

func TestGroupsService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormURLValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	groups, _, err := client.Groups.List(Page(1), PerPage(2))
	if err != nil {
		t.Errorf("Groups.List returned unexpected error: %v", err)
	}

	want := []*Group{{Name: "Test"}}
	if !reflect.DeepEqual(groups, want) {
		t.Errorf("Groups.List returned %+v, want %+v", groups, want)
	}
}

func TestGroupsService_Create(t *testing.T) {
	setup()
	defer teardown()

	input := &GroupRequest{
		Name:        "name",
		Description: "desc",
	}

	mux.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		v := &GroupRequest{}
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Groups.Create body is %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"name": "name"}`)
	})

	group, _, err := client.Groups.Create(input)
	if err != nil {
		t.Errorf("Groups.Create returned unexpected error: %v", err)
	}

	want := &Group{Name: "name"}
	if !reflect.DeepEqual(group, want) {
		t.Errorf("Groups.Create returned %+v, want %+v", group, want)
	}
}

func TestGroupsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormURLValues(t, r, values{
			"fields": "name",
		})
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	group, _, err := client.Groups.Get("1", Fields([]string{"name"}))
	if err != nil {
		t.Errorf("Groups.Get returned unexpected error: %v", err)
	}

	want := &Group{Name: "Test"}
	if !reflect.DeepEqual(group, want) {
		t.Errorf("Groups.Get returned %+v, want %+v", group, want)
	}
}

func TestGroupsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Groups.Delete("1")
	if err != nil {
		t.Errorf("Groups.Delete returned unexpected error: %v", err)
	}
}

func TestGroupsService_ListUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups/1/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormURLValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	users, _, err := client.Groups.ListUser("1", Page(1), PerPage(2))
	if err != nil {
		t.Errorf("Groups.ListUser returned unexpected error: %v", err)
	}

	want := []*User{{Name: "Test"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Groups.ListUser returned %+v, want %+v", users, want)
	}
}

func TestGroupsService_ListVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups/1/videos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormURLValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	videos, _, err := client.Groups.ListVideo("1", Page(1), PerPage(2))
	if err != nil {
		t.Errorf("Groups.ListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Groups.ListVideo returned %+v, want %+v", videos, want)
	}
}

func TestGroupsService_GetVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups/gr/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Groups.GetVideo("gr", 1)
	if err != nil {
		t.Errorf("Groups.GetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Groups.GetVideo returned %+v, want %+v", video, want)
	}
}

func TestGroupsService_DeleteVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/groups/gr/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Groups.DeleteVideo("gr", 1)
	if err != nil {
		t.Errorf("Groups.DeleteVideo returned unexpected error: %v", err)
	}
}
