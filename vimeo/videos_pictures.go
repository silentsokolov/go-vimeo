package vimeo

import "fmt"

type dataListPictures struct {
	Data []*Pictures `json:"data,omitempty"`
	pagination
}

// Pictures internal object provides access to pictures.
type Pictures struct {
	URI         string         `json:"uri,omitempty"`
	Active      bool           `json:"active"`
	Type        string         `json:"type,omitempty"`
	Sizes       []*PictureSize `json:"sizes,omitempty"`
	ResourceKey string         `json:"resource_key,omitempty"`
}

// PictureSize internal object provides access to picture size.
type PictureSize struct {
	Width              int    `json:"width,omitempty"`
	Height             int    `json:"height,omitempty"`
	Link               string `json:"link,omitempty"`
	LinkWithPlayButton string `json:"link_with_play_button,omitempty"`
}

// Header internal object provides access to header pictures.
type Header struct {
	URI         string         `json:"uri,omitempty"`
	Active      bool           `json:"active"`
	Type        string         `json:"type,omitempty"`
	Sizes       []*PictureSize `json:"sizes,omitempty"`
	ResourceKey string         `json:"resource_key,omitempty"`
}

// PicturesRequest represents a request to create/edit an pictures.
type PicturesRequest struct {
	Time   float32 `json:"time,omitempty"`
	Active bool    `json:"active,omitempty"`
}

// ListPictures lists thumbnails.
//
// https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/pictures
func (s *VideosService) ListPictures(vid int) ([]*Pictures, *Response, error) {
	u := fmt.Sprintf("videos/%d/pictures", vid)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	pictures := &dataListPictures{}

	resp, err := s.client.Do(req, pictures)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(pictures)

	return pictures.Data, resp, err
}

// CreatePictures create a thumbnail.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/pictures
func (s *VideosService) CreatePictures(vid int, r *PicturesRequest) (*Pictures, *Response, error) {
	u := fmt.Sprintf("videos/%d/pictures", vid)
	req, err := s.client.NewRequest("POST", u, r)
	if err != nil {
		return nil, nil, err
	}

	pictures := &Pictures{}
	resp, err := s.client.Do(req, pictures)
	if err != nil {
		return nil, resp, err
	}

	return pictures, resp, nil
}

// GetPictures get one thumbnail.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/pictures
func (s *VideosService) GetPictures(vid int, pid int) (*Pictures, *Response, error) {
	u := fmt.Sprintf("videos/%d/pictures/%d", vid, pid)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	pictures := &Pictures{}

	resp, err := s.client.Do(req, pictures)
	if err != nil {
		return nil, resp, err
	}

	return pictures, resp, err
}

// EditPictures edit specific pictures by ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/pictures/%7Bpicture_id%7D
func (s *VideosService) EditPictures(vid int, pid int, r *PicturesRequest) (*Pictures, *Response, error) {
	u := fmt.Sprintf("videos/%d/pictures/%d", vid, pid)
	req, err := s.client.NewRequest("PATCH", u, r)
	if err != nil {
		return nil, nil, err
	}

	pictures := &Pictures{}
	resp, err := s.client.Do(req, pictures)
	if err != nil {
		return nil, resp, err
	}

	return pictures, resp, nil
}

// DeletePictures delete specific pictures by ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/comments/%7Bcomment_id%7D
func (s *VideosService) DeletePictures(vid int, pid int) (*Response, error) {
	u := fmt.Sprintf("videos/%d/pictures/%d", vid, pid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
