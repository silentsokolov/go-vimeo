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

// ListAlbum lists the album for an current user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/albums
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

// CreateAlbum a new album.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/albums
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

// GetAlbum specific album by name.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/albums/%7Balbum_id%7D
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

// EditAlbum specific album by name.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/albums/%7Balbum_id%7D
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

// DeleteAlbum specific album by name.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/albums/%7Balbum_id%7D
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

// AlbumListVideo lists the video for an album.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/albums/%7Balbum_id%7D/videos
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

// AlbumGetVideo get specific video by album name and video ID.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/albums/%7Balbum_id%7D/videos/%7Bvideo_id%7D
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

// AlbumAddVideo add specific video by album name and video ID.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/albums/%7Balbum_id%7D/videos/%7Bvideo_id%7D
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

// AlbumDeleteVideo delete specific video by album name and video ID.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/albums/%7Balbum_id%7D/videos/%7Bvideo_id%7D
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
