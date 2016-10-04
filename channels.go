package vimeo

import (
	"fmt"
	"strings"
	"time"
)

// ChannelsService handles communication with the channels related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/endpoints/channels
type ChannelsService service

type dataListChannel struct {
	Data []*Channel `json:"data"`
	pagination
}

// Category represents a channel.
type Channel struct {
	URI          string    `json:"uri,omitempty"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	Link         string    `json:"link,omitempty"`
	CreatedTime  time.Time `json:"created_time,omitempty"`
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	User         *User     `json:"user,omitempty"`
	Pictures     *Pictures `json:"pictures,omitempty"`
	Header       *Header   `json:"header,omitempty"`
	Privacy      *Privacy  `json:"privacy,omitempty"`
	ResourceKey  string    `json:"resource_key,omitempty"`
}

// ChannelRequest represents a request to create/edit an channel.
type ChannelRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Privacy     string `json:"privacy,omitempty"`
}

// ListChannelOptions specifies the optional parameters to the
// CategoriesService.ListChannel method.
type ListChannelOptions struct {
	Query  string `url:"query,omitempty"`
	Filter string `url:"filter,omitempty"`
	ListOptions
}

// GetID returns the numeric identifier (ID) of the channel.
func (c Channel) GetID() string {
	l := strings.SplitN(c.URI, "/", -1)
	id := l[len(l)-1]
	return id
}

func listChannel(c *Client, url string, opt *ListChannelOptions) ([]*Channel, *Response, error) {
	u, err := addOptions(url, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	channels := &dataListChannel{}

	resp, err := c.Do(req, channels)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(channels)

	return channels.Data, resp, err
}

// ListChannel lists the channel for an category.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/channels
func (s *ChannelsService) List(opt *ListChannelOptions) ([]*Channel, *Response, error) {
	u, err := addOptions("channels", opt)
	if err != nil {
		return nil, nil, err
	}

	channels, resp, err := listChannel(s.client, u, opt)

	return channels, resp, err
}

// Create a new channel.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/channels
func (s *ChannelsService) Create(r *ChannelRequest) (*Channel, *Response, error) {
	req, err := s.client.NewRequest("POST", "channels", r)
	if err != nil {
		return nil, nil, err
	}

	channel := &Channel{}
	resp, err := s.client.Do(req, channel)
	if err != nil {
		return nil, resp, err
	}

	return channel, resp, nil
}

// Get specific channel by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/channels/%7Bchannel_id%7D
func (s *ChannelsService) Get(ch string) (*Channel, *Response, error) {
	u := fmt.Sprintf("channels/%s", ch)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	channel := &Channel{}

	resp, err := s.client.Do(req, channel)
	if err != nil {
		return nil, resp, err
	}

	return channel, resp, err
}

// Edit specific channel by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/channels/%7Bchannel_id%7D
func (s *ChannelsService) Edit(ch string, r *ChannelRequest) (*Channel, *Response, error) {
	u := fmt.Sprintf("channels/%s", ch)
	req, err := s.client.NewRequest("PATCH", u, r)
	if err != nil {
		return nil, nil, err
	}

	channel := &Channel{}
	resp, err := s.client.Do(req, channel)
	if err != nil {
		return nil, resp, err
	}

	return channel, resp, nil
}

// Delete specific channel by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/channels/%7Bchannel_id%7D
func (s *ChannelsService) Delete(ch string) (*Response, error) {
	u := fmt.Sprintf("channels/%s", ch)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListUser lists the user for an channel.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/channels/%7Bchannel_id%7D/users
func (s *ChannelsService) ListUser(ch string, opt *ListUserOptions) ([]*User, *Response, error) {
	u := fmt.Sprintf("channels/%s/users", ch)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	users := &dataListUser{}

	resp, err := s.client.Do(req, users)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(users)

	return users.Data, resp, err
}

// ListVideo lists the video for an channel.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/channels/%7Bchannel_id%7D/videos
func (s *ChannelsService) ListVideo(ch string, opt *ListVideoOptions) ([]*Video, *Response, error) {
	u := fmt.Sprintf("channels/%s/videos", ch)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	videos := &dataListVideo{}

	resp, err := s.client.Do(req, videos)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(videos)

	return videos.Data, resp, err
}

// GetVideo specific video by channel name and video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/channels/%7Bchannel_id%7D/videos/%7Bvideo_id%7D
func (s *ChannelsService) GetVideo(ch string, vid int) (*Video, *Response, error) {
	u := fmt.Sprintf("channels/%s/videos/%d", ch, vid)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	video := &Video{}

	resp, err := s.client.Do(req, video)
	if err != nil {
		return nil, resp, err
	}

	return video, resp, err
}

// Delete specific video by channel name and video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/channels/%7Bchannel_id%7D/videos/%7Bvideo_id%7D
func (s *ChannelsService) DeleteVideo(ch string, vid int) (*Response, error) {
	u := fmt.Sprintf("channels/%s/videos/%d", ch, vid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
