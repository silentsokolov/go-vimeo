package vimeo

import "fmt"

// TagsService handles communication with the tag related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/endpoints/tags
type TagsService service

type dataListTag struct {
	Data []*Tag `json:"data"`
	pagination
}

// Tag represents a tag.
type Tag struct {
	URI         string `json:"uri,omitempty"`
	Name        string `json:"name,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Canonical   string `json:"canonical,omitempty"`
	ResourceKey string `json:"resource_key,omitempty"`
}

func listTag(c *Client, url string) ([]*Tag, *Response, error) {
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	categories := &dataListTag{}

	resp, err := c.Do(req, categories)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(categories)

	return categories.Data, resp, err
}

func getTag(c *Client, url string) (*Tag, *Response, error) {
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	tag := &Tag{}

	resp, err := c.Do(req, tag)
	if err != nil {
		return nil, resp, err
	}

	return tag, resp, err
}

// Get specific tag by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/tags/%7Bword%7D
func (s *TagsService) Get(t string) (*Tag, *Response, error) {
	u := fmt.Sprintf("tags/%s", t)
	tag, resp, err := getTag(s.client, u)

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
