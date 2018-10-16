package vimeo

import "fmt"

type dataListCredit struct {
	Data []*Credit `json:"data,omitempty"`
	pagination
}

// Credit represents a creadit.
type Credit struct {
	URI   string `json:"uri,omitempty"`
	Name  string `json:"name,omitempty"`
	Role  string `json:"role,omitempty"`
	User  *User  `json:"user,omitempty"`
	Video *Video `json:"video,omitempty"`
}

// CreditRequest represents a request to create/edit an creadit.
type CreditRequest struct {
	Role    string `json:"role,omitempty"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	UserURI string `json:"user_uri,omitempty"`
}

// ListCredit lists the credits.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/credits
func (s *VideosService) ListCredit(vid int, opt ...CallOption) ([]*Credit, *Response, error) {
	u, err := addOptions(fmt.Sprintf("videos/%d/credits", vid), opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	credits := &dataListCredit{}

	resp, err := s.client.Do(req, credits)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(credits)

	return credits.Data, resp, err
}

// AddCredit add credit.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/credits
func (s *VideosService) AddCredit(vid int, r *CreditRequest) (*Credit, *Response, error) {
	u := fmt.Sprintf("videos/%d/credits", vid)
	req, err := s.client.NewRequest("POST", u, r)
	if err != nil {
		return nil, nil, err
	}

	credit := &Credit{}
	resp, err := s.client.Do(req, credit)
	if err != nil {
		return nil, resp, err
	}

	return credit, resp, nil
}

// GetCredit get specific credit by ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/credits/%7Bcredit_id%7D
func (s *VideosService) GetCredit(vid int, cid int, opt ...CallOption) (*Credit, *Response, error) {
	u, err := addOptions(fmt.Sprintf("videos/%d/credits/%d", vid, cid), opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	credit := &Credit{}

	resp, err := s.client.Do(req, credit)
	if err != nil {
		return nil, resp, err
	}

	return credit, resp, err
}

// EditCredit edit specific credit by ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/credits/%7Bcredit_id%7D
func (s *VideosService) EditCredit(vid int, cid int, r *CreditRequest) (*Credit, *Response, error) {
	u := fmt.Sprintf("videos/%d/credits/%d", vid, cid)
	req, err := s.client.NewRequest("PATCH", u, r)
	if err != nil {
		return nil, nil, err
	}

	credit := &Credit{}
	resp, err := s.client.Do(req, credit)
	if err != nil {
		return nil, resp, err
	}

	return credit, resp, nil
}

// DeleteCredit delete specific credit by ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/credits/%7Bcredit_id%7D
func (s *VideosService) DeleteCredit(vid int, cid int) (*Response, error) {
	u := fmt.Sprintf("videos/%d/credits/%d", vid, cid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
