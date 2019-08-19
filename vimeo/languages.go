package vimeo

// LanguagesService handles communication with the languages related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#languages
type LanguagesService service

type languageList struct {
	Data []*Language `json:"data"`
	pagination
}

// Language represents a language.
type Language struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// List method returns all the video languages that Vimeo supports.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#get_languages
func (s *LanguagesService) List(opt ...CallOption) ([]*Language, *Response, error) {
	u, err := addOptions("languages", opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	languages := &languageList{}

	resp, err := s.client.Do(req, languages)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(languages)

	return languages.Data, resp, err
}
