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

func TestUsersService_AlbumAddVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/albums/a/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.AlbumAddVideo("1", "a", 1)
	if err != nil {
		t.Errorf("Users.AlbumAddVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_AlbumAddVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/albums/a/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.AlbumAddVideo("", "a", 1)
	if err != nil {
		t.Errorf("Users.AlbumAddVideo returned unexpected error: %v", err)
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

func TestUsersService_Feed(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/feed", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"uri": "/1"}]}`)
	})

	opt := &ListFeedOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}
	feed, _, err := client.Users.Feed("1", opt)
	if err != nil {
		t.Errorf("Users.Feed returned unexpected error: %v", err)
	}

	want := []*Feed{{URI: "/1"}}
	if !reflect.DeepEqual(feed, want) {
		t.Errorf("Users.Feed returned %+v, want %+v", feed, want)
	}
}

func TestUsersService_Feed_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/feed", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"uri": "/1"}]}`)
	})

	opt := &ListFeedOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}
	feed, _, err := client.Users.Feed("", opt)
	if err != nil {
		t.Errorf("Users.Feed returned unexpected error: %v", err)
	}

	want := []*Feed{{URI: "/1"}}
	if !reflect.DeepEqual(feed, want) {
		t.Errorf("Users.Feed returned %+v, want %+v", feed, want)
	}
}

func TestUsersService_ListFollower(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/followers", func(w http.ResponseWriter, r *http.Request) {
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
	users, _, err := client.Users.ListFollower("1", opt)
	if err != nil {
		t.Errorf("Users.ListFollower returned unexpected error: %v", err)
	}

	want := []*User{{Name: "Test"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Users.ListFollower returned %+v, want %+v", users, want)
	}
}

func TestUsersService_ListFollower_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/followers", func(w http.ResponseWriter, r *http.Request) {
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
	users, _, err := client.Users.ListFollower("", opt)
	if err != nil {
		t.Errorf("Users.ListFollower returned unexpected error: %v", err)
	}

	want := []*User{{Name: "Test"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Users.ListFollower returned %+v, want %+v", users, want)
	}
}

func TestUsersService_ListFollowed(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/following", func(w http.ResponseWriter, r *http.Request) {
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
	users, _, err := client.Users.ListFollowed("1", opt)
	if err != nil {
		t.Errorf("Users.ListFollowed returned unexpected error: %v", err)
	}

	want := []*User{{Name: "Test"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Users.ListFollowed returned %+v, want %+v", users, want)
	}
}

func TestUsersService_ListFollowed_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/following", func(w http.ResponseWriter, r *http.Request) {
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
	users, _, err := client.Users.ListFollowed("", opt)
	if err != nil {
		t.Errorf("Users.ListFollowed returned unexpected error: %v", err)
	}

	want := []*User{{Name: "Test"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Users.ListFollowed returned %+v, want %+v", users, want)
	}
}

func TestUsersService_FollowUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/following/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.FollowUser("1", "2")
	if err != nil {
		t.Errorf("Users.FollowUser returned unexpected error: %v", err)
	}
}

func TestUsersService_FollowUser_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/following/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.FollowUser("", "2")
	if err != nil {
		t.Errorf("Users.FollowUser returned unexpected error: %v", err)
	}
}

func TestUsersService_UnfollowUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/following/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.UnfollowUser("1", "2")
	if err != nil {
		t.Errorf("Users.UnfollowUser returned unexpected error: %v", err)
	}
}

func TestUsersService_UnfollowUser_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/following/2", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.UnfollowUser("", "2")
	if err != nil {
		t.Errorf("Users.UnfollowUser returned unexpected error: %v", err)
	}
}

func TestUsersService_ListGroup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/groups", func(w http.ResponseWriter, r *http.Request) {
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
	groups, _, err := client.Users.ListGroup("1", opt)
	if err != nil {
		t.Errorf("Users.ListGroup returned unexpected error: %v", err)
	}

	want := []*Group{{Name: "Test"}}
	if !reflect.DeepEqual(groups, want) {
		t.Errorf("Users.ListGroup returned %+v, want %+v", groups, want)
	}
}

