package collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TODO CHANGE INSECURE - do we have to query server directly? query it through plex.tv?
const URLSessions = "http://%s:%d/status/sessions"
const URLLibrarySections = "http://%s:%d/library/sections"
const URLLibrarySectionsIDAll = "http://%s:%d/library/sections/%s/all"
const URLTranscodeSessions = "http://%s:%d/transcode/sessions"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// API Response for /status/sections which give the number of current active sessions
type APIStatusSessions struct {
	MediaContainer APIStatusSessionsMediaContainer `json:"MediaContainer"`
}

type APIStatusSessionsMediaContainer struct {
	Size int `json:"size"`
}

// API Response for /library/sections which gives the library list
type APILibrarySections struct {
	MediaContainer APILibrarySectionsMediaContainer `json:"MediaContainer"`
}

type APILibrarySectionsMediaContainer struct {
	Directory []APILibrarySectionsDirectory `json:"Directory"`
}

type APILibrarySectionsDirectory struct {
	Key   string `json:"key"`
	Type  string `json:"type"`
	Title string `json:"title"`
}

// API Response for /library/sections/id/all which gives the number of items by library
type APILibrarySectionsIDAll struct {
	MediaContainer APILibrarySectionsIDAllMediaContainer `json:"MediaContainer"`
}

type APILibrarySectionsIDAllMediaContainer struct {
	Size int `json:"size"`
}

// API Response for /transcode/sessions which give the number of current active transcoding sessions
type APITranscodeSessions struct {
	MediaContainer APITranscodeSessionsMediaContainer `json:"MediaContainer"`
}

type APITranscodeSessionsMediaContainer struct {
	Size int `json:"size"`
}

// CLean model of Plex Media Server library
type Library struct {
	Name string
	Type string
	Size int
}

type Directory struct{}

type PlexMediaServer struct {
	Address    string
	Port       int
	Token      string
	HTTPClient HTTPClient
}

func (p *PlexMediaServer) CurrentSessionsCount() (int, error) {
	url := fmt.Sprintf(URLSessions, p.Address, p.Port)

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("X-Plex-Token", p.Token)
	request.Header.Add("Accept", "application/json")
	response, err := p.HTTPClient.Do(request)
	if err != nil {
		return 0, err
	}
	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Got bad status code %d from server", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}
	var sessionContainer APIStatusSessions

	if err := json.Unmarshal([]byte(body), &sessionContainer); err != nil {
		return 0, err
	}

	return sessionContainer.MediaContainer.Size, nil
}

// GetTranscodeSessions returns the number of current transcoding sessions
func (p *PlexMediaServer) GetTranscodeSessions() (int, error) {
	url := fmt.Sprintf(URLTranscodeSessions, p.Address, p.Port)
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("X-Plex-Token", p.Token)
	request.Header.Add("Accept", "application/json")
	response, err := p.HTTPClient.Do(request)
	if err != nil {
		return 0, err
	}
	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Got bad status code %d from server", response.StatusCode)
	}

	body, _ := ioutil.ReadAll(response.Body)

	var transCodeSessionsContainer APITranscodeSessions

	if err := json.Unmarshal([]byte(body), &transCodeSessionsContainer); err != nil {
		return 0, err
	}

	return transCodeSessionsContainer.MediaContainer.Size, nil
}

// GetLibraries returns the list of all libraries present on the Plex media object
func (p *PlexMediaServer) GetLibraries() ([]Library, error) {
	URL := fmt.Sprintf(URLLibrarySections, p.Address, p.Port)

	request, _ := http.NewRequest(http.MethodGet, URL, nil)
	request.Header.Add("X-Plex-Token", p.Token)
	request.Header.Add("Accept", "application/json")
	response, _ := p.HTTPClient.Do(request)

	body, _ := ioutil.ReadAll(response.Body)

	var librarySectionsContainer APILibrarySections

	err := json.Unmarshal([]byte(body), &librarySectionsContainer)
	if err != nil {
		return nil, err
	}

	var libraries []Library

	for _, directory := range librarySectionsContainer.MediaContainer.Directory {
		URL := fmt.Sprintf(URLLibrarySectionsIDAll, p.Address, p.Port, directory.Key)

		request, _ := http.NewRequest(http.MethodGet, URL, nil)
		request.Header.Add("X-Plex-Token", p.Token)
		request.Header.Add("Accept", "application/json")
		response, _ := p.HTTPClient.Do(request)

		body, _ := ioutil.ReadAll(response.Body)

		var librarySectionsIDAllContainer APILibrarySectionsIDAll

		err := json.Unmarshal([]byte(body), &librarySectionsIDAllContainer)
		if err != nil {
			return nil, err
		}

		libraries = append(libraries, Library{Name: directory.Title, Type: directory.Type, Size: librarySectionsIDAllContainer.MediaContainer.Size})
	}
	return libraries, nil
}
