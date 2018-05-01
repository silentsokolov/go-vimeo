package vimeo

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// VideosService handles communication with the videos related
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

// File internal object provides access to video file information
type File struct {
	Quality     string    `json:"quality,omitempty"`
	Type        string    `json:"type,omitempty"`
	Width       int       `json:"width,omitempty"`
	Height      int       `json:"height,omitempty"`
	Link        string    `json:"link,omitempty"`
	CreatedTime time.Time `json:"created_time,omitempty"`
	FPS         float32   `json:"fps,omitempty"`
	Size        int       `json:"size,omitempty"`
	MD5         string    `json:"md5,omitempty"`
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

// Upload represents a request to upload video.
type Upload struct {
	Status      string `json:"status,omitempty"`
	UploadLink  string `json:"upload_link,omitempty"`
	RedirectURL string `json:"redirect_url,omitempty"`
	Link        string `json:"link,omitempty"`
	Rorm        string `json:"form,omitempty"`
	Approach    string `json:"approach,omitempty"`
	Size        int64  `json:"size,omitempty"`
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
	Categories    []*Category   `json:"categories,omitempty"`
	User          *User         `json:"user,omitempty"`
	Files         []*File       `json:"files,omitempty"`
	App           *App          `json:"app,omitempty"`
	Status        string        `json:"status,omitempty"`
	ResourceKey   string        `json:"resource_key,omitempty"`
	EmbedPresets  *EmbedPresets `json:"embed_presets,omitempty"`
	Upload        *Upload       `json:"upload,omitempty"`
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

// RatingsRequest a request to edit an embed settings.
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

// ReviewPageRequest represents a request to edit an video.
type ReviewPageRequest struct {
	Active bool `json:"name,active"`
}

// VideoRequest represents a request to edit an video.
type VideoRequest struct {
	Name          string             `json:"name,omitempty"`
	Description   string             `json:"description,omitempty"`
	License       string             `json:"license,omitempty"`
	Privacy       *Privacy           `json:"privacy,omitempty"`
	Password      string             `json:"password,omitempty"`
	Locale        string             `json:"locale,omitempty"`
	ContentRating []string           `json:"content_rating,omitempty"`
	Embed         *EmbedRequest      `json:"embed,omitempty"`
	ReviewPage    *ReviewPageRequest `json:"review_page,omitempty"`
}

// GetID returns the numeric identifier (ID) of the video.
func (v Video) GetID() int {
	l := strings.SplitN(v.URI, "/", -1)
	ID, _ := strconv.Atoi(l[len(l)-1])
	return ID
}

// UploadVideoRequest specifies the optional parameters to the
// uploadVideo method.
type UploadVideoRequest struct {
	FileName string  `json:"file_name"`
	Upload   *Upload `json:"upload,omitempty"`
}

func listVideo(c *Client, url string, opt ...CallOption) ([]*Video, *Response, error) {
	u, err := addOptions(url, opt...)
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

func getVideo(c *Client, url string, opt ...CallOption) (*Video, *Response, error) {
	u, err := addOptions(url, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
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

func getUploadVideo(c *Client, method string, uri string, reqUpload *UploadVideoRequest) (*Video, *Response, error) {
	req, err := c.NewRequest(method, uri, reqUpload)
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

func uploadVideo(c *Client, method string, url string, file *os.File) (*Video, *Response, error) {
	if c.Config.Uploader == nil {
		return nil, nil, errors.New("uploader can't be nil if you need upload video")
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, nil, err
	}

	if stat.IsDir() {
		return nil, nil, errors.New("the video file can't be a directory")
	}

	reqUpload := &UploadVideoRequest{
		FileName: file.Name(),
		Upload: &Upload{
			Approach: "tus",
			Size:     stat.Size(),
		},
	}

	video, _, err := getUploadVideo(c, method, url, reqUpload)
	if err != nil {
		return nil, nil, err
	}

	err = c.Config.Uploader.UploadFromFile(c, video.Upload.UploadLink, file)
	if err != nil {
		return nil, nil, err
	}

	u := fmt.Sprintf("videos/%d", video.GetID())
	completeVideo, resp, err := getVideo(c, u)

	return completeVideo, resp, err
}

func uploadVideoByURL(c *Client, uri, videoURL string) (*Video, *Response, error) {
	reqUpload := &UploadVideoRequest{
		Upload: &Upload{
			Approach: "pull",
			Link:     videoURL,
		},
	}

	req, err := c.NewRequest("POST", uri, reqUpload)
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

func addVideo(c *Client, url string) (*Video, *Response, error) {
	req, err := c.NewRequest("PUT", url, nil)
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

// List lists the videos.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos
func (s *VideosService) List(opt ...CallOption) ([]*Video, *Response, error) {
	videos, resp, err := listVideo(s.client, "videos", opt...)

	return videos, resp, err
}

// Get specific video by ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D
func (s *VideosService) Get(vid int, opt ...CallOption) (*Video, *Response, error) {
	u := fmt.Sprintf("videos/%d", vid)
	video, resp, err := getVideo(s.client, u, opt...)

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
func (s *VideosService) ListCategory(vid int, opt ...CallOption) ([]*Category, *Response, error) {
	u := fmt.Sprintf("videos/%d/categories", vid)
	catogories, resp, err := listCategory(s.client, u, opt...)

	return catogories, resp, err
}

// LikeList lists users who liked this video.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/likes
func (s *VideosService) LikeList(vid int, opt ...CallOption) ([]*User, *Response, error) {
	u := fmt.Sprintf("videos/%d/likes", vid)
	users, resp, err := listUser(s.client, u, opt...)

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
func (s *VideosService) ListDomain(vid int, opt ...CallOption) ([]*Domain, *Response, error) {
	u, err := addOptions(fmt.Sprintf("videos/%d/privacy/domains", vid), opt...)
	if err != nil {
		return nil, nil, err
	}

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

// ListUser list the all allowed users
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/privacy/users
func (s *VideosService) ListUser(vid int, opt ...CallOption) ([]*User, *Response, error) {
	u := fmt.Sprintf("videos/%d/privacy/users", vid)
	users, resp, err := listUser(s.client, u, opt...)

	return users, resp, err
}

// AllowUsers allow users to view this video.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/privacy/users
func (s *VideosService) AllowUsers(vid int) (*Response, error) {
	u := fmt.Sprintf("videos/%d/privacy/users", vid)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// AllowUser allow users to view this video.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/privacy/users/%7Buser_id%7D
func (s *VideosService) AllowUser(vid int, uid string) (*Response, error) {
	u := fmt.Sprintf("videos/%d/privacy/users/%s", vid, uid)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DisallowUser disallow user from viewing this video.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/privacy/users/%7Buser_id%7D
func (s *VideosService) DisallowUser(vid int, uid string) (*Response, error) {
	u := fmt.Sprintf("videos/%d/privacy/users/%s", vid, uid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListTag list a video's tags
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/tags
func (s *VideosService) ListTag(vid int, opt ...CallOption) ([]*Tag, *Response, error) {
	u := fmt.Sprintf("videos/%d/tags", vid)
	tags, resp, err := listTag(s.client, u, opt...)

	return tags, resp, err
}

// GetTag specific tag by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/tags/%7Bword%7D
func (s *VideosService) GetTag(vid int, t string, opt ...CallOption) (*Tag, *Response, error) {
	u := fmt.Sprintf("videos/%d/tags/%s", vid, t)
	tag, resp, err := getTag(s.client, u, opt...)

	return tag, resp, err
}

// AssignTag specific tag by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/tags/%7Bword%7D
func (s *VideosService) AssignTag(vid int, t string) (*Response, error) {
	u := fmt.Sprintf("videos/%d/tags/%s", vid, t)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnassignTag specific tag by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/tags/%7Bword%7D
func (s *VideosService) UnassignTag(vid int, t string) (*Response, error) {
	u := fmt.Sprintf("videos/%d/tags/%s", vid, t)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListRelatedVideo lists the related video.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/videos
func (s *VideosService) ListRelatedVideo(vid int, opt ...CallOption) ([]*Video, *Response, error) {
	u := fmt.Sprintf("videos/%d/videos", vid)
	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}

// ReplaceFile upload video file/replace video file.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/files
func (s *VideosService) ReplaceFile(vid int, file *os.File) (*Video, *Response, error) {
	u := fmt.Sprintf("videos/%d/versions", vid)
	video, resp, err := uploadVideo(s.client, "POST", u, file)

	return video, resp, err
}
