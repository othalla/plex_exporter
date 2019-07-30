package plex

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHTTPClient struct {
	response *http.Response
	err      error
}

func (c *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.response, c.err
}

func TestCollectorPlexServerCurrentSessionsCount(t *testing.T) {

	client := MockHTTPClient{
		response: &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"MediaContainer": {"size": 3}}`)),
		},
		err: nil,
	}

	plexServer := CollectorPlexServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: &client}
	sessionCounter, _ := plexServer.CurrentSessionsCount()
	assert.Equal(t, sessionCounter, 3)
}

func TestPCollectorlexServerCurrentSessionsCountBadJsonResponse(t *testing.T) {
	client := MockHTTPClient{
		response: &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`malformed`)),
		},
		err: nil,
	}

	plexServer := CollectorPlexServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: &client}
	_, err := plexServer.CurrentSessionsCount()
	assert.NotNil(t, err)
}

func TestCollectorPlexServerCurrentSessionsCountBadStatusCode(t *testing.T) {
	client := MockHTTPClient{
		response: &http.Response{
			StatusCode: 500,
		},
		err: nil,
	}
	plexServer := CollectorPlexServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: &client}
	_, err := plexServer.CurrentSessionsCount()
	assert.NotNil(t, err)
	assert.Equal(t, err, fmt.Errorf("Got bad status code 500 from server"))
}

func TestCollectorPlexServerCurrentSessionsCountHTTPRequestError(t *testing.T) {
	client := MockHTTPClient{
		response: nil,
		err:      errors.New("http error"),
	}
	plexServer := CollectorPlexServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: &client}
	_, err := plexServer.CurrentSessionsCount()
	assert.NotNil(t, err)
}
