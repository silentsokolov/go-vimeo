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
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos
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

// Download internal object provides access to video information for download
type Download struct {
	Quality     string    `json:"quality"`
	Type        string    `json:"type"`
	Width       int       `json:"width"`
	Height      int       `json:"height"`
	Expires     time.Time `json:"expires"`
	Link        string    `json:"link"`
	CreatedTime time.Time `json:"created_time"`
	Fps         int       `json:"fps"`
	Size        int       `json:"size"`
	Md5         string    `json:"md5"`
	PublicName  string    `json:"public_name"`
	SizeShort   string    `json:"size_short"`
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

// TransCode represents a request to upload video.
type TransCode struct {
	Status string `json:"status,omitempty"`
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
	Download      []*Download   `json:"download,omitempty"`
	App           *App          `json:"app,omitempty"`
	Status        string        `json:"status,omitempty"`
	ResourceKey   string        `json:"resource_key,omitempty"`
	EmbedPresets  *EmbedPresets `json:"embed_presets,omitempty"`
	Upload        *Upload       `json:"upload,omitempty"`
	TransCode     *TransCode    `json:"transcode,omitempty"`
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
	Active bool `json:"name,omitempty"`
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

func getUploadVideo(c *Client, method string, uri string, reqUpload *UploadVideoRequest) (*Video, *Response, error) { // nolint: unparam
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

// List method returns all the videos that match custom search criteria.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#search_videos
func (s *VideosService) List(opt ...CallOption) ([]*Video, *Response, error) {
	videos, resp, err := listVideo(s.client, "videos", opt...)

	return videos, resp, err
}

// Get method returns a single video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#get_video
func (s *VideosService) Get(vid int, opt ...CallOption) (*Video, *Response, error) {
	u := fmt.Sprintf("videos/%d", vid)
	video, resp, err := getVideo(s.client, u, opt...)

	return video, resp, err
}

// Edit method edits the specified video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#edit_video
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

// Delete method deletes the specified video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#delete_video
func (s *VideosService) Delete(vid int) (*Response, error) {
	u := fmt.Sprintf("videos/%d", vid)
	resp, err := deleteVideo(s.client, u)

	return resp, err
}

// ListCategory method gets all the categories that contain a particular video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/categories#get_video_categories
func (s *VideosService) ListCategory(vid int, opt ...CallOption) ([]*Category, *Response, error) {
	u := fmt.Sprintf("videos/%d/categories", vid)
	catogories, resp, err := listCategory(s.client, u, opt...)

	return catogories, resp, err
}

// LikeList method gets all the users who have liked a particular video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/likes#get_video_likes
func (s *VideosService) LikeList(vid int, opt ...CallOption) ([]*User, *Response, error) {
	u := fmt.Sprintf("videos/%d/likes", vid)
	users, resp, err := listUser(s.client, u, opt...)

	return users, resp, err
}

// GetPreset method determines whether the specified video uses a particular embed preset.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/embed-presets#get_video_embed_preset
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

// AssignPreset method assigns an embed preset to the specified video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/embed-presets#add_video_embed_preset
func (s *VideosService) AssignPreset(vid int, p int) (*Response, error) {
	u := fmt.Sprintf("videos/%d/presets/%d", vid, p)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnassignPreset method removes the embed preset from the specified video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/embed-presets#delete_video_embed_preset
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

// ListDomain method returns all the domains on the specified video's whitelist.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#get_video_privacy_domains
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

// AllowDomain method adds the specified domain to a video's whitelist.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#add_video_privacy_domain
func (s *VideosService) AllowDomain(vid int, d string) (*Response, error) {
	u := fmt.Sprintf("videos/%d/privacy/domains/%s", vid, d)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DisallowDomain method removes the specified domain from a video's whitelist.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#delete_video_privacy_domain
func (s *VideosService) DisallowDomain(vid int, d string) (*Response, error) {
	u := fmt.Sprintf("videos/%d/privacy/domains/%s", vid, d)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListUser method returns all the users who have access to the specified private video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#get_video_privacy_users
func (s *VideosService) ListUser(vid int, opt ...CallOption) ([]*User, *Response, error) {
	u := fmt.Sprintf("videos/%d/privacy/users", vid)
	users, resp, err := listUser(s.client, u, opt...)

	return users, resp, err
}

// AllowUsers method gives multiple users permission to view the specified private video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#add_video_privacy_users
func (s *VideosService) AllowUsers(vid int) (*Response, error) {
	u := fmt.Sprintf("videos/%d/privacy/users", vid)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// AllowUser method gives a single user permission to view the specified private video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#add_video_privacy_user
func (s *VideosService) AllowUser(vid int, uid string) (*Response, error) {
	u := fmt.Sprintf("videos/%d/privacy/users/%s", vid, uid)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// DisallowUser method prevents a user from being able to view the specified private video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#delete_video_privacy_user
func (s *VideosService) DisallowUser(vid int, uid string) (*Response, error) {
	u := fmt.Sprintf("videos/%d/privacy/users/%s", vid, uid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListTag method returns all the tags associated with a video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#get_video_tags
func (s *VideosService) ListTag(vid int, opt ...CallOption) ([]*Tag, *Response, error) {
	u := fmt.Sprintf("videos/%d/tags", vid)
	tags, resp, err := listTag(s.client, u, opt...)

	return tags, resp, err
}

// GetTag method determines whether a particular tag has been added to a video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#check_video_for_tag
func (s *VideosService) GetTag(vid int, t string, opt ...CallOption) (*Tag, *Response, error) {
	u := fmt.Sprintf("videos/%d/tags/%s", vid, t)
	tag, resp, err := getTag(s.client, u, opt...)

	return tag, resp, err
}

// AssignTag method adds a single tag to the specified video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#add_video_tag
func (s *VideosService) AssignTag(vid int, t string) (*Response, error) {
	u := fmt.Sprintf("videos/%d/tags/%s", vid, t)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnassignTag method removes the specified tag from a video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#delete_video_tag
func (s *VideosService) UnassignTag(vid int, t string) (*Response, error) {
	u := fmt.Sprintf("videos/%d/tags/%s", vid, t)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListRelatedVideo method returns all the related videos of a particular video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#get_related_videos
func (s *VideosService) ListRelatedVideo(vid int, opt ...CallOption) ([]*Video, *Response, error) {
	u := fmt.Sprintf("videos/%d/videos", vid)
	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}

// ReplaceFile method adds a version to the specified video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#create_video_version
func (s *VideosService) ReplaceFile(vid int, file *os.File) (*Video, *Response, error) {
	u := fmt.Sprintf("videos/%d/versions", vid)
	video, resp, err := uploadVideo(s.client, "POST", u, file)

	return video, resp, err
}
