package ssllabs

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
)

const (
	BaseURL    = "https://api.ssllabs.com/api/v3"
	BaseURLDev = "https://api.dev.ssllabs.com/api/v3"
)

type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	baseURL    string
	httpClient httpClient
}

type Param struct {
	key   string
	value string
}

func New(options ...Option) *Client {
	c := &Client{
		baseURL:    BaseURL,
		httpClient: &http.Client{},
	}

	for _, opt := range options {
		opt(c)
	}

	return c
}

func (c *Client) request(req *http.Request, buf *bytes.Buffer) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(buf, resp.Body); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) newRequest(method string, path string, params ...*Param) (*http.Request, error) {
	values := url.Values{}

	for _, param := range params {
		values.Set(param.key, param.value)
	}

	req, err := http.NewRequest(
		method,
		c.baseURL+path+"?"+values.Encode(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	return req, nil
}
