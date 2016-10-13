package vimeo

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// VideosService handles communication with the tag related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/endpoints/videos
type VideosService service

type dataListVideo struct {
	Data []*Video `json:"data,omitempty"`
	pagination
}

// Embed internal object provides access to HTML embed code.
type Embed struct {
	HTML string `json:"html,omitempty"`
}

// Stats internal object provides access to video statistic.
type Stats struct {
	Plays int `json:"plays,omitempty"`
}

// App internal object provides access to specific app.
type App struct {
	URI  string `json:"uri,omitempty"`
	Name string `json:"name,omitempty"`
}

// Buttons internal object embed settings.
type Buttons struct {
	Like       bool `json:"like"`
	WatchLater bool `json:"watchlater"`
	Share      bool `json:"share"`
	Embed      bool `json:"embed"`
	Vote       bool `json:"vote"`
	HD         bool `json:"HD"`
}

// Logos internal object embed settings.
type Logos struct {
	Vimeo        bool `json:"vimeo"`
	Custom       bool `json:"custom"`
	StickyCustom bool `json:"sticky_custom"`
}

// EmbedSettings internal object provides access to embed settings.
type EmbedSettings struct {
	Buttons                         *Buttons `json:"buttons,omitempty"`
	Logos                           *Logos   `json:"logos,omitempty"`
	Outro                           string   `json:"outro,omitempty"`
	Portrait                        string   `json:"portrait,omitempty"`
	Title                           string   `json:"title,omitempty"`
	ByLine                          string   `json:"byline,omitempty"`
	Badge                           bool     `json:"badge"`
	ByLineBadge                     bool     `json:"byline_badge"`
	CollectionsButton               bool     `json:"collections_button"`
	PlayBar                         bool     `json:"playbar"`
	Volume                          bool     `json:"volume"`
	FullscreenButton                bool     `json:"fullscreen_button"`
	ScalingButton                   bool     `json:"scaling_button"`
	Autoplay                        bool     `json:"autoplay"`
	Autopause                       bool     `json:"autopause"`
	Loop                            bool     `json:"loop"`
	Color                           string   `json:"color,omitempty"`
	Link                            bool     `json:"link"`
	OverlayEmailCapture             int      `json:"overlay_email_capture,omitempty"`
	OverlayEmailCaptureText         string   `json:"overlay_email_capture_text,omitempty"`
	OverlayEmailCaptureConfirmation string   `json:"overlay_email_capture_confirmation,omitempty"`
}

// EmbedPresets internal object present settings.
type EmbedPresets struct {
	URI      string         `json:"uri,omitempty"`
	Name     string         `json:"name,omitempty"`
	Settings *EmbedSettings `json:"settings,omitempty"`
	User     *User          `json:"user,omitempty"`
}

// Video represents a video.
type Video struct {
	URI           string        `json:"uri,omitempty"`
	Name          string        `json:"name,omitempty"`
	Description   string        `json:"description,omitempty"`
	Link          string        `json:"link,omitempty"`
	Duration      int           `json:"duration,omitempty"`
	Width         int           `json:"width,omitempty"`
	Height        int           `json:"height,omitempty"`
	Language      string        `json:"language,omitempty"`
	Embed         *Embed        `json:"embed,omitempty"`
	CreatedTime   time.Time     `json:"created_time,omitempty"`
	ModifiedTime  time.Time     `json:"modified_time,omitempty"`
	ReleaseTime   time.Time     `json:"release_time,omitempty"`
	ContentRating []string      `json:"content_rating,omitempty"`
	License       string        `json:"license,omitempty"`
	Privacy       *Privacy      `json:"privacy,omitempty"`
	Pictures      *Pictures     `json:"pictures,omitempty"`
	Tags          []*Tag        `json:"tags,omitempty"`
	Stats         *Stats        `json:"stats,omitempty"`
	User          *User         `json:"user,omitempty"`
	App           *App          `json:"app,omitempty"`
	Status        string        `json:"status,omitempty"`
	ResourceKey   string        `json:"resource_key,omitempty"`
	EmbedPresets  *EmbedPresets `json:"embed_presets,omitempty"`
}

