package ssllabs

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yagihash/ssllabs/ssllabstest"
)

func TestInfo(t *testing.T) {
	t.Parallel()

	r := ssllabstest.NewRouterInfo(t)
	ts := ssllabstest.NewServer(t, r)

	t.Run("TestInfo", func(t *testing.T) {
		client, err := New(OptionBaseURL(ts.URL + ssllabstest.Prefix))
		if assert.NoError(t, err) {
			want := &Info{}
			_ = json.Unmarshal(ssllabstest.SampleResponseInfoBytes(), want)
			got, _, err := client.Info()
			if assert.NoError(t, err) {
				assert.Equal(t, want, got)
			}
		}
	})
}
