package vimeo

import (
	"fmt"
	"time"
)

type dataListAlbum struct {
	Data []*Album `json:"data"`
	pagination
}

// Privacy internal object provides access to privacy.
type Privacy struct {
	View     string `json:"view,omitempty"`
	Join     string `json:"join,omitempty"`
	Videos   string `json:"videos,omitempty"`
	Comment  string `json:"comment,omitempty"`
	Forums   string `json:"forums,omitempty"`
	Invite   string `json:"invite,omitempty"`
	Embed    string `json:"embed,omitempty"`
	Download bool   `json:"download"`
	Add      bool   `json:"add"`
}

// Album represents a album.
type Album struct {
	URI          string    `json:"uri,omitempty"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	Link         string    `json:"link,omitempty"`
	Duration     int       `json:"duration,omitempty"`
	CreatedTime  time.Time `json:"created_time,omitempty"`
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	User         *User     `json:"user,omitempty"`
	Pictures     *Pictures `json:"pictures,omitempty"`
	Privacy      *Privacy  `json:"privacy,omitempty"`
}

// AlbumRequest represents a request to create/edit an album.
type AlbumRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Privacy     string `json:"privacy,omitempty"`
	Password    string `json:"password,omitempty"`
	Sort        string `json:"sort,omitempty"`
}

// ListAlbum method gets all the albums from the specified user's account.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/albums#get_albums
func (s *UsersService) ListAlbum(uid string, opt ...CallOption) ([]*Album, *Response, error) {
	var u string
	if uid == "" {
		u = "me/albums"
	} else {
		u = fmt.Sprintf("users/%s/albums", uid)
	}

	u, err := addOptions(u, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	albums := &dataListAlbum{}

	resp, err := s.client.Do(req, albums)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(albums)

	return albums.Data, resp, err
}

// CreateAlbum method creates a new album for the specified user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/albums#create_album
func (s *UsersService) CreateAlbum(uid string, r *AlbumRequest) (*Album, *Response, error) {
	var u string
	if uid == "" {
		u = "me/albums"
	} else {
		u = fmt.Sprintf("users/%s/albums", uid)
	}

	req, err := s.client.NewRequest("POST", u, r)
	if err != nil {
		return nil, nil, err
	}

	album := &Album{}
	resp, err := s.client.Do(req, album)
	if err != nil {
		return nil, resp, err
	}

	return album, resp, nil
}

// GetAlbum method gets a single album.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/albums#get_album
func (s *UsersService) GetAlbum(uid string, ab string, opt ...CallOption) (*Album, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/albums/%s", ab)
	} else {
		u = fmt.Sprintf("users/%s/albums/%s", uid, ab)
	}

	u, err := addOptions(u, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	album := &Album{}

	resp, err := s.client.Do(req, album)
	if err != nil {
		return nil, resp, err
	}

	return album, resp, err
}

// EditAlbum method edits an album.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/albums#edit_album
func (s *UsersService) EditAlbum(uid string, ab string, r *AlbumRequest) (*Album, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/albums/%s", ab)
	} else {
		u = fmt.Sprintf("users/%s/albums/%s", uid, ab)
	}

	req, err := s.client.NewRequest("PATCH", u, r)
	if err != nil {
		return nil, nil, err
	}

	album := &Album{}
	resp, err := s.client.Do(req, album)
	if err != nil {
		return nil, resp, err
	}

	return album, resp, nil
}

// DeleteAlbum method deletes an album from the owner's account.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/albums#delete_album
func (s *UsersService) DeleteAlbum(uid string, ab string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/albums/%s", ab)
	} else {
		u = fmt.Sprintf("users/%s/albums/%s", uid, ab)
	}

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// AlbumListVideo method gets all the videos from the specified album.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/albums#get_album_videos
func (s *UsersService) AlbumListVideo(uid string, ab string, opt ...CallOption) ([]*Video, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/albums/%s/videos", ab)
	} else {
		u = fmt.Sprintf("users/%s/albums/%s/videos", uid, ab)
	}
	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}

// AlbumGetVideo method gets a single video from an album. You can use this method to determine whether the album contains the specified video.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/albums#get_album_video
func (s *UsersService) AlbumGetVideo(uid string, ab string, vid int, opt ...CallOption) (*Video, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/albums/%s/videos/%d", ab, vid)
	} else {
		u = fmt.Sprintf("users/%s/albums/%s/videos/%d", uid, ab, vid)
	}
	video, resp, err := getVideo(s.client, u, opt...)

	return video, resp, err
}

// AlbumAddVideo method adds a single video to the specified album.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/albums#add_video_to_album
func (s *UsersService) AlbumAddVideo(uid string, ab string, vid int) (*Video, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/albums/%s/videos/%d", ab, vid)
	} else {
		u = fmt.Sprintf("users/%s/albums/%s/videos/%d", uid, ab, vid)
	}
	video, resp, err := addVideo(s.client, u)

	return video, resp, err
}

// AlbumDeleteVideo method removes a video from the specified album.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/albums#remove_video_from_album
func (s *UsersService) AlbumDeleteVideo(uid string, ab string, vid int) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/albums/%s/videos/%d", ab, vid)
	} else {
		u = fmt.Sprintf("users/%s/albums/%s/videos/%d", uid, ab, vid)
	}

	resp, err := deleteVideo(s.client, u)

	return resp, err
}
