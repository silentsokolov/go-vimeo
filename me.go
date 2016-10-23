package vimeo

import (
	"fmt"
	"time"
)

// MeService handles communication with the me related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/endpoints/me
type MeService service

// Get show current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me
func (s *MeService) Get() (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", "me", nil)
	if err != nil {
		return nil, nil, err
	}

	user := &User{}

	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, err
}

// Edit current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me
func (s *MeService) Edit(r *UserRequest) (*User, *Response, error) {
	req, err := s.client.NewRequest("PATCH", "me", r)
	if err != nil {
		return nil, nil, err
	}

	user := &User{}
	resp, err := s.client.Do(req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// ListAlbum lists the album for an current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/albums
func (s *MeService) ListAlbum(opt *ListAlbumOptions) ([]*Album, *Response, error) {
	u, err := addOptions("me/albums", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	albums := &dataListAlbum{}

	resp, err := s.client.Do(req, albums)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(albums)

	return albums.Data, resp, err
}

// CreateAlbum a new album.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/albums
func (s *MeService) CreateAlbum(r *AlbumRequest) (*Album, *Response, error) {
	req, err := s.client.NewRequest("POST", "me/albums", r)
	if err != nil {
		return nil, nil, err
	}

	album := &Album{}
	resp, err := s.client.Do(req, album)
	if err != nil {
		return nil, resp, err
	}

	return album, resp, nil
}

// GetAlbum specific album by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/albums/%7Balbum_id%7D
func (s *MeService) GetAlbum(ab string) (*Album, *Response, error) {
	u := fmt.Sprintf("me/albums/%s", ab)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	album := &Album{}

	resp, err := s.client.Do(req, album)
	if err != nil {
		return nil, resp, err
	}

	return album, resp, err
}

// EditAlbum specific album by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/albums/%7Balbum_id%7D
func (s *MeService) EditAlbum(ab string, r *AlbumRequest) (*Album, *Response, error) {
	u := fmt.Sprintf("me/albums/%s", ab)
	req, err := s.client.NewRequest("PATCH", u, r)
	if err != nil {
		return nil, nil, err
	}

	album := &Album{}
	resp, err := s.client.Do(req, album)
	if err != nil {
		return nil, resp, err
	}

	return album, resp, nil
}

// DeleteAlbum specific album by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/albums/%7Balbum_id%7D
func (s *MeService) DeleteAlbum(ab string) (*Response, error) {
	u := fmt.Sprintf("me/albums/%s", ab)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// AlbumListVideo lists the video for an album.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/albums/%7Balbum_id%7D/videos
func (s *MeService) AlbumListVideo(ab string, opt *ListVideoOptions) ([]*Video, *Response, error) {
	u := fmt.Sprintf("me/albums/%s/videos", ab)
	videos, resp, err := listVideo(s.client, u, opt)

	return videos, resp, err
}

// AlbumGetVideo specific video by album name and video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/albums/%7Balbum_id%7D/videos/%7Bvideo_id%7D
func (s *MeService) AlbumGetVideo(ab string, vid int) (*Video, *Response, error) {
	u := fmt.Sprintf("me/albums/%s/videos/%d", ab, vid)
	video, resp, err := getVideo(s.client, u)

	return video, resp, err
}

// AlbumDeleteVideo specific video by album name and video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/albums/%7Balbum_id%7D/videos/%7Bvideo_id%7D
func (s *MeService) AlbumDeleteVideo(ab string, vid int) (*Response, error) {
	u := fmt.Sprintf("me/albums/%s/videos/%d", ab, vid)

	resp, err := deleteVideo(s.client, u)

	return resp, err
}

// ListAppearance all videos a user is credited in.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/appearances
func (s *MeService) ListAppearance(opt *ListVideoOptions) ([]*Video, *Response, error) {
	u := fmt.Sprintf("me/appearances")

	videos, resp, err := listVideo(s.client, u, opt)

	return videos, resp, err
}

// ListCategory list the subscribed category for current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/categories
func (s *MeService) ListCategory(opt *ListCategoryOptions) ([]*Category, *Response, error) {
	categories, resp, err := listCategory(s.client, "me/categories", opt)

	return categories, resp, err
}

// SubscribeCategory subscribe category current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/categories/%7Bcategory%7D
func (s *MeService) SubscribeCategory(cat string) (*Response, error) {
	u := fmt.Sprintf("me/categories/%s", cat)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnsubscribeCategory unsubscribe category current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/categories/%7Bcategory%7D
func (s *MeService) UnsubscribeCategory(cat string) (*Response, error) {
	u := fmt.Sprintf("me/categories/%s", cat)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListChannel list the subscribed channel for current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/channels
func (s *MeService) ListChannel(opt *ListChannelOptions) ([]*Channel, *Response, error) {
	categories, resp, err := listChannel(s.client, "me/channels", opt)

	return categories, resp, err
}

// SubscribeChannel subscribe channel current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/channels/%7Bchannel_id%7D
func (s *MeService) SubscribeChannel(ch string) (*Response, error) {
	u := fmt.Sprintf("me/channels/%s", ch)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnsubscribeChannel unsubscribe channel current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/channels/%7Bchannel_id%7D
func (s *MeService) UnsubscribeChannel(ch string) (*Response, error) {
	u := fmt.Sprintf("me/channels/%s", ch)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Feed lists the feed for an current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/feed
func (s *MeService) Feed(opt *ListFeedOptions) ([]*Feed, *Response, error) {
	u, err := addOptions("me/feed", opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	feed := &dataListFeed{}

	resp, err := s.client.Do(req, feed)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(feed)

	return feed.Data, resp, err
}

// ListFollower lists the followers.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/channels/%7Bchannel_id%7D/users
func (s *MeService) ListFollower(opt *ListUserOptions) ([]*User, *Response, error) {
	users, resp, err := listUser(s.client, "/me/followers", opt)

	return users, resp, err
}

// ListFollowed lists the following.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/following
func (s *MeService) ListFollowed(opt *ListUserOptions) ([]*User, *Response, error) {
	users, resp, err := listUser(s.client, "/me/following", opt)

	return users, resp, err
}

// FollowUser follow a user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/following/%7Bfollow_user_id%7D
func (s *MeService) FollowUser(uid string) (*Response, error) {
	u := fmt.Sprintf("/me/following/%s", uid)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnfollowUser unfolloe a user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/following/%7Bfollow_user_id%7D
func (s *MeService) UnfollowUser(uid string) (*Response, error) {
	u := fmt.Sprintf("me/following/%s", uid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListGroup lists all joined groups.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/groups
func (s *MeService) ListGroup(opt *ListGroupOptions) ([]*Group, *Response, error) {
	groups, resp, err := listGroup(s.client, "/me/groups", opt)

	return groups, resp, err
}

// JoinGroup join current user to group.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/groups/%7Bgroup_id%7D
func (s *MeService) JoinGroup(uid string) (*Response, error) {
	u := fmt.Sprintf("/me/groups/%s", uid)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// LeaveGroup leaved current user from group.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/groups/%7Bgroup_id%7D
func (s *MeService) LeaveGroup(uid string) (*Response, error) {
	u := fmt.Sprintf("me/groups/%s", uid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListLikedVideo all liked videos.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/likes
func (s *MeService) ListLikedVideo(opt *ListVideoOptions) ([]*Video, *Response, error) {
	u := fmt.Sprintf("/me/likes")

	videos, resp, err := listVideo(s.client, u, opt)

	return videos, resp, err
}

// LikeVideo like one video.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/likes/%7Bvideo_id%7D
func (s *MeService) LikeVideo(vid int) (*Response, error) {
	u := fmt.Sprintf("/me/likes/%d", vid)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnlikeVideo unlike one video.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/likes/%7Bvideo_id%7D
func (s *MeService) UnlikeVideo(vid int) (*Response, error) {
	u := fmt.Sprintf("me/likes/%d", vid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// RemovePortrait removed specific a portrait.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/pictures/%7Bportraitset_id%7D
func (s *MeService) RemovePortrait(pid string) (*Response, error) {
	u := fmt.Sprintf("me/pictures/%s", pid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

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

// ListPortfolioOptions specifies the optional parameters to the
// ListPortfolio method.
type ListPortfolioOptions struct {
	Query     string `url:"query,omitempty"`
	Sort      string `url:"sort,omitempty"`
	Direction string `url:"direction,omitempty"`
	ListOptions
}

// ListPortfolio lists the portfolio for an current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/portfolios
func (s *MeService) ListPortfolio(opt *ListPortfolioOptions) ([]*Portfolio, *Response, error) {
	u, err := addOptions("me/portfolios", opt)
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
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/portfolios/%7Bportfolio_id%7D
func (s *MeService) GetProtfolio(p string) (*Portfolio, *Response, error) {
	u := fmt.Sprintf("me/portfolios/%s", p)
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
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/portfolios/%7Bportfolio_id%7D/videos
func (s *MeService) ProtfolioListVideo(p string, opt *ListVideoOptions) ([]*Video, *Response, error) {
	u := fmt.Sprintf("me/portfolios/%s/videos", p)
	videos, resp, err := listVideo(s.client, u, opt)

	return videos, resp, err
}

// ProtfolioGetVideo get specific video by portfolio name and video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/portfolios/%7Bportfolio_id%7D/videos/%7Bvideo_id%7D
func (s *MeService) ProtfolioGetVideo(p string, vid int) (*Video, *Response, error) {
	u := fmt.Sprintf("me/portfolios/%s/videos/%d", p, vid)
	video, resp, err := getVideo(s.client, u)

	return video, resp, err
}

// ProtfolioDeleteVideo delete specific video by portfolio name and video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/portfolios/%7Bportfolio_id%7D/videos/%7Bvideo_id%7D
func (s *MeService) ProtfolioDeleteVideo(p string, vid int) (*Response, error) {
	u := fmt.Sprintf("me/portfolios/%s/videos/%d", p, vid)
	resp, err := deleteVideo(s.client, u)

	return resp, err
}

type dataListPreset struct {
	Data []*Preset `json:"data,omitempty"`
	pagination
}

// Preset represents a preset.
type Preset struct {
	URI  string `json:"uri,omitempty"`
	Name string `json:"name,omitempty"`
}

// ListPresetOptions specifies the optional parameters to the
// ListPreset method.
type ListPresetOptions struct {
	Query string `url:"query,omitempty"`
	ListOptions
}

// ListPreset lists the preset for an current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/presets
func (s *MeService) ListPreset(opt *ListPresetOptions) ([]*Preset, *Response, error) {
	u, err := addOptions("me/presets", opt)
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
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/presets/%7Bpreset_id%7D
func (s *MeService) GetPreset(p int) (*Preset, *Response, error) {
	u := fmt.Sprintf("me/presets/%d", p)
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
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/presets/%7Bpreset_id%7D/videos
func (s *MeService) PresetListVideo(p string, opt *ListVideoOptions) ([]*Video, *Response, error) {
	u := fmt.Sprintf("me/presets/%s/videos", p)
	videos, resp, err := listVideo(s.client, u, opt)

	return videos, resp, err
}

// ListVideo lists the video for an current user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/videos
func (s *MeService) ListVideo(opt *ListVideoOptions) ([]*Video, *Response, error) {
	videos, resp, err := listVideo(s.client, "me/videos", opt)

	return videos, resp, err
}

// GetVideo get specific video by video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/videos/%7Bvideo_id%7D
func (s *MeService) GetVideo(vid int) (*Video, *Response, error) {
	u := fmt.Sprintf("me/videos/%d", vid)
	video, resp, err := getVideo(s.client, u)

	return video, resp, err
}

// WatchLaterListVideo lists the video.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/watchlater
func (s *MeService) WatchLaterListVideo(opt *ListVideoOptions) ([]*Video, *Response, error) {
	videos, resp, err := listVideo(s.client, "me/watchlater", opt)

	return videos, resp, err
}

// WatchLaterGetVideo get specific video by video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/watchlater/%7Bvideo_id%7D
func (s *MeService) WatchLaterGetVideo(vid int) (*Video, *Response, error) {
	u := fmt.Sprintf("me/watchlater/%d", vid)
	video, resp, err := getVideo(s.client, u)

	return video, resp, err
}

// WatchLaterAddVideo add video to watch later list.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/watchlater/%7Bvideo_id%7D
func (s *MeService) WatchLaterAddVideo(vid int) (*Response, error) {
	u := fmt.Sprintf("me/watchlater/%d", vid)
	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// WatchLaterDeleteVideo delete video from watch later list.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/watchlater/%7Bvideo_id%7D
func (s *MeService) WatchLaterDeleteVideo(vid int) (*Response, error) {
	u := fmt.Sprintf("me/watchlater/%d", vid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// WatchedListVideo lists the video.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/watched/videos
func (s *MeService) WatchedListVideo(opt *ListVideoOptions) ([]*Video, *Response, error) {
	videos, resp, err := listVideo(s.client, "me/watched/videos", opt)

	return videos, resp, err
}

// ClearWatchedList delete all video from watch history.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/watchlater/%7Bvideo_id%7D
func (s *MeService) ClearWatchedList() (*Response, error) {
	req, err := s.client.NewRequest("DELETE", "me/watched/videos", nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// WatchedDeleteVideo delete specific video from watch history.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/watched/videos/%7Bvideo_id%7D
func (s *MeService) WatchedDeleteVideo(vid int) (*Response, error) {
	u := fmt.Sprintf("me/watched/videos/%d", vid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
