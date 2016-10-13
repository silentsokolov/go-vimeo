package vimeo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMeService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	user, _, err := client.Me.Get()
	if err != nil {
		t.Errorf("Me.Get returned unexpected error: %v", err)
	}

	want := &User{Name: "Test"}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Me.Get returned %+v, want %+v", user, want)
	}
}

func TestMeService_Edit(t *testing.T) {
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
			t.Errorf("Me.Edit body is %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"name": "name"}`)
	})

	user, _, err := client.Me.Edit(input)
	if err != nil {
		t.Errorf("Me.Edit returned unexpected error: %v", err)
	}

	want := &User{Name: "name"}
	if !reflect.DeepEqual(user, want) {
		t.Errorf("Me.Edit returned %+v, want %+v", user, want)
	}
}

func TestMeService_ListAlbum(t *testing.T) {
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
	albums, _, err := client.Me.ListAlbum(opt)
	if err != nil {
		t.Errorf("Me.ListAlbum returned unexpected error: %v", err)
	}

	want := []*Album{{Name: "Test"}}
	if !reflect.DeepEqual(albums, want) {
		t.Errorf("Me.ListAlbum returned %+v, want %+v", albums, want)
	}
}

func TestMeService_CreateAlbum(t *testing.T) {
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
			t.Errorf("Me.CreateAlbum body is %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"name": "name"}`)
	})

	album, _, err := client.Me.CreateAlbum(input)
	if err != nil {
		t.Errorf("Me.CreateAlbum returned unexpected error: %v", err)
	}

	want := &Album{Name: "name"}
	if !reflect.DeepEqual(album, want) {
		t.Errorf("Me.CreateAlbum returned %+v, want %+v", album, want)
	}
}

func TestMeService_GetAlbum(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/albums/a", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	album, _, err := client.Me.GetAlbum("a")
	if err != nil {
		t.Errorf("Me.GetAlbum returned unexpected error: %v", err)
	}

	want := &Album{Name: "Test"}
	if !reflect.DeepEqual(album, want) {
		t.Errorf("Me.GetAlbum returned %+v, want %+v", album, want)
	}
}

func TestMeService_EditAlbum(t *testing.T) {
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
			t.Errorf("Me.EditAlbum body is %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"name": "name"}`)
	})

	album, _, err := client.Me.EditAlbum("a", input)
	if err != nil {
		t.Errorf("Me.Edit returned unexpected error: %v", err)
	}

	want := &Album{Name: "name"}
	if !reflect.DeepEqual(album, want) {
		t.Errorf("Me.EditAlbum returned %+v, want %+v", album, want)
	}
}

func TestMeService_DeleteAlbum(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/albums/a", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Me.DeleteAlbum("a")
	if err != nil {
		t.Errorf("Me.DeleteAlbum returned unexpected error: %v", err)
	}
}

func TestMeService_AlbumListVideo(t *testing.T) {
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
	videos, _, err := client.Me.AlbumListVideo("a", opt)
	if err != nil {
		t.Errorf("Me.AlbumListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Me.AlbumListVideo returned %+v, want %+v", videos, want)
	}
}

func TestMeService_AlbumGetVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/albums/a/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Me.AlbumGetVideo("a", 1)
	if err != nil {
		t.Errorf("Me.AlbumGetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Me.AlbumGetVideo returned %+v, want %+v", video, want)
	}
}

func TestMeService_AlbumDeleteVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/albums/a/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Me.AlbumDeleteVideo("a", 1)
	if err != nil {
		t.Errorf("Me.AlbumDeleteVideo returned unexpected error: %v", err)
	}
}

func TestMeService_ListAppearance(t *testing.T) {
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
	videos, _, err := client.Me.ListAppearance(opt)
	if err != nil {
		t.Errorf("Me.ListAppearance returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Me.ListAppearance returned %+v, want %+v", videos, want)
	}
}

func TestMeService_ListCategory(t *testing.T) {
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
	categories, _, err := client.Me.ListCategory(opt)
	if err != nil {
		t.Errorf("Me.ListCategory returned unexpected error: %v", err)
	}

	want := []*Category{{Name: "Test"}}
	if !reflect.DeepEqual(categories, want) {
		t.Errorf("Me.ListCategory returned %+v, want %+v", categories, want)
	}
}

func TestMeService_SubscribeCategory(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/categories/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Me.SubscribeCategory("1")
	if err != nil {
		t.Errorf("Me.SubscribeCategory returned unexpected error: %v", err)
	}
}

