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

// ListPortfolio lists the portfolio for user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/portfolios
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

// GetProtfolio get portfolio by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/portfolios/%7Bportfolio_id%7D
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

// ProtfolioListVideo lists the video for an portfolio.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/portfolios/%7Bportfolio_id%7D/videos
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

// ProtfolioGetVideo get specific video by portfolio name and video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/portfolios/%7Bportfolio_id%7D/videos/%7Bvideo_id%7D
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

// ProtfolioAddVideo add one video.
//
// Vimeo API docs: hhttps://developer.vimeo.com/api/playground/users/%7Buser_id%7D/portfolios/%7Bportfolio_id%7D/videos/%7Bvideo_id%7D
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

// ProtfolioDeleteVideo delete specific video by portfolio name and video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/portfolios/%7Bportfolio_id%7D/videos/%7Bvideo_id%7D
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
