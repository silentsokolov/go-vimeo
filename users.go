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

// ListCategory list the subscribed category for user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/categories
func (s *UsersService) ListCategory(uid string, opt *ListCategoryOptions) ([]*Category, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/categories")
	} else {
		u = fmt.Sprintf("users/%s/categories", uid)
	}

	categories, resp, err := listCategory(s.client, u, opt)

	return categories, resp, err
}

// SubscribeCategory subscribe category user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/categories/%7Bcategory%7D
func (s *UsersService) SubscribeCategory(uid string, cat string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/categories/%s", cat)
	} else {
		u = fmt.Sprintf("users/%s/categories/%s", uid, cat)
	}

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnsubscribeCategory unsubscribe category current user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/categories/%7Bcategory%7D
func (s *UsersService) UnsubscribeCategory(uid string, cat string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/categories/%s", cat)
	} else {
		u = fmt.Sprintf("users/%s/categories/%s", uid, cat)
	}

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListChannel list the subscribed channel for user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/channels
func (s *UsersService) ListChannel(uid string, opt *ListChannelOptions) ([]*Channel, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/channels")
	} else {
		u = fmt.Sprintf("users/%s/channels", uid)
	}

	categories, resp, err := listChannel(s.client, u, opt)

	return categories, resp, err
}

// SubscribeChannel subscribe channel user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/channels/%7Bchannel_id%7D
func (s *UsersService) SubscribeChannel(uid string, ch string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/channels/%s", ch)
	} else {
		u = fmt.Sprintf("users/%s/channels/%s", uid, ch)
	}

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnsubscribeChannel unsubscribe channel user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/channels/%7Bchannel_id%7D
func (s *UsersService) UnsubscribeChannel(uid string, ch string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/channels/%s", ch)
	} else {
		u = fmt.Sprintf("users/%s/channels/%s", uid, ch)
	}

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
