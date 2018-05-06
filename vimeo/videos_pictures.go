package vimeo

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

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
	Link        string         `json:"link,omitempty"`
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

// GetID returns the numeric identifier (ID) of the video.
func (v Pictures) GetID() int {
	l := strings.SplitN(v.URI, "/", -1)
	ID, _ := strconv.Atoi(l[len(l)-1])
	return ID
}

// ListPictures lists thumbnails.
//
// https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/pictures
func (s *VideosService) ListPictures(vid int, opt ...CallOption) ([]*Pictures, *Response, error) {
	u, err := addOptions(fmt.Sprintf("videos/%d/pictures", vid), opt...)
	if err != nil {
		return nil, nil, err
	}

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
func (s *VideosService) GetPictures(vid int, pid int, opt ...CallOption) (*Pictures, *Response, error) {
	u, err := addOptions(fmt.Sprintf("videos/%d/pictures/%d", vid, pid), opt...)
	if err != nil {
		return nil, nil, err
	}

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
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/pictures/%7Bpicture_id%7D
func (s *VideosService) DeletePictures(vid int, pid int) (*Response, error) {
	u := fmt.Sprintf("videos/%d/pictures/%d", vid, pid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UploadPicture shortcut upload picture file.
func (s *VideosService) UploadPicture(vid int, r *PicturesRequest, file *os.File) (*Pictures, *Response, error) {
	pictures, _, err := s.CreatePictures(vid, r)
	if err != nil {
		return nil, nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, nil, err
	}

	if stat.IsDir() {
		return nil, nil, errors.New("the video file can't be a directory")
	}

	req, err := http.NewRequest("PUT", pictures.Link, file)
	if err != nil {
		return nil, nil, err
	}

	_, err = s.client.Do(req, nil)
	if err != nil {
		return nil, nil, err
	}

	pictures, resp, err := s.GetPictures(vid, pictures.GetID())
	if err != nil {
		return nil, nil, err
	}

	return pictures, resp, err
}
