package vimeo

import (
	"fmt"
	"time"
)

type dataListPortfolio struct {
	Data []*Portfolio `json:"data,omitempty"`
	pagination
}

// Portfolio represents a portfolio.
type Portfolio struct {
	URI          string    `json:"uri,omitempty"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	Link         string    `json:"link,omitempty"`
	CreatedTime  time.Time `json:"created_time,omitempty"`
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	Sort         string    `json:"sort,omitempty"`
}

// ListPortfolio method gets all the specified user's portfolios.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/portfolios#get_portfolios
func (s *UsersService) ListPortfolio(uid string, opt ...CallOption) ([]*Portfolio, *Response, error) {
	var u string
	if uid == "" {
		u = "me/portfolios"
	} else {
		u = fmt.Sprintf("users/%s/portfolios", uid)
	}

	u, err := addOptions(u, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	portfolio := &dataListPortfolio{}

	resp, err := s.client.Do(req, portfolio)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(portfolio)

	return portfolio.Data, resp, err
}

// GetProtfolio method gets a single portfolio from the specified user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/portfolios#get_portfolio
func (s *UsersService) GetProtfolio(uid string, p string, opt ...CallOption) (*Portfolio, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/portfolios/%s", p)
	} else {
		u = fmt.Sprintf("users/%s/portfolios/%s", uid, p)
	}

	u, err := addOptions(u, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	portf := &Portfolio{}

	resp, err := s.client.Do(req, portf)
	if err != nil {
		return nil, resp, err
	}

	return portf, resp, err
}

// ProtfolioListVideo method gets all the videos from the specified portfolio.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/portfolios#get_portfolio_videos
func (s *UsersService) ProtfolioListVideo(uid string, p string, opt ...CallOption) ([]*Video, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/portfolios/%s/videos", p)
	} else {
		u = fmt.Sprintf("users/%s/portfolios/%s/videos", uid, p)
	}

	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}

// ProtfolioGetVideo method gets a single video from the specified portfolio.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/portfolios#get_portfolio_video
func (s *UsersService) ProtfolioGetVideo(uid string, p string, vid int, opt ...CallOption) (*Video, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/portfolios/%s/videos/%d", p, vid)
	} else {
		u = fmt.Sprintf("users/%s/portfolios/%s/videos/%d", uid, p, vid)
	}

	video, resp, err := getVideo(s.client, u, opt...)

	return video, resp, err
}

// ProtfolioAddVideo method adds a video to the specified portfolio.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/portfolios#add_video_to_portfolio
func (s *UsersService) ProtfolioAddVideo(uid string, p string, vid int) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/portfolios/%s/videos/%d", p, vid)
	} else {
		u = fmt.Sprintf("users/%s/portfolios/%s/videos/%d", uid, p, vid)
	}

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ProtfolioDeleteVideo method removes a video from the specified portfolio.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/portfolios#delete_video_from_portfolio
func (s *UsersService) ProtfolioDeleteVideo(uid string, p string, vid int) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/portfolios/%s/videos/%d", p, vid)
	} else {
		u = fmt.Sprintf("users/%s/portfolios/%s/videos/%d", uid, p, vid)
	}

	resp, err := deleteVideo(s.client, u)

	return resp, err
}
