[![Build Status](https://travis-ci.org/silentsokolov/go-vimeo.svg?branch=master)](https://travis-ci.org/silentsokolov/go-vimeo)
[![GoDoc](https://godoc.org/github.com/silentsokolov/go-vimeo?status.svg)](https://godoc.org/github.com/silentsokolov/go-vimeo/) [![codecov](https://codecov.io/gh/silentsokolov/go-vimeo/branch/master/graph/badge.svg)](https://codecov.io/gh/silentsokolov/go-vimeo)
[![Go Report Card](https://goreportcard.com/badge/github.com/silentsokolov/go-vimeo)](https://goreportcard.com/report/github.com/silentsokolov/go-vimeo)

# go-vimeo

go-vimeo is a Go client library for accessing the [Vimeo API](https://developer.vimeo.com/api).

## Basic usage ##

```go
import "github.com/silentsokolov/go-vimeo/vimeo"


func main() {
	client := vimeo.NewClient(nil)

	// Specific optional parameters
	cats, _, err := client.Categories.List(Page(1), PerPage(2))
}
```

### Authentication ###

The go-vimeo library does not directly handle authentication. Instead, when creating a new client, pass an http.Client that can handle authentication for you, for example the [oauth2](https://github.com/golang/oauth2).

```go
import (
	"golang.org/x/oauth2"
	"github.com/silentsokolov/go-vimeo/vimeo"
)

func main() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := vimeo.NewClient(tc)

	cats, _, err := client.Categories.List()
}
```


### Pagination ###

```go
func main() {
	client := ...

	// Any "List" request
	_, resp, _ := client.Categories.List(Page(2), PerPage(2))

	fmt.Printf("Current page: %d\n", resp.Page)
	fmt.Printf("Next page: %s\n", resp.NextPage)
	fmt.Printf("Prev page: %s\n", resp.PrevPage)
	fmt.Printf("Total pages: %d\n", resp.TotalPages)
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

```go
import (
	"os"

	"golang.org/x/oauth2"
	"github.com/silentsokolov/go-vimeo/vimeo"
)

func main() {
	client := ...
	filePath := "/Users/user/Videos/Awesome.mp4"

	f, _ := os.Open(filePath)

	video, resp, _ := client.Users.UploadVideo("", f)

	fmt.Println(video, resp)
}
```
