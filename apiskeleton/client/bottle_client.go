package client

import (
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
)

// ShowBottlePath computes a request path to the show action of bottle.
func ShowBottlePath(bottleID int) string {
	return fmt.Sprintf("/bottles/%v", bottleID)
}

// Get bottle by id
func (c *Client) ShowBottle(ctx context.Context, path string) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}
