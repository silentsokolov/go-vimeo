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

func TestUsersService_ListAlbum(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/albums", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	opt := &ListAlbumOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}
	albums, _, err := client.Users.ListAlbum("1", opt)
	if err != nil {
		t.Errorf("Users.ListAlbum returned unexpected error: %v", err)
	}

	want := []*Album{{Name: "Test"}}
	if !reflect.DeepEqual(albums, want) {
		t.Errorf("Users.ListAlbum returned %+v, want %+v", albums, want)
	}
}

func TestUsersService_ListAlbum_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/albums", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	opt := &ListAlbumOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}
	albums, _, err := client.Users.ListAlbum("", opt)
	if err != nil {
		t.Errorf("Users.ListAlbum returned unexpected error: %v", err)
	}

	want := []*Album{{Name: "Test"}}
	if !reflect.DeepEqual(albums, want) {
		t.Errorf("Users.ListAlbum returned %+v, want %+v", albums, want)
	}
}

func TestUsersService_CreateAlbum(t *testing.T) {
	setup()
	defer teardown()

	input := &AlbumRequest{
		Name:        "name",
		Description: "desc",
	}

	mux.HandleFunc("/users/1/albums", func(w http.ResponseWriter, r *http.Request) {
		v := &AlbumRequest{}
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Users.CreateAlbum body is %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"name": "name"}`)
	})

	album, _, err := client.Users.CreateAlbum("1", input)
	if err != nil {
		t.Errorf("Users.CreateAlbum returned unexpected error: %v", err)
	}

	want := &Album{Name: "name"}
	if !reflect.DeepEqual(album, want) {
		t.Errorf("Users.CreateAlbum returned %+v, want %+v", album, want)
	}
}

func TestUsersService_CreateAlbum_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	input := &AlbumRequest{
		Name:        "name",
		Description: "desc",
	}

	mux.HandleFunc("/me/albums", func(w http.ResponseWriter, r *http.Request) {
		v := &AlbumRequest{}
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Users.CreateAlbum body is %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"name": "name"}`)
	})

	album, _, err := client.Users.CreateAlbum("", input)
	if err != nil {
		t.Errorf("Users.CreateAlbum returned unexpected error: %v", err)
	}

	want := &Album{Name: "name"}
	if !reflect.DeepEqual(album, want) {
		t.Errorf("Users.CreateAlbum returned %+v, want %+v", album, want)
	}
}

func TestUsersService_GetAlbum(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/albums/a", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	album, _, err := client.Users.GetAlbum("1", "a")
	if err != nil {
		t.Errorf("Users.GetAlbum returned unexpected error: %v", err)
	}

	want := &Album{Name: "Test"}
	if !reflect.DeepEqual(album, want) {
		t.Errorf("Users.GetAlbum returned %+v, want %+v", album, want)
	}
}

func TestUsersService_GetAlbum_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/albums/a", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	album, _, err := client.Users.GetAlbum("", "a")
	if err != nil {
		t.Errorf("Users.GetAlbum returned unexpected error: %v", err)
	}

	want := &Album{Name: "Test"}
	if !reflect.DeepEqual(album, want) {
		t.Errorf("Users.GetAlbum returned %+v, want %+v", album, want)
	}
}

func TestUsersService_EditAlbum(t *testing.T) {
	setup()
	defer teardown()

	input := &AlbumRequest{
		Name:        "name",
		Description: "desc",
		Privacy:     "anybody",
	}

	mux.HandleFunc("/users/1/albums/a", func(w http.ResponseWriter, r *http.Request) {
		v := &AlbumRequest{}
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PATCH")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Users.EditAlbum body is %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"name": "name"}`)
	})

	album, _, err := client.Users.EditAlbum("1", "a", input)
	if err != nil {
		t.Errorf("Users.Edit returned unexpected error: %v", err)
	}

	want := &Album{Name: "name"}
	if !reflect.DeepEqual(album, want) {
		t.Errorf("Users.EditAlbum returned %+v, want %+v", album, want)
	}
}

func TestUsersService_EditAlbum_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	input := &AlbumRequest{
		Name:        "name",
		Description: "desc",
		Privacy:     "anybody",
	}

	mux.HandleFunc("/me/albums/a", func(w http.ResponseWriter, r *http.Request) {
		v := &AlbumRequest{}
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PATCH")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Users.EditAlbum body is %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"name": "name"}`)
	})

	album, _, err := client.Users.EditAlbum("", "a", input)
	if err != nil {
		t.Errorf("Users.Edit returned unexpected error: %v", err)
	}

	want := &Album{Name: "name"}
	if !reflect.DeepEqual(album, want) {
		t.Errorf("Users.EditAlbum returned %+v, want %+v", album, want)
	}
}

func TestUsersService_DeleteAlbum(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/albums/a", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.DeleteAlbum("1", "a")
	if err != nil {
		t.Errorf("Users.DeleteAlbum returned unexpected error: %v", err)
	}
}

func TestUsersService_DeleteAlbum_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/albums/a", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.DeleteAlbum("", "a")
	if err != nil {
		t.Errorf("Users.DeleteAlbum returned unexpected error: %v", err)
	}
}