func TestMeService_UnsubscribeCategory(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/categories/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Me.UnsubscribeCategory("1")
	if err != nil {
		t.Errorf("Me.UnsubscribeCategory returned unexpected error: %v", err)
	}
}

func TestMeService_ListChannel(t *testing.T) {
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
	channels, _, err := client.Me.ListChannel(opt)
	if err != nil {
		t.Errorf("Me.ListChannel returned unexpected error: %v", err)
	}

	want := []*Channel{{Name: "Test"}}
	if !reflect.DeepEqual(channels, want) {
		t.Errorf("Me.ListChannel returned %+v, want %+v", channels, want)
	}
}

func TestMeService_SubscribeChannel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/channels/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Me.SubscribeChannel("1")
	if err != nil {
		t.Errorf("Me.SubscribeChannel returned unexpected error: %v", err)
	}
}

func TestMeService_UnsubscribeChannel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/channels/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Me.UnsubscribeChannel("1")
	if err != nil {
		t.Errorf("Me.UnsubscribeChannel returned unexpected error: %v", err)
	}
}

func TestMeService_Feed(t *testing.T) {
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
	feed, _, err := client.Me.Feed(opt)
	if err != nil {
		t.Errorf("Me.Feed returned unexpected error: %v", err)
	}

	want := []*Feed{{URI: "/1"}}
	if !reflect.DeepEqual(feed, want) {
		t.Errorf("Me.Feed returned %+v, want %+v", feed, want)
	}
}

func TestMeService_ListFollower(t *testing.T) {
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
	users, _, err := client.Me.ListFollower(opt)
	if err != nil {
		t.Errorf("Me.ListFollower returned unexpected error: %v", err)
	}

	want := []*User{{Name: "Test"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Me.ListFollower returned %+v, want %+v", users, want)
	}
}

func TestMeService_ListFollowed(t *testing.T) {
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
	users, _, err := client.Me.ListFollowed(opt)
	if err != nil {
		t.Errorf("Me.ListFollowed returned unexpected error: %v", err)
	}

	want := []*User{{Name: "Test"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Me.ListFollowed returned %+v, want %+v", users, want)
	}
}

func TestMeService_FollowUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/following/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Me.FollowUser("1")
	if err != nil {
		t.Errorf("Me.FollowUser returned unexpected error: %v", err)
	}
}

func TestMeService_UnfollowUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/following/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Me.UnfollowUser("1")
	if err != nil {
		t.Errorf("Me.UnfollowUser returned unexpected error: %v", err)
	}
}

func TestMeService_ListGroup(t *testing.T) {
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
	groups, _, err := client.Me.ListGroup(opt)
	if err != nil {
		t.Errorf("Me.ListGroup returned unexpected error: %v", err)
	}

	want := []*Group{{Name: "Test"}}
	if !reflect.DeepEqual(groups, want) {
		t.Errorf("Me.ListGroup returned %+v, want %+v", groups, want)
	}
}

func TestMeService_JoinGroup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Me.JoinGroup("1")
	if err != nil {
		t.Errorf("Me.JoinGroup returned unexpected error: %v", err)
	}
}

func TestMeService_LeaveGroup(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/groups/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Me.LeaveGroup("1")
	if err != nil {
		t.Errorf("Me.LeaveGroup returned unexpected error: %v", err)
	}
}

func TestMeService_ListLikedVideo(t *testing.T) {
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
	videos, _, err := client.Me.ListLikedVideo(opt)
	if err != nil {
		t.Errorf("Me.ListLikedVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Me.ListLikedVideo returned %+v, want %+v", videos, want)
	}
}

func TestMeService_LikeVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/likes/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Me.LikeVideo(1)
	if err != nil {
		t.Errorf("Me.LikeVideo returned unexpected error: %v", err)
	}
}

func TestMeService_UnlikeVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/likes/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Me.UnlikeVideo(1)
	if err != nil {
		t.Errorf("Me.UnlikeVideo returned unexpected error: %v", err)
	}
}

func TestMeService_RemovePortrait(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/pictures/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Me.RemovePortrait("1")
	if err != nil {
		t.Errorf("Me.RemovePortrait returned unexpected error: %v", err)
	}
}

func TestMeService_ListPortfolio(t *testing.T) {
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
	portfolios, _, err := client.Me.ListPortfolio(opt)
	if err != nil {
		t.Errorf("Me.ListPortfolio returned unexpected error: %v", err)
	}

	want := []*Portfolio{{Name: "Test"}}
	if !reflect.DeepEqual(portfolios, want) {
		t.Errorf("Me.ListPortfolio returned %+v, want %+v", portfolios, want)
	}
}

