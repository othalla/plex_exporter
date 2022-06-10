package collector

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
	// Responses that will be returned by the MockHTTPClient
	// Note that responses are consummed once returned
	responses     []*http.Response
	err           error
	responseIndex int
}

func NewMockHTTPClient(responses []*http.Response, err error) *MockHTTPClient {
	return &MockHTTPClient{responses: responses, err: err, responseIndex: 0}
}

func (c *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if c.err != nil {
		return nil, c.err
	}
	response := c.responses[c.responseIndex]
	c.responseIndex++
	return response, nil
}

func TestCollectorPlexServerGetVersion(t *testing.T) {
	responses := []*http.Response{
		{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"MediaContainer": {"version": "1.25.1"}}`)),
		},
	}
	client := NewMockHTTPClient(responses, nil)

	plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	version, err := plexServer.GetVersion()
	assert.NoError(t, err)
	assert.Equal(t, version, "1.25.1")
}

func TestPCollectorlexServerGetVersionBadJsonResponse(t *testing.T) {
	responses := []*http.Response{
		{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`malformed`)),
		},
	}
	client := NewMockHTTPClient(responses, nil)

	plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	_, err := plexServer.GetVersion()
	assert.NotNil(t, err)
}

func TestCollectorPlexServerGetVersionBadStatusCode(t *testing.T) {
	responses := []*http.Response{
		{
			StatusCode: 500,
		},
	}
	client := NewMockHTTPClient(responses, nil)

	plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	_, err := plexServer.GetVersion()
	assert.NotNil(t, err)
	assert.Equal(t, err, fmt.Errorf("got bad status code 500 from server"))
}

func TestCollectorPlexServerGetVersionHTTPRequestError(t *testing.T) {
	client := NewMockHTTPClient(nil, errors.New("http error"))

	plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	_, err := plexServer.GetVersion()
	assert.NotNil(t, err)
}

func TestCollectorPlexServerCurrentSessionsCount(t *testing.T) {

	responses := []*http.Response{
		{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"MediaContainer": {"size": 3}}`)),
		},
	}
	client := NewMockHTTPClient(responses, nil)

	plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	sessionCounter, _ := plexServer.CurrentSessionsCount()
	assert.Equal(t, sessionCounter, 3)
}

func TestPCollectorlexServerCurrentSessionsCountBadJsonResponse(t *testing.T) {
	responses := []*http.Response{
		{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`malformed`)),
		},
	}
	client := NewMockHTTPClient(responses, nil)

	plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	_, err := plexServer.CurrentSessionsCount()
	assert.NotNil(t, err)
}

func TestCollectorPlexServerCurrentSessionsCountBadStatusCode(t *testing.T) {
	responses := []*http.Response{
		{
			StatusCode: 500,
		},
	}
	client := NewMockHTTPClient(responses, nil)

	plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	_, err := plexServer.CurrentSessionsCount()
	assert.NotNil(t, err)
	assert.Equal(t, err, fmt.Errorf("Got bad status code 500 from server"))
}

func TestCollectorPlexServerCurrentSessionsCountHTTPRequestError(t *testing.T) {
	client := NewMockHTTPClient(nil, errors.New("http error"))

	plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	_, err := plexServer.CurrentSessionsCount()
	assert.NotNil(t, err)
}

func TestCollectorPlexServerGetTranscodeSessions(t *testing.T) {
	responses := []*http.Response{
		{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"MediaContainer": {"size": 2}}`)),
		},
	}
	client := NewMockHTTPClient(responses, nil)

	plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	sessionCounter, _ := plexServer.GetTranscodeSessions()
	assert.Equal(t, sessionCounter, 2)
}

func TestCollectorPlexServerGetTranscodeSessionsBadJsonResponse(t *testing.T) {
	responses := []*http.Response{
		{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`malformed`)),
		},
	}
	client := NewMockHTTPClient(responses, nil)

	plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	_, err := plexServer.GetTranscodeSessions()
	assert.NotNil(t, err)
}

func TestCollectorPlexServerGetTranscodeSessionsHTTPRequestError(t *testing.T) {
	client := NewMockHTTPClient(nil, errors.New("http error"))

	plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	_, err := plexServer.GetTranscodeSessions()
	assert.NotNil(t, err)
}

func TestCollectorPlexServerGetTranscodeSessionsCountBadStatusCode(t *testing.T) {
	responses := []*http.Response{
		{
			StatusCode: 500,
		},
	}
	client := NewMockHTTPClient(responses, nil)

	plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	_, err := plexServer.GetTranscodeSessions()
	assert.NotNil(t, err)
	assert.Equal(t, err, fmt.Errorf("Got bad status code 500 from server"))
}

func TestCollectorPlexServerGetLibrares(t *testing.T) {
	responses := []*http.Response{
		{
			StatusCode: http.StatusOK,
			Body: ioutil.NopCloser(bytes.NewBufferString(`
				{
					"MediaContainer": {
						"Directory": [
							{"key": "1", "title": "First", "type": "show"},
							{"key": "2", "title": "Another", "type": "film"}
						]
					}
				}`),
			),
		},
		{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"MediaContainer": {"size": 100}}`)),
		},
		{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString(`{"MediaContainer": {"size": 200}}`)),
		},
	}
	client := NewMockHTTPClient(responses, nil)

	plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
	libraries, _ := plexServer.GetLibraries()

	assert.Equal(t, libraries[0].Name, "First")
	assert.Equal(t, libraries[0].Type, "show")
	assert.Equal(t, libraries[0].Size, 100)

	assert.Equal(t, libraries[1].Name, "Another")
	assert.Equal(t, libraries[1].Type, "film")
	assert.Equal(t, libraries[1].Size, 200)
}

func TestCollectorPlexServerGetLibrariesBadJsonResponse(t *testing.T) {
	responseScenarios := [][]*http.Response{
		{
			&http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`malformed`)),
			},
			&http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`{"MediaContainer": {"size": 100}}`)),
			},
		},
		{
			&http.Response{
				StatusCode: http.StatusOK,
				Body: ioutil.NopCloser(bytes.NewBufferString(`
					{"MediaContainer": {"Directory": [{"key": "1", "title": "First", "type": "show"}]}}`)),
			},
			&http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`malformed`)),
			},
		},
	}

	for _, responseScenario := range responseScenarios {
		client := NewMockHTTPClient(responseScenario, nil)

		plexServer := PlexMediaServer{Address: "127.0.0.1", Port: 32400, Token: "auth-token", HTTPClient: client}
		_, err := plexServer.GetLibraries()

		assert.NotNil(t, err)
	}
}
