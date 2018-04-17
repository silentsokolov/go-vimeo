package vimeo

import "os"

type Uploader interface {
	UploadFromFile(c *Client, uploadURL string, f *os.File) error
}
