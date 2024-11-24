package internal

import (
	"errors"
	"os"
)

var (
    CLIENT_URL string = os.Getenv("CLIENT_URL")
)

// Client provides an interface to access internals such as
// api keys, etc that are not stored in memory.
type Client struct {
    ApiKey string
}

func (c *Client) RequestKey() error { 
    c.ApiKey = ""
    return errors.New("Function not implemented")
}