func TestUsersService_ListGroup_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/groups", func(w http.ResponseWriter, r *http.Request) {
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
	groups, _, err := client.Users.ListGroup("", opt)
	if err != nil {
		t.Errorf("Users.ListGroup returned unexpected error: %v", err)
	}

	want := []*Group{{Name: "Test"}}
	if !reflect.DeepEqual(groups, want) {
		t.Errorf("Users.ListGroup returned %+v, want %+v", groups, want)
	}
}

func TestUsersService_JoinGroup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.JoinGroup("1", "1")
	if err != nil {
		t.Errorf("Users.JoinGroup returned unexpected error: %v", err)
	}
}

func TestUsersService_JoinGroup_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.JoinGroup("", "1")
	if err != nil {
		t.Errorf("Users.JoinGroup returned unexpected error: %v", err)
	}
}

func TestUsersService_LeaveGroup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.LeaveGroup("1", "1")
	if err != nil {
		t.Errorf("Users.LeaveGroup returned unexpected error: %v", err)
	}
}

func TestUsersService_LeaveGroup_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.LeaveGroup("", "1")
	if err != nil {
		t.Errorf("Users.LeaveGroup returned unexpected error: %v", err)
	}
}

func TestUsersService_ListLikedVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/likes", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.ListLikedVideo("1", opt)
	if err != nil {
		t.Errorf("Users.ListLikedVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.ListLikedVideo returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_ListLikedVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/likes", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.ListLikedVideo("", opt)
	if err != nil {
		t.Errorf("Users.ListLikedVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.ListLikedVideo returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_LikeVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/likes/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.LikeVideo("1", 1)
	if err != nil {
		t.Errorf("Users.LikeVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_LikeVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/likes/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.LikeVideo("", 1)
	if err != nil {
		t.Errorf("Users.LikeVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_UnlikeVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/likes/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.UnlikeVideo("1", 1)
	if err != nil {
		t.Errorf("Users.UnlikeVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_UnlikeVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/likes/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.UnlikeVideo("", 1)
	if err != nil {
		t.Errorf("Users.UnlikeVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_RemovePortrait(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/pictures/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.RemovePortrait("1", "1")
	if err != nil {
		t.Errorf("Users.RemovePortrait returned unexpected error: %v", err)
	}
}

func TestUsersService_RemovePortrait_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/pictures/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.RemovePortrait("", "1")
	if err != nil {
		t.Errorf("Users.RemovePortrait returned unexpected error: %v", err)
	}
}

func TestUsersService_ListPortfolio(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/portfolios", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	opt := &ListPortfolioOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}
	portfolios, _, err := client.Users.ListPortfolio("1", opt)
	if err != nil {
		t.Errorf("Users.ListPortfolio returned unexpected error: %v", err)
	}

	want := []*Portfolio{{Name: "Test"}}
	if !reflect.DeepEqual(portfolios, want) {
		t.Errorf("Users.ListPortfolio returned %+v, want %+v", portfolios, want)
	}
}

func TestUsersService_ListPortfolio_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/portfolios", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	opt := &ListPortfolioOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}
	portfolios, _, err := client.Users.ListPortfolio("", opt)
	if err != nil {
		t.Errorf("Users.ListPortfolio returned unexpected error: %v", err)
	}

	want := []*Portfolio{{Name: "Test"}}
	if !reflect.DeepEqual(portfolios, want) {
		t.Errorf("Users.ListPortfolio returned %+v, want %+v", portfolios, want)
	}
}

func TestUsersService_GetProtfolio(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/portfolios/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	portfolio, _, err := client.Users.GetProtfolio("1", "1")
	if err != nil {
		t.Errorf("Users.GetProtfolio returned unexpected error: %v", err)
	}

	want := &Portfolio{Name: "Test"}
	if !reflect.DeepEqual(portfolio, want) {
		t.Errorf("Users.GetProtfolio returned %+v, want %+v", portfolio, want)
	}
}

func TestUsersService_GetProtfolio_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/portfolios/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	portfolio, _, err := client.Users.GetProtfolio("", "1")
	if err != nil {
		t.Errorf("Users.GetProtfolio returned unexpected error: %v", err)
	}

	want := &Portfolio{Name: "Test"}
	if !reflect.DeepEqual(portfolio, want) {
		t.Errorf("Users.GetProtfolio returned %+v, want %+v", portfolio, want)
	}
}

func TestUsersService_ProtfolioListVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/portfolios/1/videos", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.ProtfolioListVideo("1", "1", opt)
	if err != nil {
		t.Errorf("Users.ProtfolioListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.ProtfolioListVideo returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_ProtfolioListVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/portfolios/1/videos", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.ProtfolioListVideo("", "1", opt)
	if err != nil {
		t.Errorf("Users.ProtfolioListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.ProtfolioListVideo returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_ProtfolioGetVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/portfolios/1/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Users.ProtfolioGetVideo("1", "1", 1)
	if err != nil {
		t.Errorf("Users.ProtfolioGetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Users.ProtfolioGetVideo returned %+v, want %+v", video, want)
	}
}

func TestUsersService_ProtfolioGetVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/portfolios/1/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Users.ProtfolioGetVideo("", "1", 1)
	if err != nil {
		t.Errorf("Users.ProtfolioGetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Users.ProtfolioGetVideo returned %+v, want %+v", video, want)
	}
}

func TestUsersService_ProtfolioAddVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/portfolios/1/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.ProtfolioAddVideo("1", "1", 1)
	if err != nil {
		t.Errorf("Users.ProtfolioDeleteVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_ProtfolioAddVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/portfolios/1/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.ProtfolioAddVideo("", "1", 1)
	if err != nil {
		t.Errorf("Users.ProtfolioDeleteVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_ProtfolioDeleteVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/portfolios/1/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.ProtfolioDeleteVideo("1", "1", 1)
	if err != nil {
		t.Errorf("Users.ProtfolioDeleteVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_ProtfolioDeleteVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/portfolios/1/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.ProtfolioDeleteVideo("", "1", 1)
	if err != nil {
		t.Errorf("Users.ProtfolioDeleteVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_ListPreset(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/presets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	opt := &ListPresetOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}
	presets, _, err := client.Users.ListPreset("1", opt)
	if err != nil {
		t.Errorf("Users.ListPreset returned unexpected error: %v", err)
	}

	want := []*Preset{{Name: "Test"}}
	if !reflect.DeepEqual(presets, want) {
		t.Errorf("Users.ListPreset returned %+v, want %+v", presets, want)
	}
}

func TestUsersService_ListPreset_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/presets", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"page":     "1",
			"per_page": "2",
		})
		fmt.Fprint(w, `{"data": [{"name": "Test"}]}`)
	})

	opt := &ListPresetOptions{
		ListOptions: ListOptions{Page: 1, PerPage: 2},
	}
	presets, _, err := client.Users.ListPreset("", opt)
	if err != nil {
		t.Errorf("Users.ListPreset returned unexpected error: %v", err)
	}

	want := []*Preset{{Name: "Test"}}
	if !reflect.DeepEqual(presets, want) {
		t.Errorf("Users.ListPreset returned %+v, want %+v", presets, want)
	}
}

func TestUsersService_GetPreset(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/presets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	preset, _, err := client.Users.GetPreset("1", 1)
	if err != nil {
		t.Errorf("Users.GetPreset returned unexpected error: %v", err)
	}

	want := &Preset{Name: "Test"}
	if !reflect.DeepEqual(preset, want) {
		t.Errorf("Users.GetPreset returned %+v, want %+v", preset, want)
	}
}

func TestUsersService_GetPreset_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/presets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	preset, _, err := client.Users.GetPreset("", 1)
	if err != nil {
		t.Errorf("Users.GetPreset returned unexpected error: %v", err)
	}

	want := &Preset{Name: "Test"}
	if !reflect.DeepEqual(preset, want) {
		t.Errorf("Users.GetPreset returned %+v, want %+v", preset, want)
	}
}

func TestUsersService_PresetListVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/presets/1/videos", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.PresetListVideo("1", 1, opt)
	if err != nil {
		t.Errorf("Users.PresetListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.PresetListVideo returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_PresetListVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/presets/1/videos", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.PresetListVideo("", 1, opt)
	if err != nil {
		t.Errorf("Users.PresetListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.PresetListVideo returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_ListVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/videos", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.ListVideo("1", opt)
	if err != nil {
		t.Errorf("Users.ListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.ListVideo returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_ListVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/videos", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.ListVideo("", opt)
	if err != nil {
		t.Errorf("Users.ListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.ListVideo returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_GetVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Users.GetVideo("1", 1)
	if err != nil {
		t.Errorf("Users.GetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Users.GetVideo returned %+v, want %+v", video, want)
	}
}

func TestUsersService_GetVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Users.GetVideo("", 1)
	if err != nil {
		t.Errorf("Users.GetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Users.GetVideo returned %+v, want %+v", video, want)
	}
}

func TestUsersService_WatchLaterListVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/watchlater", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.WatchLaterListVideo("1", opt)
	if err != nil {
		t.Errorf("Users.WatchLaterListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.WatchLaterListVideo returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_WatchLaterListVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/watchlater", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.WatchLaterListVideo("", opt)
	if err != nil {
		t.Errorf("Users.WatchLaterListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.WatchLaterListVideo returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_WatchLaterGetVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/watchlater/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Users.WatchLaterGetVideo("1", 1)
	if err != nil {
		t.Errorf("Users.WatchLaterGetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Users.WatchLaterGetVideo returned %+v, want %+v", video, want)
	}
}

func TestUsersService_WatchLaterGetVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/watchlater/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Users.WatchLaterGetVideo("", 1)
	if err != nil {
		t.Errorf("Users.WatchLaterGetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Users.WatchLaterGetVideo returned %+v, want %+v", video, want)
	}
}

func TestUsersService_WatchLaterAddVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/watchlater/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.WatchLaterAddVideo("1", 1)
	if err != nil {
		t.Errorf("Users.WatchLaterAddVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_WatchLaterAddVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/watchlater/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Users.WatchLaterAddVideo("", 1)
	if err != nil {
		t.Errorf("Users.WatchLaterAddVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_WatchLaterDeleteVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/users/1/watchlater/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.WatchLaterDeleteVideo("1", 1)
	if err != nil {
		t.Errorf("Users.WatchLaterDeleteVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_WatchLaterDeleteVideo_authenticatedUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/watchlater/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.WatchLaterDeleteVideo("", 1)
	if err != nil {
		t.Errorf("Users.WatchLaterDeleteVideo returned unexpected error: %v", err)
	}
}

func TestUsersService_WatchedListVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/watched/videos", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Users.WatchedListVideo("", opt)
	if err != nil {
		t.Errorf("Users.WatchedListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Users.WatchedListVideo returned %+v, want %+v", videos, want)
	}
}

func TestUsersService_ClearWatchedList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/watched/videos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.ClearWatchedList("")
	if err != nil {
		t.Errorf("Users.ClearWatchedList returned unexpected error: %v", err)
	}
}

func TestUsersService_WatchedDeleteVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/watched/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Users.WatchedDeleteVideo("", 1)
	if err != nil {
		t.Errorf("Users.WatchedDeleteVideo returned unexpected error: %v", err)
	}
}
