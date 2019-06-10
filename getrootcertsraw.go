package ssllabs

import (
	"bytes"
	"net/http"
	"time"
)

const (
	TrustStoreMozilla = 1
	TrustStoreMacOS   = 2
	TrustStoreAndroid = 3
	TrustStoreJava    = 4
	TrustStoreWindows = 5
)

type Certs struct {
	TrustStore string
	List       []Cert
}

type Cert struct {
	Name        string
	Subject     map[string]string
	KeyType     string
	KeyLength   uint
	NotBefore   time.Time
	NotAfter    time.Time
	Certificate string
}

func (c *Client) GetRootCertsRaw(trustStore uint8) (*Certs, *http.Response, error) {
	param := &Param{
		Key:   "trustStore",
		Value: string(trustStore),
	}
	req, err := c.newRequest("GET", "/getRootCertsRaw", param)
	if err != nil {
		return nil, nil, err
	}

	var buf bytes.Buffer

	resp, err := c.request(req, &buf)
	if err != nil {
		return nil, nil, err
	}

	return c.parseRootCerts(buf.String()), resp, nil
}

func (c *Client) parseRootCerts(raw string) *Certs {
	return &Certs{}
}
