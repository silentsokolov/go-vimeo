package vimeo

import "fmt"

// TagsService handles communication with the tag related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/endpoints/tags
type TagsService service

// Tag represents a tag.
type Tag struct {
	URI         string `json:"uri,omitempty"`
	Name        string `json:"name,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Canonical   string `json:"canonical,omitempty"`
	ResourceKey string `json:"resource_key,omitempty"`
}

// Get specific tag by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/tags/%7Bword%7D
func (s *TagsService) Get(t string) (*Tag, *Response, error) {
	u := fmt.Sprintf("tags/%s", t)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	tag := &Tag{}

	resp, err := s.client.Do(req, tag)
	if err != nil {
		return nil, resp, err
	}

	return tag, resp, err
}

// ListVideo lists the video for an tag.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/tags/%7Bword%7D/videos
func (s *TagsService) ListVideo(t string, opt *ListVideoOptions) ([]*Video, *Response, error) {
	u := fmt.Sprintf("tags/%s/videos", t)
	videos, resp, err := listVideo(s.client, u, opt)

	return videos, resp, err
}
