package vimeo

import (
	"fmt"
	"os"
	"time"
)

// UsersService handles communication with the users related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/users
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

// Search method information about this method appears below.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/users#search_users
func (s *UsersService) Search(opt ...CallOption) ([]*User, *Response, error) {
	users, resp, err := listUser(s.client, "users", opt...)

	return users, resp, err
}

// Get method returns the representation of the authenticated user.
// Passing the empty string will authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/users#get_user
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

// Edit method edits the representation of the authenticated user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/users#edit_user
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

// ListAppearance method returns all the videos in which the authenticated user has a credited appearance.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#get_appearances
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

// ListCategory method gets all the categories to which a particular user has subscribed.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/categories#get_category_subscriptions
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

// SubscribeCategory method subscribes the current user to a specified category.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/categories#subscribe_to_category
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

// UnsubscribeCategory method unsubscribes the current user from a specified category.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/categories#unsubscribe_from_category
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

// ListChannel method gets all the channels to which the specified user subscribes.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels#get_channel_subscriptions
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

// SubscribeChannel method causes a user to become the follower of the channel in question.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels#subscribe_to_channel
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

// UnsubscribeChannel method causes a user to stop following the channel in question.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/channels#unsubscribe_from_channel
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

// Feed method returns all the videos in the authenticated user's feed.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/users#get_feed
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

// ListFollower method returns all the followers of the authenticated user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/users#get_followers
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

// ListFollowed method causes the authenticated user to become the follower of multiple users.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/users#follow_users
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

// FollowUser method causes the authenticated user to become the follower of another user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/users#follow_user
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

// UnfollowUser method causes the authenticated user to stop following another user.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/users#unfollow_user
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

// ListGroup method returns all the groups to which a particular user belongs.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/groups#get_user_groups
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

// JoinGroup method adds a single user to the specified group.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/groups#join_group
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

// LeaveGroup method removes a single user from the specified group.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/groups#leave_group
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

// ListLikedVideo method gets all the videos that the specified user has liked.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/likes#get_likes
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

// LikeVideo method checks if the specified user has liked a particular video.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/likes#like_video
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

// UnlikeVideo method causes the specified user to unlike a video that they previously liked.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/likes#unlike_video
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

// RemovePortrait method removes a portrait image from the authenticated user's Vimeo account.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/users#delete_picture
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

// ListVideo method returns all the videos that the authenticated user has uploaded.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#get_videos
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

// GetVideo method determines whether a particular user is the owner of the specified video.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#check_if_user_owns_video
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

// UploadVideo method begins the video upload process for the authenticated user. For more information, see upload documentation.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#upload_video
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
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#upload_video
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

// WatchLaterListVideo method gets all the videos from the specified user's Watch Later queue.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/watch-later-queue#get_watch_later_queue
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

// WatchLaterGetVideo method checks the specified user's Watch Later queue for a particular video.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/watch-later-queue#check_watch_later_queue
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

// WatchLaterAddVideo method adds a single video to the specified user's Watch Later queue.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/watch-later-queue#add_video_to_watch_later
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

// WatchLaterDeleteVideo method removes a single video from the specified user's Watch Later queue.
// Passing the empty string will edit authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/watch-later-queue#delete_video_from_watch_later
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
