package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	HTTP      *http.Client
	AuthName  string // "PRIVATE-TOKEN"
	AuthValue string // the token
}

func NewClient(authName, authValue string) *Client {
	return &Client{
		HTTP:      &http.Client{Timeout: 30 * time.Second},
		AuthName:  authName,
		AuthValue: authValue,
	}
}

// GetJSON fetches url, decodes the JSON body into v, and returns
// the response headers (needed for pagination).
func (c *Client) GetJSON(ctx context.Context, url string, v any) (http.Header, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(c.AuthName, c.AuthValue)

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: HTTP %d", url, resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
		return nil, fmt.Errorf("decode %s: %w", url, err)
	}
	return resp.Header, nil
}
