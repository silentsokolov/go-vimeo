package vimeo

import "fmt"

type dataListPreset struct {
	Data []*Preset `json:"data,omitempty"`
	pagination
}

// Preset represents a preset.
type Preset struct {
	URI  string `json:"uri,omitempty"`
	Name string `json:"name,omitempty"`
}

// ListPreset lists the preset for an current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/presets
func (s *UsersService) ListPreset(uid string, opt ...CallOption) ([]*Preset, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/presets")
	} else {
		u = fmt.Sprintf("users/%s/presets", uid)
	}

	u, err := addOptions(u, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	preset := &dataListPreset{}

	resp, err := s.client.Do(req, preset)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(preset)

	return preset.Data, resp, err
}

// GetPreset get preset by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/presets/%7Bpreset_id%7D
func (s *UsersService) GetPreset(uid string, p int, opt ...CallOption) (*Preset, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/presets/%d", p)
	} else {
		u = fmt.Sprintf("users/%s/presets/%d", uid, p)
	}

	u, err := addOptions(u, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	portf := &Preset{}

	resp, err := s.client.Do(req, portf)
	if err != nil {
		return nil, resp, err
	}

	return portf, resp, err
}

// PresetListVideo lists the preset for an preset.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/presets/%7Bpreset_id%7D/videos
func (s *UsersService) PresetListVideo(uid string, p int, opt ...CallOption) ([]*Video, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/presets/%d/videos", p)
	} else {
		u = fmt.Sprintf("users/%s/presets/%d/videos", uid, p)
	}

	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}