func TestMeService_GetProtfolio(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/portfolios/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	portfolio, _, err := client.Me.GetProtfolio("1")
	if err != nil {
		t.Errorf("Me.GetProtfolio returned unexpected error: %v", err)
	}

	want := &Portfolio{Name: "Test"}
	if !reflect.DeepEqual(portfolio, want) {
		t.Errorf("Me.GetProtfolio returned %+v, want %+v", portfolio, want)
	}
}

func TestMeService_ProtfolioListVideo(t *testing.T) {
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
	videos, _, err := client.Me.ProtfolioListVideo("1", opt)
	if err != nil {
		t.Errorf("Me.ProtfolioListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Me.ProtfolioListVideo returned %+v, want %+v", videos, want)
	}
}

func TestMeService_ProtfolioGetVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/portfolios/1/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Me.ProtfolioGetVideo("1", 1)
	if err != nil {
		t.Errorf("Me.ProtfolioGetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Me.ProtfolioGetVideo returned %+v, want %+v", video, want)
	}
}

func TestMeService_ProtfolioDeleteVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/portfolios/1/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Me.ProtfolioDeleteVideo("1", 1)
	if err != nil {
		t.Errorf("Me.ProtfolioDeleteVideo returned unexpected error: %v", err)
	}
}

func TestMeService_ListPreset(t *testing.T) {
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
	presets, _, err := client.Me.ListPreset(opt)
	if err != nil {
		t.Errorf("Me.ListPreset returned unexpected error: %v", err)
	}

	want := []*Preset{{Name: "Test"}}
	if !reflect.DeepEqual(presets, want) {
		t.Errorf("Me.ListPreset returned %+v, want %+v", presets, want)
	}
}

func TestMeService_GetPreset(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/presets/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	preset, _, err := client.Me.GetPreset(1)
	if err != nil {
		t.Errorf("Me.GetPreset returned unexpected error: %v", err)
	}

	want := &Preset{Name: "Test"}
	if !reflect.DeepEqual(preset, want) {
		t.Errorf("Me.GetPreset returned %+v, want %+v", preset, want)
	}
}

func TestMeService_PresetListVideo(t *testing.T) {
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
	videos, _, err := client.Me.PresetListVideo("1", opt)
	if err != nil {
		t.Errorf("Me.PresetListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Me.PresetListVideo returned %+v, want %+v", videos, want)
	}
}

func TestMeService_ListVideo(t *testing.T) {
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
	videos, _, err := client.Me.ListVideo(opt)
	if err != nil {
		t.Errorf("Me.ListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Me.ListVideo returned %+v, want %+v", videos, want)
	}
}

func TestMeService_GetVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Me.GetVideo(1)
	if err != nil {
		t.Errorf("Me.GetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Me.GetVideo returned %+v, want %+v", video, want)
	}
}

func TestMeService_WatchLaterListVideo(t *testing.T) {
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
	videos, _, err := client.Me.WatchLaterListVideo(opt)
	if err != nil {
		t.Errorf("Me.WatchLaterListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Me.WatchLaterListVideo returned %+v, want %+v", videos, want)
	}
}

func TestMeService_WatchLaterGetVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/watchlater/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Me.WatchLaterGetVideo(1)
	if err != nil {
		t.Errorf("Me.WatchLaterGetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Me.WatchLaterGetVideo returned %+v, want %+v", video, want)
	}
}

func TestMeService_WatchLaterAddVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/watchlater/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Me.WatchLaterAddVideo(1)
	if err != nil {
		t.Errorf("Me.WatchLaterAddVideo returned unexpected error: %v", err)
	}
}

func TestMeService_WatchLaterDeleteVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/watchlater/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Me.WatchLaterDeleteVideo(1)
	if err != nil {
		t.Errorf("Me.WatchLaterDeleteVideo returned unexpected error: %v", err)
	}
}

func TestMeService_WatchedListVideo(t *testing.T) {
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
	videos, _, err := client.Me.WatchedListVideo(opt)
	if err != nil {
		t.Errorf("Me.WatchedListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Me.WatchedListVideo returned %+v, want %+v", videos, want)
	}
}

func TestMeService_ClearWatchedList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/watched/videos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Me.ClearWatchedList()
	if err != nil {
		t.Errorf("Me.ClearWatchedList returned unexpected error: %v", err)
	}
}

func TestMeService_WatchedDeleteVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/me/watched/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Me.WatchedDeleteVideo(1)
	if err != nil {
		t.Errorf("Me.WatchedDeleteVideo returned unexpected error: %v", err)
	}
}
