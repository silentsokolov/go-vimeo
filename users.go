package vimeo

import (
	"fmt"
	"time"
)

// UsersService handles communication with the users related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/endpoints/users
type UsersService service

type dataListUser struct {
	Data []*User `json:"data"`
	pagination
}

// WebSite represents a web site.
type WebSite struct {
	Name        string `json:"name,omitempty"`
	Link        string `json:"link,omitempty"`
	Description string `json:"description,omitempty"`
}

// User represents a user.
type User struct {
	URI           string     `json:"uri,omitempty"`
	Name          string     `json:"name,omitempty"`
	Link          string     `json:"link,omitempty"`
	Location      string     `json:"location,omitempty"`
	Bio           string     `json:"bio,omitempty"`
	CreatedTime   time.Time  `json:"created_time,omitempty"`
	Account       string     `json:"account,omitempty"`
	Pictures      *Pictures  `json:"pictures,omitempty"`
	WebSites      []*WebSite `json:"websites,omitempty"`
	ContentFilter []string   `json:"content_filter,omitempty"`
	ResourceKey   string     `json:"resource_key,omitempty"`
}

// ListUserOptions specifies the optional parameters to the
// ListUser method.
type ListUserOptions struct {
	Query  string `url:"query,omitempty"`
	Filter string `url:"filter,omitempty"`
	ListOptions
}

// UserRequest represents a request to create/edit an user.
type UserRequest struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
	Bio      string `json:"bio,omitempty"`
}

func listUser(c *Client, url string, opt *ListUserOptions) ([]*User, *Response, error) {
	u, err := addOptions(url, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	users := &dataListUser{}

	resp, err := c.Do(req, users)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(users)

	return users.Data, resp, err
}

// Search users.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/channels/%7Bchannel_id%7D/users
func (s *UsersService) Search(opt *ListUserOptions) ([]*User, *Response, error) {
	users, resp, err := listUser(s.client, "users", opt)

	return users, resp, err
}

// Get show one user.
// Passing the empty string will authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D
func (s *UsersService) Get(uid string) (*User, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me")
	} else {
		u = fmt.Sprintf("users/%s", uid)
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	user := &User{}

	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, err
}

// Edit one user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D
func (s *UsersService) Edit(uid string, r *UserRequest) (*User, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me")
	} else {
		u = fmt.Sprintf("users/%s", uid)
	}

	req, err := s.client.NewRequest("PATCH", u, r)
	if err != nil {
		return nil, nil, err
	}

	user := &User{}
	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// ListAppearance all videos a user is credited in.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/appearances
func (s *UsersService) ListAppearance(uid string, opt *ListVideoOptions) ([]*Video, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/appearances")
	} else {
		u = fmt.Sprintf("users/%s/appearances", uid)
	}

	videos, resp, err := listVideo(s.client, u, opt)

	return videos, resp, err
}
