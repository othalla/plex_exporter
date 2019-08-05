package collector

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

type MockPlexMediaServer struct {
	Sessions  int
	Libraries []Library
}

func (mps *MockPlexMediaServer) CurrentSessionsCount() (int, error) {
	return mps.Sessions, nil
}

func (mps *MockPlexMediaServer) GetLibraries() ([]Library, error) {
	return mps.Libraries, nil
}

func TestPlexMediaServerCollectorMetricsSessions(t *testing.T) {

	mockServer := &MockPlexMediaServer{Sessions: 17}
	collector := NewPlexMediaServerCollector(mockServer)

	expected := `
# HELP plex_sessions_active_count Number of active Plex sessions
# TYPE plex_sessions_active_count Gauge
plex_sessions_active_count 17
	`

	if err := testutil.CollectAndCompare(collector, strings.NewReader(expected)); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}
}

func TestPlexMediaServerCollectorMetricsLibraries(t *testing.T) {
	mockServer := &MockPlexMediaServer{Libraries: []Library{
		{Name: "mylib", Type: "TV Shows", Size: 200},
		{Name: "anotherlib", Type: "Movie", Size: 340},
	}}
	collector := NewPlexMediaServerCollector(mockServer)

	expected := `
# HELP plex_media_server_library_media_count Number of medias in a plex library
# TYPE plex_media_server_library_media_count Gauge
plex_media_server_library_media_count{name="mylib", type="TV Shows"} 200
plex_media_server_library_media_count{name="anotherlib", type="Movie"} 340
# HELP plex_sessions_active_count Number of active Plex sessions
# TYPE plex_sessions_active_count Gauge
plex_sessions_active_count 0
	`

	if err := testutil.CollectAndCompare(collector, strings.NewReader(expected)); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}
}
