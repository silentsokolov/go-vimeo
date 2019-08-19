package vimeo

import "fmt"

// CategoriesService handles communication with the categories related
// methods of the Vimeo API.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/categories
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

// List method gets all existing categories.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/categories#get_categories
func (s *CategoriesService) List(opt ...CallOption) ([]*Category, *Response, error) {
	categories, resp, err := listCategory(s.client, "categories", opt...)

	return categories, resp, err
}

// Get method gets a single category.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/categories#get_category
func (s *CategoriesService) Get(cat string, opt ...CallOption) (*Category, *Response, error) {
	u := fmt.Sprintf("categories/%s", cat)
	category, resp, err := getCategory(s.client, u, opt...)

	return category, resp, err
}

// ListChannel method gets all the channels that belong to a category.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/categories#get_category_channels
func (s *CategoriesService) ListChannel(cat string, opt ...CallOption) ([]*Channel, *Response, error) {
	u := fmt.Sprintf("categories/%s/channels", cat)
	channels, resp, err := listChannel(s.client, u, opt...)

	return channels, resp, err
}

// ListGroup method gets all the groups that belong to a category.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/categories#get_category_groups
func (s *CategoriesService) ListGroup(cat string, opt ...CallOption) ([]*Group, *Response, error) {
	u := fmt.Sprintf("categories/%s/groups", cat)
	groups, resp, err := listGroup(s.client, u, opt...)

	return groups, resp, err
}

// ListVideo method gets all the videos that belong to a category.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/categories#get_category_videos
func (s *CategoriesService) ListVideo(cat string, opt ...CallOption) ([]*Video, *Response, error) {
	u := fmt.Sprintf("categories/%s/videos", cat)
	videos, resp, err := listVideo(s.client, u, opt...)

	return videos, resp, err
}

// GetVideo method gets a single video from a category. Use it to determine whether the video belongs to the category.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/categories#check_category_for_video
func (s *CategoriesService) GetVideo(cat string, vid int, opt ...CallOption) (*Video, *Response, error) {
	u := fmt.Sprintf("categories/%s/videos/%d", cat, vid)
	video, resp, err := getVideo(s.client, u, opt...)

	return video, resp, err
}
