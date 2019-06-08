/*
Package ssllabs provides API access on Qualys SSL Labs.
The use of this client is subject to https://www.ssllabs.com/about/terms.html
*/
package ssllabs

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"

	"github.com/pkg/errors"
	"golang.org/x/net/publicsuffix"
)

const (
	// Base string for API endpoints
	BaseURL = "https://api.ssllabs.com/api/v3"
	// Base string for beta API endpoints
	BaseURLDev = "https://api.dev.ssllabs.com/api/v3"
)

// httpClient defines minimal interface for client
type httpClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Client is the ssllabs client for API access
type Client struct {
	baseURL    string
	httpClient httpClient
}

// Param is parameter for API call
type Param struct {
	Key   string
	Value string
}

// New creates a ssllabs client with given options.
func New(options ...Option) (*Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar: jar,
	}

	c := &Client{
		baseURL:    BaseURL,
		httpClient: client,
	}

	for _, opt := range options {
		opt(c)
	}

	if err := c.normalizeBaseURL(); err != nil {
		return c, errors.Wrap(err, "Invalid URL was given")
	}

	return c, nil
}

func (c *Client) normalizeBaseURL() error {
	u, err := url.Parse(c.baseURL)
	if err != nil {
		return err
	}

	if u.Scheme == "" {
		return fmt.Errorf("failed to identify the scheme: %s", u)
	}

	r := regexp.MustCompile(`^https?$`)
	if !r.MatchString(u.Scheme) {
		return fmt.Errorf("unsupported scheme: %s", u.Scheme)
	}

	if u.Host == "" {
		return fmt.Errorf("failed to identify the hostname: %s", u)
	}

	if c.baseURL[len(c.baseURL)-1:] == "/" {
		c.baseURL = c.baseURL[:len(c.baseURL)-1]
	}
	return nil
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
		values.Set(param.Key, param.Value)
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
