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

// ListPreset method returns all the embed presets that belong to the specified user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/embed-presets#get_embed_presets
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

// GetPreset method returns a single embed preset that belongs to the specified user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/embed-presets#get_embed_preset
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

// PresetListVideo method edits an embed present belonging to the specified user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/embed-presets#edit_embed_preset
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
