[![build](https://github.com/silentsokolov/go-vimeo/actions/workflows/build.yaml/badge.svg)](https://github.com/silentsokolov/go-vimeo/actions/workflows/build.yaml)
[![GoDoc](https://godoc.org/github.com/silentsokolov/go-vimeo/v2?status.svg)](https://godoc.org/github.com/silentsokolov/go-vimeo/v2/vimeo) [![codecov](https://codecov.io/gh/silentsokolov/go-vimeo/branch/master/graph/badge.svg)](https://codecov.io/gh/silentsokolov/go-vimeo)
[![Go Report Card](https://goreportcard.com/badge/github.com/silentsokolov/go-vimeo/v2)](https://goreportcard.com/report/github.com/silentsokolov/go-vimeo/v2)

# go-vimeo

go-vimeo is a Go client library for accessing the [Vimeo API](https://developer.vimeo.com/api).

## Basic usage ##

```go
import "github.com/silentsokolov/go-vimeo/v2/vimeo"


func main() {
	client := vimeo.NewClient(tokenContext, nil)

	// Specific optional parameters
	cats, _, err := client.Categories.List(OptPage(1), OptPerPage(2), OptFields([]string{"name"}))
}
```

### Authentication ###

The go-vimeo library does not directly handle authentication. Instead, when creating a new client, pass an http.Client that can handle authentication for you, for example the [oauth2](https://github.com/golang/oauth2).

```go
import (
	"golang.org/x/oauth2"
	"github.com/silentsokolov/go-vimeo/v2/vimeo"
)

func main() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := vimeo.NewClient(tc, nil)

	cats, _, err := client.Categories.List()
}
```


### Pagination ###

```go
func main() {
	client := ...

	// Any "List" request
	_, resp, _ := client.Categories.List(OptPage(2), OptPerPage(2))

	fmt.Printf("Current page: %d\n", resp.Page)
	fmt.Printf("Next page: %s\n", resp.NextPage)
	fmt.Printf("Prev page: %s\n", resp.PrevPage)
	fmt.Printf("Total objects: %d\n", resp.Total)
}
```


### Created/Updated request ###

```go
func main() {
	client := ...

	// Specific request instance
	req := &vimeo.ChannelRequest{
		Name:        "My Channel",
		Description: "Awesome",
		Privacy:     "anybody",
	}

	ch, _, _ := client.Channels.Create(req)

	fmt.Println(ch)
}
```


### Where "Me" service? ###

The "Me" service repeats the "Users" service, passing the empty string will authenticated user.

```go
func main() {
	client := ...

	// Call /me API method.
	// Return current authenticated user.
	me, _, _ := client.Users.Get("")

	fmt.Println(me)


	// Call /me/videos API method.
	// Return videos for current authenticated user.
	videos, _, _ := client.Users.ListVideo("")

	fmt.Println(videos)
}
```

### Upload video ###

Since the release of Vimeo API version 3.4 used to unload the video the [tus protocol](https://tus.io/). Necessary to implement the process manually. You can use the implementation of [tus protocol on golang](https://github.com/eventials/go-tus).

```go
import (
	"os"

	"golang.org/x/oauth2"
	"github.com/silentsokolov/go-vimeo/v2/vimeo"

	tus "github.com/eventials/go-tus"
)

type Uploader struct{}

func (u Uploader) UploadFromFile(c *vimeo.Client, uploadURL string, f *os.File) error {
	tusClient, err := tus.NewClient(uploadURL, nil)
	if err != nil {
		return err
	}

	upload, err := tus.NewUploadFromFile(f)
	if err != nil {
		return err
	}

	uploader := tus.NewUploader(tusClient, uploadURL, upload, 0)

	return uploader.Upload()
}

func main() {
	config := vimeo.Config{
		Uploader: &Uploader{},
	}

	tc := ...
	client := vimeo.NewClient(tc, &config)

	filePath := "/Users/user/Videos/Awesome.mp4"

	f, _ := os.Open(filePath)

	video, resp, _ := client.Users.UploadVideo("", f)

	fmt.Println(video, resp)
}
```
