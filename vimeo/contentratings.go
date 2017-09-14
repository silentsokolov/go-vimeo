package vimeo

// ContentRatingsService handles communication with the content ratings related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/endpoints/contentratings
type ContentRatingsService service

type contentRatingList struct {
	Data []*ContentRating `json:"data"`
	pagination
}

// ContentRating represents a content rating.
type ContentRating struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// List the content rating.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/contentratings
func (s *ContentRatingsService) List(opt ...CallOption) ([]*ContentRating, *Response, error) {
	u, err := addOptions("contentratings", opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	ratings := &contentRatingList{}

	resp, err := s.client.Do(req, ratings)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(ratings)

	return ratings.Data, resp, err
}
