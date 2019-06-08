package ssllabs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptionBaseURL(t *testing.T) {
	cases := []struct {
		name  string
		input string
		want  string
	}{
		{name: "ValidFormat", input: "https://example.com/api", want: "https://example.com/api"},
		{name: "UnnecessarySlash", input: "https://example.com/api/", want: "https://example.com/api"},
		{name: "SpecificPort", input: "http://example.com:8080/api", want: "http://example.com:8080/api"},
		{name: "InvalidScheme", input: "foobar://example.com/api", want: ""},
		{name: "NoScheme", input: "example.com/api", want: ""},
		{name: "InvalidFormat", input: "http:example.com/api", want: ""},
		{name: "ControlCharacter", input: "https://exam\x00ple.com/api", want: ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			opt := OptionBaseURL(c.input)

			client, err := New(opt)
			if c.want == "" {
				assert.Error(t, err)
			} else {
				if assert.NoError(t, err) {
					got := client.baseURL
					assert.Equal(t, c.want, got)
				}
			}
		})
	}
}