// TitleRequest a request to edit an embed settings.
type TitleRequest struct {
	Owner    string `json:"owner,omitempty"`
	Portrait string `json:"portrait,omitempty"`
	Name     string `json:"name,omitempty"`
}

// RatingTVRequest a request to edit video.
type RatingTVRequest struct {
	Rating string `json:"rating,omitempty"`
	Reason string `json:"reason,omitempty"`
}

// RatingMPAARequest a request to edit video.
type RatingMPAARequest struct {
	Rating string `json:"rating,omitempty"`
	Reason string `json:"reason,omitempty"`
}

// TitleRequest a request to edit an embed settings.
type RatingsRequest struct {
	RatingTVRequest   string `json:"tv,omitempty"`
	RatingMPAARequest string `json:"mpaa,omitempty"`
}

// ExtraLinksRequest a request to edit video.
type ExtraLinksRequest struct {
	IMDB           string `json:"imdb,omitempty"`
	RottenTomatoes string `json:"rotten_tomatoes,omitempty"`
}

// EmbedRequest a request to edit an embed settings.
type EmbedRequest struct {
	Buttons                         *Buttons           `json:"buttons,omitempty"`
	Logos                           *Logos             `json:"logos,omitempty"`
	Outro                           string             `json:"outro,omitempty"`
	Portrait                        string             `json:"portrait,omitempty"`
	Title                           *TitleRequest      `json:"title,omitempty"`
	ByLine                          string             `json:"byline,omitempty"`
	Badge                           bool               `json:"badge"`
	ByLineBadge                     bool               `json:"byline_badge"`
	CollectionsButton               bool               `json:"collections_button"`
	PlayBar                         bool               `json:"playbar"`
	Volume                          bool               `json:"volume"`
	FullscreenButton                bool               `json:"fullscreen_button"`
	ScalingButton                   bool               `json:"scaling_button"`
	Autoplay                        bool               `json:"autoplay"`
	Autopause                       bool               `json:"autopause"`
	Loop                            bool               `json:"loop"`
	Color                           string             `json:"color,omitempty"`
	Link                            bool               `json:"link"`
	RatingsRequest                  *RatingsRequest    `json:"ratings,omitempty"`
	ExtraLinks                      *ExtraLinksRequest `json:"external_links,omitempty"`
	OverlayEmailCapture             int                `json:"overlay_email_capture,omitempty"`
	OverlayEmailCaptureText         string             `json:"overlay_email_capture_text,omitempty"`
	OverlayEmailCaptureConfirmation string             `json:"overlay_email_capture_confirmation,omitempty"`
}

// VideoRequest represents a request to edit an video.
type VideoRequest struct {
	Name          string        `json:"name,omitempty"`
	Description   string        `json:"description,omitempty"`
	License       string        `json:"license,omitempty"`
	Privacy       *Privacy      `json:"privacy,omitempty"`
	Password      string        `json:"password,omitempty"`
	ReviewLink    bool          `json:"review_link"`
	Locale        string        `json:"locale,omitempty"`
	ContentRating []string      `json:"content_rating,omitempty"`
	Embed         *EmbedRequest `json:"embed,omitempty"`
}

// GetID returns the numeric identifier (ID) of the video.
func (v Video) GetID() int {
	l := strings.SplitN(v.URI, "/", -1)
	ID, _ := strconv.Atoi(l[len(l)-1])
	return ID
}

// ListVideoOptions specifies the optional parameters to the
// CategoriesService.ListVideo method.
type ListVideoOptions struct {
	Query            string `url:"query,omitempty"`
	Filter           string `url:"filter,omitempty"`
	FilterEmbeddable string `url:"filter_embeddable,omitempty"`
	Sort             string `url:"sort,omitempty"`
	Direction        string `url:"direction,omitempty"`
	FilterPlayable   string `url:"filter_playable,omitempty"`
	ListOptions
}

