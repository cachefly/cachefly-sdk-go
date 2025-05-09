package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"time"
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
		http: &http.Client{
			Timeout: 35 * time.Second, // Set a reasonable timeout
		},
		baseURL: cfg.BaseURL,
		token:   cfg.AuthToken,
	}
}

// Post performs a POST request with a JSON payload and decodes the JSON response.
func (c *Client) Post(ctx context.Context, endpoint string, body interface{}, out interface{}) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.fullURL(endpoint), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")
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
