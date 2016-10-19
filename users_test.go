package vimeo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestUsersService_Search(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	opt := &ListUserOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}
	users, _, err := client.Users.Search(opt)
	if err != nil {
		t.Errorf("Users.Search returned unexpected error: %v", err)
	}

	want := []*User{{Name: "Test"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Users.Search returned %+v, want %+v", users, want)
	}
}

func TestUsersService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	user, _, err := client.Users.Get("1")
	if err != nil {
		t.Errorf("Users.Get returned unexpected error: %v", err)
	}

	want := &User{Name: "Test"}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Users.Get returned %+v, want %+v", user, want)
	}
}

func TestUsersService_Get_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	user, _, err := client.Users.Get("")
	if err != nil {
		t.Errorf("Users.Get returned unexpected error: %v", err)
	}

	want := &User{Name: "Test"}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Users.Get returned %+v, want %+v", user, want)
	}
}

func TestUsersService_Edit(t *testing.T) {
	setup()
	defer teardown()

	input := &UserRequest{
		Name: "name",
	}

	mux.HandleFunc("/users/1", func(w http.ResponseWriter, r *http.Request) {
		v := &UserRequest{}
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PATCH")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Users.Edit body is %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"name": "name"}`)
	})

	user, _, err := client.Users.Edit("1", input)
	if err != nil {
		t.Errorf("Users.Edit returned unexpected error: %v", err)
	}

	want := &User{Name: "name"}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Users.Edit returned %+v, want %+v", user, want)
	}
}

func TestUsersService_Edit_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	input := &UserRequest{
		Name: "name",
	}

	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		v := &UserRequest{}
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PATCH")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Users.Edit body is %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"name": "name"}`)
	})

	user, _, err := client.Users.Edit("", input)
	if err != nil {
		t.Errorf("Users.Edit returned unexpected error: %v", err)
	}

	want := &User{Name: "name"}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Users.Edit returned %+v, want %+v", user, want)
	}
}
