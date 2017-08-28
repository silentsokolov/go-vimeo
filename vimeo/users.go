package vimeo

import (
	"fmt"
	"os"
	"time"
)

// UsersService handles communication with the users related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/endpoints/users
type UsersService service

type dataListUser struct {
	Data []*User `json:"data"`
	pagination
}

// WebSite represents a web site.
type WebSite struct {
	Name        string `json:"name,omitempty"`
	Link        string `json:"link,omitempty"`
	Description string `json:"description,omitempty"`
}

// User represents a user.
type User struct {
	URI           string     `json:"uri,omitempty"`
	Name          string     `json:"name,omitempty"`
	Link          string     `json:"link,omitempty"`
	Location      string     `json:"location,omitempty"`
	Bio           string     `json:"bio,omitempty"`
	CreatedTime   time.Time  `json:"created_time,omitempty"`
	Account       string     `json:"account,omitempty"`
	Pictures      *Pictures  `json:"pictures,omitempty"`
	WebSites      []*WebSite `json:"websites,omitempty"`
	ContentFilter []string   `json:"content_filter,omitempty"`
	ResourceKey   string     `json:"resource_key,omitempty"`
}

// UserRequest represents a request to create/edit an user.
type UserRequest struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
	Bio      string `json:"bio,omitempty"`
}

