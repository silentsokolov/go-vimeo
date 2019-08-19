package vimeo

import (
	"fmt"
	"strings"
	"time"
)

// ChannelsService handles communication with the channels related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels
type ChannelsService service

type dataListChannel struct {
	Data []*Channel `json:"data"`
	pagination
}

// Channel represents a channel.
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

// GetID returns the identifier (ID) of the channel.
func (c Channel) GetID() string {
	l := strings.SplitN(c.URI, "/", -1)
	id := l[len(l)-1]
	return id
}

func listChannel(c *Client, url string, opt ...CallOption) ([]*Channel, *Response, error) {
	u, err := addOptions(url, opt...)
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

// List method gets all existing channels.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels#get_channels
func (s *ChannelsService) List(opt ...CallOption) ([]*Channel, *Response, error) {
	channels, resp, err := listChannel(s.client, "channels", opt...)

	return channels, resp, err
}

// Create method creates a new channel.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels#create_channel
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

// Get method gets a single channel.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels#get_channel
func (s *ChannelsService) Get(ch string, opt ...CallOption) (*Channel, *Response, error) {
	u, err := addOptions(fmt.Sprintf("channels/%s", ch), opt...)
	if err != nil {
		return nil, nil, err
	}

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

// Edit method edits the specified channel.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels#edit_channel
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

// Delete method deletes the specified channel.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels#delete_channel
func (s *ChannelsService) Delete(ch string) (*Response, error) {
	u := fmt.Sprintf("channels/%s", ch)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListUser method gets all the followers of a specific channel.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels#get_channel_subscribers
func (s *ChannelsService) ListUser(ch string, opt ...CallOption) ([]*User, *Response, error) {
	u := fmt.Sprintf("channels/%s/users", ch)
	users, resp, err := listUser(s.client, u, opt...)

	return users, resp, err
}

// ListVideo method gets all the videos in a specific channel.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels#get_channel_videos
func (s *ChannelsService) ListVideo(ch string, opt ...CallOption) ([]*Video, *Response, error) {
	u := fmt.Sprintf("channels/%s/videos", ch)
	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}

// GetVideo method returns a specific video in a channel. You can use it to determine whether the video is in the channel.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels#get_channel_video
func (s *ChannelsService) GetVideo(ch string, vid int, opt ...CallOption) (*Video, *Response, error) {
	u := fmt.Sprintf("channels/%s/videos/%d", ch, vid)
	video, resp, err := getVideo(s.client, u, opt...)

	return video, resp, err
}

// AddVideo method adds a single video to the specified channel.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels#add_video_to_channel
func (s *ChannelsService) AddVideo(ch string, vid int) (*Video, *Response, error) {
	u := fmt.Sprintf("channels/%s/videos/%d", ch, vid)
	video, resp, err := addVideo(s.client, u)

	return video, resp, err
}

// DeleteVideo method removes a single video from the channel in question.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels#delete_video_from_channel
func (s *ChannelsService) DeleteVideo(ch string, vid int) (*Response, error) {
	u := fmt.Sprintf("channels/%s/videos/%d", ch, vid)
	resp, err := deleteVideo(s.client, u)

	return resp, err
}
