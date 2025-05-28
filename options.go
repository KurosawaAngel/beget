package beget

import "net/http"

type Option func(*Client)

// WithHTTPClient sets the HTTP client to use for requests.
func WithHTTPClient(h *http.Client) Option {
	return func(c *Client) {
		c.h = h
	}
}

// WithBaseUrl sets the base URL to use for requests.
func WithBaseUrl(url string) Option {
	return func(c *Client) {
		c.baseUrl = url
	}
}
