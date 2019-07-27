package plex

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RoundTripFunction func(req *http.Request) *http.Response

func (f RoundTripFunction) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(function RoundTripFunction) *http.Client {
	return &http.Client{
		Transport: function,
	}
}

func TestPlexServerCurrentSessionsCount(t *testing.T) {

	client := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, req.URL.String(), "https://127.0.0.1:32400/status/sessions")
		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"MediaContainer": {"size": 3}}`)),
		}
	})

	plexServer := PlexServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	sessionCounter := plexServer.CurrentSessionsCount()
	assert.Equal(t, sessionCounter, 3)
}
