package vimeo

// CreativeCommonsService handles communication with the creative commons related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/endpoints/creativecommons
type CreativeCommonsService service

type creativeCommonList struct {
	Data []*CreativeCommon `json:"data"`
	pagination
}

// CreativeCommon represents a creative common.
type CreativeCommon struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// List the creative common.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/creativecommons
func (s *CreativeCommonsService) List(opt ...CallOption) ([]*CreativeCommon, *Response, error) {
	u, err := addOptions("creativecommons", opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	commons := &creativeCommonList{}

	resp, err := s.client.Do(req, commons)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(commons)

	return commons.Data, resp, err
}
