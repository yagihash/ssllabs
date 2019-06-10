package ssllabs

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type StatusCodes struct {
	StatusDetails map[string]string `json:"statusDetails"`
}

// GetStatusCodes is interface to /getStatusCodes
func (c *Client) GetStatusCodes() (*StatusCodes, *http.Response, error) {
	req, err := c.newRequest("GET", "/getStatusCodes")
	if err != nil {
		return nil, nil, err
	}

	var buf bytes.Buffer

	resp, err := c.request(req, &buf)
	if err != nil {
		return nil, nil, err
	}

	data := &StatusCodes{}
	if err := json.Unmarshal(buf.Bytes(), data); err != nil {
		return nil, nil, err
	}
	return data, resp, nil
}
