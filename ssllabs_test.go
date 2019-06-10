package ssllabs

import (
	"net/http"
	"net/http/cookiejar"
	"testing"

	"golang.org/x/net/publicsuffix"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	assert.NoError(t, err)
	defaultClient := &Client{
		baseURL: BaseURL,
		httpClient: &http.Client{
			Jar: jar,
		},
	}

	t.Run("Default", func(t *testing.T) {
		want := defaultClient
		got, err := New()
		assert.NoError(t, err)
		assert.Equal(t, want, got)
	})

	cases := []struct {
		name  string
		input Option
		want  *Client
	}{
		{name: "OptionBaseURL", input: OptionBaseURL("https://example.com/test"), want: &Client{
			baseURL: "https://example.com/test",
			httpClient: &http.Client{
				Jar: jar,
			},
		}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := New(c.input)
			if assert.NoError(t, err) {
				assert.Equal(t, c.want, got)
			}
		})
	}
}
