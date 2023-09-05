package web

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// Client is a wrapper for http.Client
// more fields can be added in furure to make it more useful
type Client struct {
	http.Client
}

// New creates new Client instance
func New(client http.Client) *Client {
	return &Client{
		Client: client,
	}
}

// Get performes a get request
func (c *Client) Get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("making request for url: %s failed, err %w", url, err)
	}

	r, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request for url: %s failed, err %w", url, err)
	}

	resp, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("reading request for url: %s failed, err %w", url, err)
	}

	return resp, nil
}