func listUser(c *Client, url string, opt ...CallOption) ([]*User, *Response, error) {
	u, err := addOptions(url, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	users := &dataListUser{}

	resp, err := c.Do(req, users)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(users)

	return users.Data, resp, err
}

// Search users.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/channels/%7Bchannel_id%7D/users
func (s *UsersService) Search(opt ...CallOption) ([]*User, *Response, error) {
	users, resp, err := listUser(s.client, "users", opt...)

	return users, resp, err
}

// Get show one user.
// Passing the empty string will authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D
func (s *UsersService) Get(uid string, opt ...CallOption) (*User, *Response, error) {
	var u string
	if uid == "" {
		u = "me"
	} else {
		u = fmt.Sprintf("users/%s", uid)
	}

	u, err := addOptions(u, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
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

// Edit one user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D
func (s *UsersService) Edit(uid string, r *UserRequest) (*User, *Response, error) {
	var u string
	if uid == "" {
		u = "me"
	} else {
		u = fmt.Sprintf("users/%s", uid)
	}

	req, err := s.client.NewRequest("PATCH", u, r)
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

// ListAppearance all videos a user is credited in.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/appearances
func (s *UsersService) ListAppearance(uid string, opt ...CallOption) ([]*Video, *Response, error) {
	var u string
	if uid == "" {
		u = "me/appearances"
	} else {
		u = fmt.Sprintf("users/%s/appearances", uid)
	}

	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}

// ListCategory list the subscribed category for user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/categories
func (s *UsersService) ListCategory(uid string, opt ...CallOption) ([]*Category, *Response, error) {
	var u string
	if uid == "" {
		u = "me/categories"
	} else {
		u = fmt.Sprintf("users/%s/categories", uid)
	}

	categories, resp, err := listCategory(s.client, u, opt...)

	return categories, resp, err
}

// SubscribeCategory subscribe category user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/categories/%7Bcategory%7D
func (s *UsersService) SubscribeCategory(uid string, cat string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/categories/%s", cat)
	} else {
		u = fmt.Sprintf("users/%s/categories/%s", uid, cat)
	}

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnsubscribeCategory unsubscribe category current user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/categories/%7Bcategory%7D
func (s *UsersService) UnsubscribeCategory(uid string, cat string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/categories/%s", cat)
	} else {
		u = fmt.Sprintf("users/%s/categories/%s", uid, cat)
	}

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListChannel list the subscribed channel for user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/channels
func (s *UsersService) ListChannel(uid string, opt ...CallOption) ([]*Channel, *Response, error) {
	var u string
	if uid == "" {
		u = "me/channels"
	} else {
		u = fmt.Sprintf("users/%s/channels", uid)
	}

	categories, resp, err := listChannel(s.client, u, opt...)

	return categories, resp, err
}

// SubscribeChannel subscribe channel user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/channels/%7Bchannel_id%7D
func (s *UsersService) SubscribeChannel(uid string, ch string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/channels/%s", ch)
	} else {
		u = fmt.Sprintf("users/%s/channels/%s", uid, ch)
	}

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnsubscribeChannel unsubscribe channel user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/channels/%7Bchannel_id%7D
func (s *UsersService) UnsubscribeChannel(uid string, ch string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/channels/%s", ch)
	} else {
		u = fmt.Sprintf("users/%s/channels/%s", uid, ch)
	}

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

type dataListFeed struct {
	Data []*Feed `json:"data"`
	pagination
}

// Feed represents a feed.
type Feed struct {
	URI  string `json:"uri,omitempty"`
	Clip *Video `json:"clip,omitempty"`
}

// Feed lists the feed for an user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/feed
func (s *UsersService) Feed(uid string, opt ...CallOption) ([]*Feed, *Response, error) {
	var u string
	if uid == "" {
		u = "me/feed"
	} else {
		u = fmt.Sprintf("users/%s/feed", uid)
	}

	u, err := addOptions(u, opt...)
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
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/followers
func (s *UsersService) ListFollower(uid string, opt ...CallOption) ([]*User, *Response, error) {
	var u string
	if uid == "" {
		u = "me/followers"
	} else {
		u = fmt.Sprintf("users/%s/followers", uid)
	}

	users, resp, err := listUser(s.client, u, opt...)

	return users, resp, err
}

// ListFollowed lists the following.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/following
func (s *UsersService) ListFollowed(uid string, opt ...CallOption) ([]*User, *Response, error) {
	var u string
	if uid == "" {
		u = "me/following"
	} else {
		u = fmt.Sprintf("users/%s/following", uid)
	}

	users, resp, err := listUser(s.client, u, opt...)

	return users, resp, err
}

// FollowUser follow a user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/following/%7Bfollow_user_id%7D
func (s *UsersService) FollowUser(uid string, fid string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/following/%s", fid)
	} else {
		u = fmt.Sprintf("users/%s/following/%s", uid, fid)
	}

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnfollowUser unfollow a user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/following/%7Bfollow_user_id%7D
func (s *UsersService) UnfollowUser(uid string, fid string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/following/%s", fid)
	} else {
		u = fmt.Sprintf("users/%s/following/%s", uid, fid)
	}

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListGroup lists all joined groups.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/groups
func (s *UsersService) ListGroup(uid string, opt ...CallOption) ([]*Group, *Response, error) {
	var u string
	if uid == "" {
		u = "me/groups"
	} else {
		u = fmt.Sprintf("users/%s/groups", uid)
	}

	groups, resp, err := listGroup(s.client, u, opt...)

	return groups, resp, err
}

// JoinGroup join user to group.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/groups/%7Bgroup_id%7D
func (s *UsersService) JoinGroup(uid string, gid string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/groups/%s", gid)
	} else {
		u = fmt.Sprintf("users/%s/groups/%s", uid, gid)
	}

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// LeaveGroup leaved user from group.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/groups/%7Bgroup_id%7D
func (s *UsersService) LeaveGroup(uid string, gid string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/groups/%s", gid)
	} else {
		u = fmt.Sprintf("users/%s/groups/%s", uid, gid)
	}

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListLikedVideo all liked videos.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/likes
func (s *UsersService) ListLikedVideo(uid string, opt ...CallOption) ([]*Video, *Response, error) {
	var u string
	if uid == "" {
		u = "me/likes"
	} else {
		u = fmt.Sprintf("users/%s/likes", uid)
	}

	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}

// LikeVideo like one video.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/likes/%7Bvideo_id%7D
func (s *UsersService) LikeVideo(uid string, vid int) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/likes/%d", vid)
	} else {
		u = fmt.Sprintf("users/%s/likes/%d", uid, vid)
	}

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// UnlikeVideo unlike one video.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/likes/%7Bvideo_id%7D
func (s *UsersService) UnlikeVideo(uid string, vid int) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/likes/%d", vid)
	} else {
		u = fmt.Sprintf("users/%s/likes/%d", uid, vid)
	}

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// RemovePortrait removed specific a portrait.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/pictures/%7Bportraitset_id%7D
func (s *UsersService) RemovePortrait(uid string, pid string) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/pictures/%s", pid)
	} else {
		u = fmt.Sprintf("users/%s/pictures/%s", uid, pid)
	}

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListVideo lists the video for user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/videos
func (s *UsersService) ListVideo(uid string, opt ...CallOption) ([]*Video, *Response, error) {
	var u string
	if uid == "" {
		u = "me/videos"
	} else {
		u = fmt.Sprintf("users/%s/videos", uid)
	}

	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}

// GetVideo get specific video by video ID.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/videos
func (s *UsersService) GetVideo(uid string, vid int, opt ...CallOption) (*Video, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/videos/%d", vid)
	} else {
		u = fmt.Sprintf("users/%s/videos/%d", uid, vid)
	}

	video, resp, err := getVideo(s.client, u, opt...)

	return video, resp, err
}

// UploadVideo upload video file.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/videos
func (s *UsersService) UploadVideo(uid string, file *os.File) (*Video, *Response, error) {
	var u string
	if uid == "" {
		u = "me/videos"
	} else {
		u = fmt.Sprintf("users/%s/videos", uid)
	}

	video, resp, err := uploadVideo(s.client, "POST", u, file)

	return video, resp, err
}

// UploadVideo upload video by url.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/videos
func (s *UsersService) UploadVideoByURL(uid string, videoURL string) (*Video, *Response, error) {
	var u string
	if uid == "" {
		u = "me/videos"
	} else {
		u = fmt.Sprintf("users/%s/videos", uid)
	}

	video, resp, err := uploadVideoByURL(s.client, u, videoURL)

	return video, resp, err
}

// WatchLaterListVideo lists the video.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/watchlater
func (s *UsersService) WatchLaterListVideo(uid string, opt ...CallOption) ([]*Video, *Response, error) {
	var u string
	if uid == "" {
		u = "me/watchlater"
	} else {
		u = fmt.Sprintf("users/%s/watchlater", uid)
	}

	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}

// WatchLaterGetVideo get specific video by video ID.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/watchlater/%7Bvideo_id%7D
func (s *UsersService) WatchLaterGetVideo(uid string, vid int) (*Video, *Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/watchlater/%d", vid)
	} else {
		u = fmt.Sprintf("users/%s/watchlater/%d", uid, vid)
	}

	video, resp, err := getVideo(s.client, u)

	return video, resp, err
}

