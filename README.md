[![Build Status](https://travis-ci.org/silentsokolov/go-vimeo.svg?branch=master)](https://travis-ci.org/silentsokolov/go-vimeo)
[![GoDoc](https://godoc.org/github.com/silentsokolov/go-vimeo?status.svg)](https://godoc.org/github.com/silentsokolov/go-vimeo/) [![codecov](https://codecov.io/gh/silentsokolov/go-vimeo/branch/master/graph/badge.svg)](https://codecov.io/gh/silentsokolov/go-vimeo)

# go-vimeo

go-vimeo is a Go client library for accessing the [Vimeo API](https://developer.vimeo.com/api).

## Basic usage ##

```go
import "github.com/silentsokolov/go-vimeo"


func main() {
    client := vimeo.NewClient(nil)

    // Specific optional parameters
    opt := &vimeo.VideoListOptions{
        ListOptions: vimeo.ListOptions{Page: 1, PerPage: 2},
    }

    cats, _, err := client.Categories.List(opt)
}
```

### Authentication ###

The go-vimeo library does not directly handle authentication. Instead, when creating a new client, pass an http.Client that can handle authentication for you, for example the [oauth2](https://github.com/golang/oauth2).

```go
import "golang.org/x/oauth2"

func main() {
    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: "... your access token ..."},
    )
    tc := oauth2.NewClient(oauth2.NoContext, ts)

    client := vimeo.NewClient(tc)

    cats, _, err := client.Categories.List(opt)
}
```


### Pagination ###

```go
import "golang.org/x/oauth2"

func main() {
    client := ...

    // Specific optional parameters
    opt := &vimeo.VideoListOptions{
        ListOptions: vimeo.ListOptions{Page: 1, PerPage: 2},
    }

    // Any "List" request
    _, resp, _ := client.Categories.List(opt)

    fmt.Printf("Current page: %d", resp.Page)
	fmt.Printf("Next page: %d", resp.NextPage)
	fmt.Printf("Prev page: %d", resp.PrevPage)
	fmt.Printf("Total pages: %d", resp.TotalPages)
}
```


### Created/Updated request ###

```go
import "golang.org/x/oauth2"

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
