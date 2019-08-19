package vimeo

import "fmt"

type dataListTextTrack struct {
	Data []*TextTrack `json:"data,omitempty"`
	pagination
}

// TextTrack represents a text track.
// TODO: Need full object.
type TextTrack struct {
	URI  string `json:"uri,omitempty"`
	Name string `json:"name,omitempty"`
}

// TextTrackRequest represents a request to create/edit text track.
type TextTrackRequest struct {
	Active   bool   `json:"role"`
	Type     string `json:"type,omitempty"`
	Language string `json:"language,omitempty"`
	Name     string `json:"name,omitempty"`
}

// ListTextTrack method returns all the text tracks of the specified video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#get_text_tracks
func (s *VideosService) ListTextTrack(vid int, opt ...CallOption) ([]*TextTrack, *Response, error) {
	u, err := addOptions(fmt.Sprintf("/videos/%d/texttracks", vid), opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	texttracks := &dataListTextTrack{}

	resp, err := s.client.Do(req, texttracks)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(texttracks)

	return texttracks.Data, resp, err
}

// AddTextTrack method adds a text track to the specified video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#create_text_track
func (s *VideosService) AddTextTrack(vid int, r *TextTrackRequest) (*TextTrack, *Response, error) {
	u := fmt.Sprintf("/videos/%d/texttracks", vid)
	req, err := s.client.NewRequest("POST", u, r)
	if err != nil {
		return nil, nil, err
	}

	textTrack := &TextTrack{}

	resp, err := s.client.Do(req, textTrack)
	if err != nil {
		return nil, resp, err
	}

	return textTrack, resp, nil
}

// GetTextTrack method returns a single text track from the specified video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#get_text_track
func (s *VideosService) GetTextTrack(vid int, tid int, opt ...CallOption) (*TextTrack, *Response, error) {
	u, err := addOptions(fmt.Sprintf("videos/%d/texttracks/%d", vid, tid), opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	textTrack := &TextTrack{}

	resp, err := s.client.Do(req, textTrack)
	if err != nil {
		return nil, resp, err
	}

	return textTrack, resp, err
}

// EditTextTrack method edits the specified text track.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#edit_text_track
func (s *VideosService) EditTextTrack(vid int, tid int, r *TextTrackRequest) (*TextTrack, *Response, error) {
	u := fmt.Sprintf("videos/%d/texttracks/%d", vid, tid)
	req, err := s.client.NewRequest("PATCH", u, r)
	if err != nil {
		return nil, nil, err
	}

	textTrack := &TextTrack{}
	resp, err := s.client.Do(req, textTrack)
	if err != nil {
		return nil, resp, err
	}

	return textTrack, resp, nil
}

// DeleteTextTrack method deletes the specified text track from a video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#delete_text_track
func (s *VideosService) DeleteTextTrack(vid int, tid int) (*Response, error) {
	u := fmt.Sprintf("videos/%d/texttracks/%d", vid, tid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
