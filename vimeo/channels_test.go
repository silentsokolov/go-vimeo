package vimeo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestChannel_GetID(t *testing.T) {
	v := &Channel{Name: "Test", URI: "/channels/1"}

	if id := v.GetID(); id != "1" {
		t.Errorf("Channel.GetID returned %+v, want %+v", id, "1")
	}
}

func TestChannelsService_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channels", func(w http.ResponseWriter, r *http.Request) {
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
	channels, _, err := client.Channels.List(opt)
	if err != nil {
		t.Errorf("Channels.List returned unexpected error: %v", err)
	}

	want := []*Channel{{Name: "Test"}}
	if !reflect.DeepEqual(channels, want) {
		t.Errorf("Channels.List returned %+v, want %+v", channels, want)
	}
}

func TestChannelsService_Create(t *testing.T) {
	setup()
	defer teardown()

	input := &ChannelRequest{
		Name:        "name",
		Description: "desc",
		Privacy:     "anybody",
	}

	mux.HandleFunc("/channels", func(w http.ResponseWriter, r *http.Request) {
		v := &ChannelRequest{}
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Channels.Create body is %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"name": "name"}`)
	})

	channel, _, err := client.Channels.Create(input)
	if err != nil {
		t.Errorf("Channels.Create returned unexpected error: %v", err)
	}

	want := &Channel{Name: "name"}
	if !reflect.DeepEqual(channel, want) {
		t.Errorf("Channels.Create returned %+v, want %+v", channel, want)
	}
}

func TestChannelsService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channels/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	channel, _, err := client.Channels.Get("1")
	if err != nil {
		t.Errorf("Channels.Get returned unexpected error: %v", err)
	}

	want := &Channel{Name: "Test"}
	if !reflect.DeepEqual(channel, want) {
		t.Errorf("Channels.Get returned %+v, want %+v", channel, want)
	}
}

func TestChannelsService_Edit(t *testing.T) {
	setup()
	defer teardown()

	input := &ChannelRequest{
		Name:        "name",
		Description: "desc",
		Privacy:     "anybody",
	}

	mux.HandleFunc("/channels/1", func(w http.ResponseWriter, r *http.Request) {
		v := &ChannelRequest{}
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "PATCH")
		if !reflect.DeepEqual(v, input) {
			t.Errorf("Channels.Edit body is %+v, want %+v", v, input)
		}

		fmt.Fprint(w, `{"name": "name"}`)
	})

	channel, _, err := client.Channels.Edit("1", input)
	if err != nil {
		t.Errorf("Channels.Edit returned unexpected error: %v", err)
	}

	want := &Channel{Name: "name"}
	if !reflect.DeepEqual(channel, want) {
		t.Errorf("Channels.Edit returned %+v, want %+v", channel, want)
	}
}

func TestChannelsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channels/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Channels.Delete("1")
	if err != nil {
		t.Errorf("Channels.Delete returned unexpected error: %v", err)
	}
}

func TestChannelsService_ListUser(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channels/1/users", func(w http.ResponseWriter, r *http.Request) {
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
	users, _, err := client.Channels.ListUser("1", opt)
	if err != nil {
		t.Errorf("Channels.ListUser returned unexpected error: %v", err)
	}

	want := []*User{{Name: "Test"}}
	if !reflect.DeepEqual(users, want) {
		t.Errorf("Channels.ListUser returned %+v, want %+v", users, want)
	}
}

func TestChannelsService_ListVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channels/1/videos", func(w http.ResponseWriter, r *http.Request) {
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
	videos, _, err := client.Channels.ListVideo("1", opt)
	if err != nil {
		t.Errorf("Channels.ListVideo returned unexpected error: %v", err)
	}

	want := []*Video{{Name: "Test"}}
	if !reflect.DeepEqual(videos, want) {
		t.Errorf("Channels.ListVideo returned %+v, want %+v", videos, want)
	}
}

func TestChannelsService_GetVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channels/ch/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name": "Test"}`)
	})

	video, _, err := client.Channels.GetVideo("ch", 1)
	if err != nil {
		t.Errorf("Channels.GetVideo returned unexpected error: %v", err)
	}

	want := &Video{Name: "Test"}
	if !reflect.DeepEqual(video, want) {
		t.Errorf("Channels.GetVideo returned %+v, want %+v", video, want)
	}
}

func TestChannelsService_DeleteVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channels/ch/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
	})

	_, err := client.Channels.DeleteVideo("ch", 1)
	if err != nil {
		t.Errorf("Channels.DeleteVideo returned unexpected error: %v", err)
	}
}

func TestChannelsService_AddVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/channels/ch/videos/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
	})

	_, err := client.Channels.AddVideo("ch", 1)
	if err != nil {
		t.Errorf("Channels.AddVideo returned unexpected error: %v", err)
	}
}
