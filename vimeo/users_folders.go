package vimeo

import (
	"fmt"
	"strings"
	"time"
)

type dataListFolder struct {
	Data []*Folder `json:"data,omitempty"`
	pagination
}

// Folder represents a Vimeo folder (project).
type Folder struct {
	URI          string    `json:"uri,omitempty"`
	Name         string    `json:"name,omitempty"`
	ResourceKey  string    `json:"resource_key,omitempty"`
	CreatedTime  time.Time `json:"created_time,omitempty"`
	ModifiedTime time.Time `json:"modified_time,omitempty"`
	ParentFolder *Folder   `json:"parent_folder,omitempty"`
}

func listFolder(c *Client, url string, opt ...CallOption) ([]*Folder, *Response, error) {
	u, err := addOptions(url, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	folders := &dataListFolder{}

	resp, err := c.Do(req, folders)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(folders)

	return folders.Data, resp, err
}

// FolderItem is a single entry returned by the folder /items endpoint.
// Its Type field is either "video" or "folder".
type FolderItem struct {
	Type   string  `json:"type,omitempty"`
	Video  *Video  `json:"video,omitempty"`
	Folder *Folder `json:"folder,omitempty"`
}

type dataListFolderItem struct {
	Data []*FolderItem `json:"data,omitempty"`
	pagination
}

func listFolderItem(c *Client, url string, opt ...CallOption) ([]*FolderItem, *Response, error) {
	u, err := addOptions(url, opt...)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	items := &dataListFolderItem{}

	resp, err := c.Do(req, items)
	if err != nil {
		return nil, resp, err
	}

	resp.setPaging(items)

	return items.Data, resp, err
}

// ListFolders lists all root-level folders for a user.
// Passing the empty string will use the authenticated user.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/folders#get_folders
func (s *UsersService) ListFolders(uid string, opt ...CallOption) ([]*Folder, *Response, error) {
	var u string
	if uid == "" {
		u = "me/folders"
	} else {
		u = fmt.Sprintf("users/%s/folders", uid)
	}
	return listFolder(s.client, u, opt...)
}

// ListFolderItems lists all items (videos and sub-folders) within a given folder
// via the /items endpoint. folderURI is the full URI returned by the API
// (e.g. "/users/12345/projects/67890").
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/folders#get_folder_items
func (s *UsersService) ListFolderItems(folderURI string, opt ...CallOption) ([]*FolderItem, *Response, error) {
	u := strings.TrimPrefix(folderURI, "/") + "/items"
	return listFolderItem(s.client, u, opt...)
}

// ListFolderVideos lists all videos within a given folder via the /videos endpoint.
// folderURI is the full URI returned by the API (e.g. "/users/12345/projects/67890").
// The response includes full download metadata.
//
// Vimeo API docs: https://developer.vimeo.com/api/reference/folders#get_folder_items
func (s *UsersService) ListFolderVideos(folderURI string, opt ...CallOption) ([]*Video, *Response, error) {
	u := strings.TrimPrefix(folderURI, "/") + "/videos"
	return listVideo(s.client, u, opt...)
}
