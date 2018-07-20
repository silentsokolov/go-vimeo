package vimeo

import "fmt"

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
	Parent                *SubCategory   `json:"parent,omitempty"`
	SubCategories         []*SubCategory `json:"subcategories,omitempty"`
	ResourceKey           string         `json:"resource_key,omitempty"`
}

// SubCategory internal object provides access to subcategory in category.
type SubCategory struct {
	URI  string `json:"URI,omitempty"`
	Name string `json:"name,omitempty"`
	Link string `json:"link,omitempty"`
}

func listCategory(c *Client, url string, opt ...CallOption) ([]*Category, *Response, error) {
	u, err := addOptions(url, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	categories := &dataListCategory{}

	resp, err := c.Do(req, categories)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(categories)

	return categories.Data, resp, err
}

func getCategory(c *Client, url string, opt ...CallOption) (*Category, *Response, error) {
	u, err := addOptions(url, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	category := &Category{}

	resp, err := c.Do(req, category)
	if err != nil {
		return nil, resp, err
	}

	return category, resp, err
}

// List the category.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/categories
func (s *CategoriesService) List(opt ...CallOption) ([]*Category, *Response, error) {
	categories, resp, err := listCategory(s.client, "categories", opt...)

	return categories, resp, err
}

// Get specific category by name.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/categories/%7Bcategory%7D
func (s *CategoriesService) Get(cat string, opt ...CallOption) (*Category, *Response, error) {
	u := fmt.Sprintf("categories/%s", cat)
	category, resp, err := getCategory(s.client, u, opt...)

	return category, resp, err
}

// ListChannel lists the channel for an category.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/categories/%7Bcategory%7D/channels
func (s *CategoriesService) ListChannel(cat string, opt ...CallOption) ([]*Channel, *Response, error) {
	u := fmt.Sprintf("categories/%s/channels", cat)
	channels, resp, err := listChannel(s.client, u, opt...)

	return channels, resp, err
}

// ListGroup lists the group for an category.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/categories/%7Bcategory%7D/groups
func (s *CategoriesService) ListGroup(cat string, opt ...CallOption) ([]*Group, *Response, error) {
	u := fmt.Sprintf("categories/%s/groups", cat)
	groups, resp, err := listGroup(s.client, u, opt...)

	return groups, resp, err
}

// ListVideo lists the video for an category.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/categories/%7Bcategory%7D/videos
func (s *CategoriesService) ListVideo(cat string, opt ...CallOption) ([]*Video, *Response, error) {
	u := fmt.Sprintf("categories/%s/videos", cat)
	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}

// GetVideo specific video by category name and video ID.
//
// Vimeo API docs: https://developer.vimeo.com/api/playground/categories/%7Bcategory%7D/videos/%7Bvideo_id%7D
func (s *CategoriesService) GetVideo(cat string, vid int, opt ...CallOption) (*Video, *Response, error) {
	u := fmt.Sprintf("categories/%s/videos/%d", cat, vid)
	video, resp, err := getVideo(s.client, u, opt...)

	return video, resp, err
}