func TestUsersService_AlbumListVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/albums/a/videos", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.AlbumListVideo("1", "a", opt)
	if err != nil {
		t.Errorf("Users.AlbumListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.AlbumListVideo returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_AlbumListVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/albums/a/videos", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.AlbumListVideo("", "a", opt)
	if err != nil {
		t.Errorf("Users.AlbumListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.AlbumListVideo returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_AlbumGetVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/albums/a/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Users.AlbumGetVideo("1", "a", 1)
	if err != nil {
		t.Errorf("Users.AlbumGetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Users.AlbumGetVideo returned %+v, want %+v", video, want)
	}
}

func TestUsersService_AlbumGetVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/albums/a/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Users.AlbumGetVideo("", "a", 1)
	if err != nil {
		t.Errorf("Users.AlbumGetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Users.AlbumGetVideo returned %+v, want %+v", video, want)
	}
}

func TestUsersService_AlbumDeleteVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/albums/a/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.AlbumDeleteVideo("1", "a", 1)
	if err != nil {
		t.Errorf("Users.AlbumDeleteVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_AlbumDeleteVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/albums/a/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.AlbumDeleteVideo("", "a", 1)
	if err != nil {
		t.Errorf("Users.AlbumDeleteVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_ListAppearance(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/appearances", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.ListAppearance("1", opt)
	if err != nil {
		t.Errorf("Users.ListAppearance returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.ListAppearance returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_ListAppearance_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/appearances", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.ListAppearance("", opt)
	if err != nil {
		t.Errorf("Users.ListAppearance returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.ListAppearance returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_ListCategory(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/categories", func(w http.ResponseWriter, r *http.Request) {
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
	categories, _, err := client.Users.ListCategory("1", opt)
	if err != nil {
		t.Errorf("Users.ListCategory returned unexpected error: %v", err)
	}

	want := []*Category{{Name: "Test"}}
	if !reflect.DeepEqual(categories, want) {
		t.Errorf("Users.ListCategory returned %+v, want %+v", categories, want)
	}
}

func TestUsersService_ListCategory_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/categories", func(w http.ResponseWriter, r *http.Request) {
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
	categories, _, err := client.Users.ListCategory("", opt)
	if err != nil {
		t.Errorf("Users.ListCategory returned unexpected error: %v", err)
	}

	want := []*Category{{Name: "Test"}}
	if !reflect.DeepEqual(categories, want) {
		t.Errorf("Users.ListCategory returned %+v, want %+v", categories, want)
	}
}

func TestUsersService_SubscribeCategory(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/categories/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.SubscribeCategory("1", "1")
	if err != nil {
		t.Errorf("Users.SubscribeCategory returned unexpected error: %v", err)
	}
}

func TestUsersService_SubscribeCategory_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/categories/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.SubscribeCategory("", "1")
	if err != nil {
		t.Errorf("Users.SubscribeCategory returned unexpected error: %v", err)
	}
}

func TestUsersService_UnsubscribeCategory(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/categories/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.UnsubscribeCategory("1", "1")
	if err != nil {
		t.Errorf("Users.UnsubscribeCategory returned unexpected error: %v", err)
	}
}

func TestUsersService_UnsubscribeCategory_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/categories/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.UnsubscribeCategory("", "1")
	if err != nil {
		t.Errorf("Users.UnsubscribeCategory returned unexpected error: %v", err)
	}
}

func TestUsersService_ListChannel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/channels", func(w http.ResponseWriter, r *http.Request) {
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
	channels, _, err := client.Users.ListChannel("1", opt)
	if err != nil {
		t.Errorf("Users.ListChannel returned unexpected error: %v", err)
	}

	want := []*Channel{{Name: "Test"}}
	if !reflect.DeepEqual(channels, want) {
		t.Errorf("Users.ListChannel returned %+v, want %+v", channels, want)
	}
}

func TestUsersService_ListChannel_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/channels", func(w http.ResponseWriter, r *http.Request) {
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
	channels, _, err := client.Users.ListChannel("", opt)
	if err != nil {
		t.Errorf("Users.ListChannel returned unexpected error: %v", err)
	}

	want := []*Channel{{Name: "Test"}}
	if !reflect.DeepEqual(channels, want) {
		t.Errorf("Users.ListChannel returned %+v, want %+v", channels, want)
	}
}

func TestUsersService_SubscribeChannel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/channels/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.SubscribeChannel("1", "1")
	if err != nil {
		t.Errorf("Users.SubscribeChannel returned unexpected error: %v", err)
	}
}

func TestUsersService_SubscribeChannel_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/channels/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.SubscribeChannel("", "1")
	if err != nil {
		t.Errorf("Users.SubscribeChannel returned unexpected error: %v", err)
	}
}

func TestUsersService_UnsubscribeChannel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/channels/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.UnsubscribeChannel("1", "1")
	if err != nil {
		t.Errorf("Users.UnsubscribeChannel returned unexpected error: %v", err)
	}
}

func TestUsersService_UnsubscribeChannel_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/channels/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.UnsubscribeChannel("", "1")
	if err != nil {
		t.Errorf("Users.UnsubscribeChannel returned unexpected error: %v", err)
	}
}
