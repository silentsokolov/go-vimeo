package vimeo

import "time"

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

// ListAlbumOptions specifies the optional parameters to the
// ListAlbum method.
type ListAlbumOptions struct {
	Query     string `url:"query,omitempty"`
	Sort      string `url:"sort,omitempty"`
	Direction string `url:"direction,omitempty"`
	ListOptions
}

// AlbumRequest represents a request to create/edit an album.
type AlbumRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Privacy     string `json:"privacy,omitempty"`
	Password    string `json:"password,omitempty"`
	Sort        string `json:"sort,omitempty"`
}
