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
	URI         string     `json:"uri,omitempty"`
	Name        string     `json:"name,omitempty"`
	Link        string     `json:"link,omitempty"`
	Location    string     `json:"location,omitempty"`
	Bio         string     `json:"bio,omitempty"`
	CreatedTime time.Time  `json:"created_time,omitempty"`
	Account     string     `json:"account,omitempty"`
	Pictures    *Pictures  `json:"pictures,omitempty"`
	WebSites    []*WebSite `json:"websites,omitempty"`
	ResourceKey string     `json:"resource_key,omitempty"`
}

// ListUserOptions specifies the optional parameters to the
// ListUser method.
type ListUserOptions struct {
	Query  string `url:"query,omitempty"`
	Filter string `url:"filter,omitempty"`
	ListOptions
}
