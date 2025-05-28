// Package beget implements a Beget API client.
package beget

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// Client is a Beget API client.
type Client struct {
	h        *http.Client
	username string
	password string
	baseUrl  string
}

const DefaultBaseURL = "https://api.beget.com/api"

// New returns a new Beget API client.
// You can set options with the provided options.
func New(username string, password string, options ...Option) *Client {
	c := &Client{username: username, password: password, baseUrl: DefaultBaseURL}
	for _, o := range options {
		o(c)
	}
	if c.h == nil {
		c.h = &http.Client{}
	}
	return c
}

func (c *Client) do(ctx context.Context, endpoint string, input any, output any) error {
	u, err := c.buildUrl(endpoint, input)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	resp, err := c.h.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(output)
}

func (c *Client) buildUrl(endpoint string, input any) (string, error) {
	u, err := url.Parse(c.baseUrl)
	if err != nil {
		return "", err
	}
	u = u.JoinPath(endpoint)
	data, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	values := url.Values{}
	values.Set("login", c.username)
	values.Set("passwd", c.password)
	values.Set("input_format", "json")
	values.Set("output_format", "json")
	values.Set("input_data", string(data))

	u.RawQuery = values.Encode()
	return u.String(), nil
}
