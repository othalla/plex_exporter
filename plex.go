package plex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const URLSessions = "https://%s:%d/status/sessions"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type SessionMediaContainer struct {
	SessionsSummary SessionsSummary `json:"MediaContainer"`
}

type SessionsSummary struct {
	Size int `json:"size"`
}

type PlexServer struct {
	Address string
	Port    int
	Token   string
}

func (ps *PlexServer) CurrentSessionsCount(HTTPClient HTTPClient) (int, error) {
	url := fmt.Sprintf(URLSessions, ps.Address, ps.Port)

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("X-Plex-Token", ps.Token)
	response, _ := HTTPClient.Do(request)
	if response.StatusCode != 200 {
		return 0, fmt.Errorf("Got bad status code %d from server", response.StatusCode)
	}

	body, _ := ioutil.ReadAll(response.Body)
	var sessionContainer SessionMediaContainer

	if err := json.Unmarshal([]byte(body), &sessionContainer); err != nil {
		return 0, err
	}

	return sessionContainer.SessionsSummary.Size, nil
}