// WatchLaterAddVideo add video to watch later list.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/watchlater/%7Bvideo_id%7D
func (s *UsersService) WatchLaterAddVideo(uid string, vid int) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/watchlater/%d", vid)
	} else {
		u = fmt.Sprintf("users/%s/watchlater/%d", uid, vid)
	}

	req, err := s.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// WatchLaterDeleteVideo delete video from watch later list.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/users/%7Buser_id%7D/watchlater/%7Bvideo_id%7D
func (s *UsersService) WatchLaterDeleteVideo(uid string, vid int) (*Response, error) {
	var u string
	if uid == "" {
		u = fmt.Sprintf("me/watchlater/%d", vid)
	} else {
		u = fmt.Sprintf("users/%s/watchlater/%d", uid, vid)
	}

	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// WatchedListVideo lists the video.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/watched/videos
func (s *UsersService) WatchedListVideo(uid string, opt ...CallOption) ([]*Video, *Response, error) {
	videos, resp, err := listVideo(s.client, "me/watched/videos", opt...)

	return videos, resp, err
}

// ClearWatchedList delete all video from watch history.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/watchlater/%7Bvideo_id%7D
func (s *UsersService) ClearWatchedList(uid string) (*Response, error) {
	req, err := s.client.NewRequest("DELETE", "me/watched/videos", nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// WatchedDeleteVideo delete specific video from watch history.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/me/watched/videos/%7Bvideo_id%7D
func (s *UsersService) WatchedDeleteVideo(uid string, vid int) (*Response, error) {
	u := fmt.Sprintf("me/watched/videos/%d", vid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
