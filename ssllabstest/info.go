package ssllabstest

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func NewRouterInfo(t *testing.T) *mux.Router {
	t.Helper()

	h := NewHandler(t, SampleResponseInfoBytes())
	r := NewRouter(t, http.MethodGet, "/info", h)

	return r
}

func SampleResponseInfoBytes() []byte {
	return []byte(`
{
  "engineVersion": "1.35.1",
  "criteriaVersion": "2009p",
  "maxAssessments": 25,
  "currentAssessments": 0,
  "newAssessmentCoolOff": 1000,
  "messages": [
    "This assessment service is provided free of charge by Qualys SSL Labs, subject to our terms and conditions: https://www.ssllabs.com/about/terms.html"
  ]
}`)
}
