package vimeo

import "fmt"

type dataListComment struct {
	Data []*Comment `json:"data,omitempty"`
	pagination
}

// Comment represents a comment.
type Comment struct {
	URI         string `json:"uri,omitempty"`
	Type        string `json:"type,omitempty"`
	Text        string `json:"text,omitempty"`
	CreatedOn   string `json:"created_on,omitempty"`
	User        *User  `json:"user,omitempty"`
	ResourceKey string `json:"resource_key,omitempty"`
}

// CommentRequest represents a request to create/edit an comment.
type CommentRequest struct {
	Text string `json:"text,omitempty"`
}

// ListComment method returns all the comments on the specified video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#get_comments
func (s *VideosService) ListComment(vid int, opt ...CallOption) ([]*Comment, *Response, error) {
	u, err := addOptions(fmt.Sprintf("videos/%d/comments", vid), opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	comments := &dataListComment{}

	resp, err := s.client.Do(req, comments)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(comments)

	return comments.Data, resp, err
}

// AddComment method adds a comment to the specified video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#create_comment
func (s *VideosService) AddComment(vid int, r *CommentRequest) (*Comment, *Response, error) {
	u := fmt.Sprintf("videos/%d/comments", vid)
	req, err := s.client.NewRequest("POST", u, r)
	if err != nil {
		return nil, nil, err
	}

	comment := &Comment{}
	resp, err := s.client.Do(req, comment)
	if err != nil {
		return nil, resp, err
	}

	return comment, resp, nil
}

// GetComment method returns the specified comment on a video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#get_comment
func (s *VideosService) GetComment(vid int, cid int, opt ...CallOption) (*Comment, *Response, error) {
	u, err := addOptions(fmt.Sprintf("videos/%d/comments/%d", vid, cid), opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	comment := &Comment{}

	resp, err := s.client.Do(req, comment)
	if err != nil {
		return nil, resp, err
	}

	return comment, resp, err
}

// EditComment method edits the specified comment on a video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#edit_comment
func (s *VideosService) EditComment(vid int, cid int, r *CommentRequest) (*Comment, *Response, error) {
	u := fmt.Sprintf("videos/%d/comments/%d", vid, cid)
	req, err := s.client.NewRequest("PATCH", u, r)
	if err != nil {
		return nil, nil, err
	}

	comment := &Comment{}
	resp, err := s.client.Do(req, comment)
	if err != nil {
		return nil, resp, err
	}

	return comment, resp, nil
}

// DeleteComment method deletes the specified comment from a video.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#delete_comment
func (s *VideosService) DeleteComment(vid int, cid int) (*Response, error) {
	u := fmt.Sprintf("videos/%d/comments/%d", vid, cid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListReplies method returns all the replies to the specified video comment.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#get_comment_replies
func (s *VideosService) ListReplies(vid int, cid int, opt ...CallOption) ([]*Comment, *Response, error) {
	u, err := addOptions(fmt.Sprintf("videos/%d/comments/%d/replies", vid, cid), opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	replies := &dataListComment{}

	resp, err := s.client.Do(req, replies)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(replies)

	return replies.Data, resp, err
}

// AddReplies method adds a reply to the specified video comment.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/videos#create_comment_reply
func (s *VideosService) AddReplies(vid int, cid int, r *CommentRequest) (*Comment, *Response, error) {
	u := fmt.Sprintf("videos/%d/comments/%d/replies", vid, cid)
	req, err := s.client.NewRequest("POST", u, r)
	if err != nil {
		return nil, nil, err
	}

	replies := &Comment{}
	resp, err := s.client.Do(req, replies)
	if err != nil {
		return nil, resp, err
	}

	return replies, resp, nil
}