func listVideo(c *Client, url string, opt *ListVideoOptions) ([]*Video, *Response, error) {
	u, err := addOptions(url, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	videos := &dataListVideo{}

	resp, err := c.Do(req, videos)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(videos)

	return videos.Data, resp, err
}

func getVideo(c *Client, url string) (*Video, *Response, error) {
	req, err := c.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	video := &Video{}

	resp, err := c.Do(req, video)
	if err != nil {
		return nil, resp, err
	}

	return video, resp, err
}

func deleteVideo(c *Client, url string) (*Response, error) {
	req, err := c.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req, nil)
}

// List lists the videos.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos
func (s *VideosService) List(opt *ListVideoOptions) ([]*Video, *Response, error) {
	videos, resp, err := listVideo(s.client, "videos", opt)

	return videos, resp, err
}

// Get specific video by ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D
func (s *VideosService) Get(vid int) (*Video, *Response, error) {
	u := fmt.Sprintf("videos/%d", vid)
	video, resp, err := getVideo(s.client, u)

	return video, resp, err
}

// Edit specific video by ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D
func (s *VideosService) Edit(vid int, r *VideoRequest) (*Video, *Response, error) {
	u := fmt.Sprintf("videos/%d", vid)
	req, err := s.client.NewRequest("PATCH", u, r)
	if err != nil {
		return nil, nil, err
	}

	video := &Video{}
	resp, err := s.client.Do(req, video)
	if err != nil {
		return nil, resp, err
	}

	return video, resp, nil
}

// Delete specific video by ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D
func (s *VideosService) Delete(vid int) (*Response, error) {
	u := fmt.Sprintf("videos/%d", vid)
	resp, err := deleteVideo(s.client, u)

	return resp, err
}

// ListCategory lists the video category.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/categories
func (s *VideosService) ListCategory(vid int, opt *ListCategoryOptions) ([]*Category, *Response, error) {
	u := fmt.Sprintf("videos/%d/categories", vid)
	catogories, resp, err := listCategory(s.client, u, opt)

	return catogories, resp, err
}

// LikeList lists users who liked this video.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/likes
func (s *VideosService) LikeList(vid int, opt *ListUserOptions) ([]*User, *Response, error) {
	u := fmt.Sprintf("videos/%d/likes", vid)
	users, resp, err := listUser(s.client, u, opt)

	return users, resp, err
}

// GetPreset get preset by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/presets/%7Bpreset_id%7D
func (s *VideosService) GetPreset(vid int, p int) (*Preset, *Response, error) {
	u := fmt.Sprintf("videos/%d/presets/%d", vid, p)
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

// AssignPreset embed preset by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/presets/%7Bpreset_id%7D
func (s *VideosService) AssignPreset(vid int, p int) (*Response, error) {
	u := fmt.Sprintf("videos/%d/presets/%d", vid, p)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnassignPreset embed preset by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/presets/%7Bpreset_id%7D
func (s *VideosService) UnassignPreset(vid int, p int) (*Response, error) {
	u := fmt.Sprintf("videos/%d/presets/%d", vid, p)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

type dataListDomain struct {
	Data []*Domain `json:"data,omitempty"`
	pagination
}

// Domain represents a domain.
type Domain struct {
	URI  string `json:"uri,omitempty"`
	Name string `json:"name,omitempty"`
}

// ListDomain lists the domains.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/privacy/domains
func (s *VideosService) ListDomain(vid int) ([]*Domain, *Response, error) {
	u := fmt.Sprintf("videos/%d/privacy/domains", vid)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	domains := &dataListDomain{}

	resp, err := s.client.Do(req, domains)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(domains)

	return domains.Data, resp, err
}

// AllowDomain embedding on a domain.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/privacy/domains/%7Bdomain%7D
func (s *VideosService) AllowDomain(vid int, d string) (*Response, error) {
	u := fmt.Sprintf("videos/%d/privacy/domains/%s", vid, d)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DisallowDomain embedding on a domain.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/privacy/domains/%7Bdomain%7D
func (s *VideosService) DisallowDomain(vid int, d string) (*Response, error) {
	u := fmt.Sprintf("videos/%d/privacy/domains/%s", vid, d)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
