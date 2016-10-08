package vimeo

import "time"

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

// ListUserOptions specifies the optional parameters to the
// ListUser method.
type ListUserOptions struct {
	Query  string `url:"query,omitempty"`
	Filter string `url:"filter,omitempty"`
	ListOptions
}

// UserRequest represents a request to create/edit an user.
type UserRequest struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
	Bio      string `json:"bio,omitempty"`
}

func listUser(c *Client, url string, opt *ListUserOptions) ([]*User, *Response, error) {
	u, err := addOptions(url, opt)
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
