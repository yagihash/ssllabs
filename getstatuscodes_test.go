package ssllabs

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yagihash/ssllabs/ssllabstest"
)

func TestGetStatusCodes(t *testing.T) {
	t.Parallel()
	r := ssllabstest.NewRouterGetStatusCodes(t)
	ts := ssllabstest.NewServer(t, r)

	t.Run("TestGetStatusCodes", func(t *testing.T) {
		client, err := New(OptionBaseURL(ts.URL + ssllabstest.Prefix))
		if assert.NoError(t, err) {
			want := &StatusCodes{}
			_ = json.Unmarshal(ssllabstest.SampleResponseGetStatusCodesBytes(), want)
			got, _, err := client.GetStatusCodes()
			if assert.NoError(t, err) {
				assert.Equal(t, want, got)
			}
		}
	})
}
