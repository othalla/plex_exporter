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
const URLLibrarySectionsIDAll = "http://%s:%d/library/sections/%d/all"

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
	Key   int    `json:"key"`
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

// CLean model of Plex Media Server library
type Library struct {
	Name string
	Type string
	Size int
}

type Directory struct{}

type CollectorPlexServer struct {
	Address    string
	Port       int
	Token      string
	HTTPClient HTTPClient
}

func (ps *CollectorPlexServer) CurrentSessionsCount() (int, error) {
	url := fmt.Sprintf(URLSessions, ps.Address, ps.Port)

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("X-Plex-Token", ps.Token)
	request.Header.Add("Accept", "application/json")
	response, err := ps.HTTPClient.Do(request)
	if err != nil {
		return 0, err
	}
	if response.StatusCode != 200 {
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

func (ps *CollectorPlexServer) GetLibraries() []Library {
	URL := fmt.Sprintf(URLLibrarySections, ps.Address, ps.Port)

	request, _ := http.NewRequest("GET", URL, nil)
	request.Header.Add("X-Plex-Token", ps.Token)
	request.Header.Add("Accept", "application/json")
	response, _ := ps.HTTPClient.Do(request)

	body, _ := ioutil.ReadAll(response.Body)

	var librarySectionsContainer APILibrarySections

	json.Unmarshal([]byte(body), &librarySectionsContainer)

	var libraries []Library

	for _, directory := range librarySectionsContainer.MediaContainer.Directory {
		URL := fmt.Sprintf(URLLibrarySectionsIDAll, ps.Address, ps.Port, directory.Key)

		request, _ := http.NewRequest("GET", URL, nil)
		request.Header.Add("X-Plex-Token", ps.Token)
		request.Header.Add("Accept", "application/json")
		response, _ := ps.HTTPClient.Do(request)

		body, _ := ioutil.ReadAll(response.Body)

		var librarySectionsIDAllContainer APILibrarySectionsIDAll

		json.Unmarshal([]byte(body), &librarySectionsIDAllContainer)

		libraries = append(libraries, Library{Name: directory.Title, Type: directory.Type, Size: librarySectionsIDAllContainer.MediaContainer.Size})
	}
	return libraries
}
