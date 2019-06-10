package ssllabs

import (
	"bytes"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	TrustStoreMozilla = "1"
	TrustStoreMacOS   = "2"
	TrustStoreAndroid = "3"
	TrustStoreJava    = "4"
	TrustStoreWindows = "5"
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

func (c *Client) GetRootCertsRaw(trustStore string) (*Certs, *http.Response, error) {
	param := &Param{
		Key:   "trustStore",
		Value: trustStore,
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

	certs, err := parseRootCerts(buf.String())
	return certs, resp, err
}

func parseRootCerts(raw string) (*Certs, error) {
	raw = strings.ReplaceAll(raw, "\r\n", "\n")
	lines := strings.Split(raw, "\n")

	ret := &Certs{
		TrustStore: strings.Split(lines[0], ": ")[1],
	}

	list := []Cert{}

	i := 1
	for i+6 < len(lines) {
		keyLength, err := strconv.Atoi(strings.Split(lines[i+3], ": ")[1])
		if err != nil {
			return nil, err
		}

		notBefore, err := time.Parse(time.UnixDate, strings.Split(lines[i+4], ": ")[1])
		if err != nil {
			return nil, err
		}

		notAfter, err := time.Parse(time.UnixDate, strings.Split(lines[i+5], ":  ")[1])
		if err != nil {
			return nil, err
		}

		ei := endIndexOfCert(lines, i)

		c := Cert{
			Name:        lines[i][2:],
			Subject:     map[string]string{},
			KeyType:     strings.Split(lines[i+2], ":   ")[1],
			KeyLength:   uint(keyLength),
			NotBefore:   notBefore,
			NotAfter:    notAfter,
			Certificate: strings.Join(lines[i+6:ei], "\n"),
		}
		list = append(list, c)
		i = ei + 2
	}

	ret.List = list

	return ret, nil
}

func endIndexOfCert(lines []string, current int) int {
	for i := current; i < len(lines); i++ {
		if strings.Contains(lines[i], "-----END CERTIFICATE-----") {
			return i
		}
	}
	return len(lines) - 1
}
