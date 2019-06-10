package ssllabs

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Info struct {
	EngineVersion        string   `json:"engineVersion"`
	CriteriaVersion      string   `json:"criteriaVersion"`
	MaxAssessments       int      `json:"maxAssessments"`
	CurrentAssessments   int      `json:"currentAssessments"`
	NewAssessmentCoolOff int      `json:"newAssessmentCoolOff"`
	Messages             []string `json:"messages"`
}

// Info is interface to /info
func (c *Client) Info() (*Info, *http.Response, error) {
	req, err := c.newRequest("GET", "/info")
	if err != nil {
		return nil, nil, err
	}

	var buf bytes.Buffer

	resp, err := c.request(req, &buf)
	if err != nil {
		return nil, nil, err
	}

	data := &Info{}
	if err := json.Unmarshal(buf.Bytes(), data); err != nil {
		return nil, nil, err
	}
	return data, resp, nil
}
