package ssllabstest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

const (
	Prefix = "/api/v3"
)

func NewServer(t *testing.T, r *mux.Router) *httptest.Server {
	t.Helper()

	return httptest.NewServer(r)
}

func NewRouter(t *testing.T, method string, path string, handler http.HandlerFunc) *mux.Router {
	t.Helper()

	r := mux.NewRouter()
	s := r.PathPrefix(Prefix).Subrouter()

	s.HandleFunc(path, handler).Methods(method)

	return r
}

func NewHandler(t *testing.T, response []byte) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("content-type", "application/json")
		_, _ = w.Write(response)
	}
}
