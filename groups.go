package vimeo

import (
	"fmt"
	"strings"
	"time"
)

// GroupsService handles communication with the group related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/endpoints/groups
type GroupsService service

type dataListGroup struct {
	Data []*Group `json:"data,omitempty"`
	pagination
}

// Group represents a group.
type Group struct {
	URI          string    `json:"uri,omitempty"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	Link         string    `json:"link,omitempty"`
	CreatedTime  time.Time `json:"created_time,omitempty"`
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	Privacy      *Privacy  `json:"privacy,omitempty"`
	Pictures     *Pictures `json:"pictures,omitempty"`
	Header       *Header   `json:"header,omitempty"`
	User         *User     `json:"user,omitempty"`
	ResourceKey  string    `json:"resource_key,omitempty"`
}

// GroupRequest represents a request to create/edit an group.
type GroupRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// ListGroupOptions specifies the optional parameters to ListGroup method.
type ListGroupOptions struct {
	Query     string `url:"query,omitempty"`
	Filter    string `url:"filter,omitempty"`
	Sort      string `url:"sort,omitempty"`
	Direction string `url:"direction,omitempty"`
	ListOptions
}

// GetID returns the identifier (ID) of the group.
func (g Group) GetID() string {
	l := strings.SplitN(g.URI, "/", -1)
	id := l[len(l)-1]
	return id
}

func listGroup(c *Client, url string, opt *ListGroupOptions) ([]*Group, *Response, error) {
	u, err := addOptions(url, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	groups := &dataListGroup{}

	resp, err := c.Do(req, groups)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(groups)

	return groups.Data, resp, err
}

// List lists the group.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/groups
func (s *GroupsService) List(opt *ListGroupOptions) ([]*Group, *Response, error) {
	groups, resp, err := listGroup(s.client, "groups", opt)

	return groups, resp, err
}

// Create a new group.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/groups
func (s *GroupsService) Create(r *GroupRequest) (*Group, *Response, error) {
	req, err := s.client.NewRequest("POST", "groups", r)
	if err != nil {
		return nil, nil, err
	}

	group := &Group{}
	resp, err := s.client.Do(req, group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}

// Get specific group by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/groups/%7Bgroup_id%7D
func (s *GroupsService) Get(gr string) (*Group, *Response, error) {
	u := fmt.Sprintf("groups/%s", gr)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	group := &Group{}

	resp, err := s.client.Do(req, group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, err
}

// Delete specific group by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/groups/%7Bgroup_id%7D
func (s *GroupsService) Delete(gr string) (*Response, error) {
	u := fmt.Sprintf("groups/%s", gr)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListUser lists the user for an group.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/groups/%7Bgroup_id%7D/users
func (s *GroupsService) ListUser(gr string, opt *ListUserOptions) ([]*User, *Response, error) {
	u := fmt.Sprintf("groups/%s/users", gr)
	users, resp, err := listUser(s.client, u, opt)

	return users, resp, err
}

// ListVideo lists the video for an group.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/groups/%7Bgroup_id%7D/videos
func (s *GroupsService) ListVideo(gr string, opt *ListVideoOptions) ([]*Video, *Response, error) {
	u := fmt.Sprintf("groups/%s/videos", gr)
	videos, resp, err := listVideo(s.client, u, opt)

	return videos, resp, err
}

// GetVideo specific video by group name and video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/groups/%7Bgroup_id%7D/videos/%7Bvideo_id%7D
func (s *GroupsService) GetVideo(gr string, vid int) (*Video, *Response, error) {
	u := fmt.Sprintf("groups/%s/videos/%d", gr, vid)
	video, resp, err := getVideo(s.client, u)

	return video, resp, err
}

// DeleteVideo specific video by group name and video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/groups/%7Bgroup_id%7D/videos/%7Bvideo_id%7D
func (s *GroupsService) DeleteVideo(gr string, vid int) (*Response, error) {
	u := fmt.Sprintf("groups/%s/videos/%d", gr, vid)
	resp, err := deleteVideo(s.client, u)

	return resp, err
}
