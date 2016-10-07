package vimeo

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// CategoriesService handles communication with the categories related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/endpoints/categories
type CategoriesService service

type dataListCategory struct {
	Data []*Category `json:"data"`
	pagination
}

// Category represents a category.
type Category struct {
	URI                   string         `json:"uri,omitempty"`
	Link                  string         `json:"link,omitempty"`
	Name                  string         `json:"name,omitempty"`
	TopLevel              bool           `json:"top_level"`
	Pictures              *Pictures      `json:"pictures,omitempty"`
	LastVideoFeaturedTime string         `json:"last_video_featured_time,omitempty"`
	Parent                string         `json:"parent,omitempty"`
	SubCategories         []*SubCategory `json:"subcategories,omitempty"`
	ResourceKey           string         `json:"resource_key,omitempty"`
}

// Pictures internal object provides access to pictures.
type Pictures struct {
	URI         string         `json:"uri,omitempty"`
	Active      bool           `json:"active"`
	Type        string         `json:"type,omitempty"`
	Sizes       []*PictureSize `json:"sizes,omitempty"`
	ResourceKey string         `json:"resource_key,omitempty"`
}

// PictureSize internal object provides access to picture size.
type PictureSize struct {
	Width              int    `json:"width,omitempty"`
	Height             int    `json:"height,omitempty"`
	Link               string `json:"link,omitempty"`
	LinkWithPlayButton string `json:"link_with_play_button,omitempty"`
}

// PictureSize internal object provides access to header pictures.
type Header struct {
	URI         string         `json:"uri,omitempty"`
	Active      bool           `json:"active"`
	Type        string         `json:"type,omitempty"`
	Sizes       []*PictureSize `json:"sizes,omitempty"`
	ResourceKey string         `json:"resource_key,omitempty"`
}

// SubCategory internal object provides access to subcategory in category.
type SubCategory struct {
	URI  string `json:"URI,omitempty"`
	Name string `json:"name,omitempty"`
	Link string `json:"link,omitempty"`
}

// ListCategoryOptions specifies the optional parameters to the
// CategoriesService.List method.
type ListCategoryOptions struct {
	ListOptions
}

// List the category.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/categories
func (s *CategoriesService) List(opt *ListCategoryOptions) ([]*Category, *Response, error) {
	u, err := addOptions("categories", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	categories := &dataListCategory{}

	resp, err := s.client.Do(req, categories)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(categories)

	return categories.Data, resp, err
}

// Get specific category by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/categories/%7Bcategory%7D
func (s *CategoriesService) Get(cat string) (*Category, *Response, error) {
	u := fmt.Sprintf("categories/%s", cat)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	category := &Category{}

	resp, err := s.client.Do(req, category)
	if err != nil {
		return nil, resp, err
	}

	return category, resp, err
}

// Privacy internal object provides access to privacy.
type Privacy struct {
	View     string `json:"view,omitempty"`
	Join     string `json:"join,omitempty"`
	Videos   string `json:"videos,omitempty"`
	Comment  string `json:"comment,omitempty"`
	Forums   string `json:"forums,omitempty"`
	Invite   string `json:"invite,omitempty"`
	Download bool   `json:"download"`
	Add      bool   `json:"add"`
}

// ListChannel lists the channel for an category.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/categories/%7Bcategory%7D/channels
func (s *CategoriesService) ListChannel(cat string, opt *ListChannelOptions) ([]*Channel, *Response, error) {
	u := fmt.Sprintf("categories/%s/channels", cat)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	channels, resp, err := listChannel(s.client, u, opt)

	return channels, resp, err
}

// ListGroup lists the group for an category.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/categories/%7Bcategory%7D/groups
func (s *CategoriesService) ListGroup(cat string, opt *ListGroupOptions) ([]*Group, *Response, error) {
	u := fmt.Sprintf("categories/%s/groups", cat)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	groups, resp, err := listGroup(s.client, u, opt)

	return groups, resp, err
}

type dataListVideo struct {
	Data []*Video `json:"data,omitempty"`
	pagination
}

// Embed internal object provides access to HTML embed code.
type Embed struct {
	HTML string `json:"html,omitempty"`
}

// Tag represents a tag.
type Tag struct {
	URI         string `json:"uri,omitempty"`
	Name        string `json:"name,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Canonical   string `json:"canonical,omitempty"`
	ResourceKey string `json:"resource_key,omitempty"`
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

// Logos internal object present settings.
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
	ListOptions
}

// ListVideo lists the video for an category.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/categories/%7Bcategory%7D/videos
func (s *CategoriesService) ListVideo(cat string, opt *ListVideoOptions) ([]*Video, *Response, error) {
	u := fmt.Sprintf("categories/%s/videos", cat)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	videos := &dataListVideo{}

	resp, err := s.client.Do(req, videos)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(videos)

	return videos.Data, resp, err
}

// GetVideo specific video by category name and video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/categories/%7Bcategory%7D/videos/%7Bvideo_id%7D
func (s *CategoriesService) GetVideo(cat string, vid int) (*Video, *Response, error) {
	u := fmt.Sprintf("categories/%s/videos/%d", cat, vid)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	video := &Video{}

	resp, err := s.client.Do(req, video)
	if err != nil {
		return nil, resp, err
	}

	return video, resp, err
}
