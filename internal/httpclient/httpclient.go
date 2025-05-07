package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
)

type Config struct {
	BaseURL   string
	AuthToken string
}

type Client struct {
	http    *http.Client
	baseURL string
	token   string
}

func New(cfg Config) *Client {
	return &Client{
		http:    http.DefaultClient,
		baseURL: cfg.BaseURL,
		token:   cfg.AuthToken,
	}
}

// Get performs a GET request and decodes the JSON response.
func (c *Client) Get(ctx context.Context, endpoint string, out interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.fullURL(endpoint), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error %d: %s", resp.StatusCode, string(body))
	}

	return json.NewDecoder(resp.Body).Decode(out)
}

func (c *Client) fullURL(endpoint string) string {
	return c.baseURL + path.Clean("/"+endpoint)
}
