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

// ListComment lists the comments.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/comments
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

// AddComment add comment.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/comments
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

// GetComment get specific comment by ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/comments/%7Bcomment_id%7D
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

// EditComment edit specific comment by ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/comments/%7Bcomment_id%7D
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

// DeleteComment delete specific comment by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/comments/%7Bcomment_id%7D
func (s *VideosService) DeleteComment(vid int, cid int) (*Response, error) {
	u := fmt.Sprintf("videos/%d/comments/%d", vid, cid)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// ListReplies lists the comment replies.
//
// https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/comments/%7Bcomment_id%7D/replies
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

// AddReplies add replies.
//
// https://developer.vimeo.com/api/playground/videos/%7Bvideo_id%7D/comments/%7Bcomment_id%7D/replies
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
