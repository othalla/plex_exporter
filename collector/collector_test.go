package collector

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

type MockPlexMediaServer struct {
	Sessions          int
	TranscodeSessions int
	Libraries         []Library
	Version           string
}

func (mps *MockPlexMediaServer) GetVersion() (string, error) {
	return mps.Version, nil
}

func (mps *MockPlexMediaServer) CurrentSessionsCount() (int, error) {
	return mps.Sessions, nil
}

func (mps *MockPlexMediaServer) GetTranscodeSessions() (int, error) {
	return mps.TranscodeSessions, nil
}

func (mps *MockPlexMediaServer) GetLibraries() ([]Library, error) {
	return mps.Libraries, nil
}

func TestPlexMediaServerCollectorMetricsInfo(t *testing.T) {

	mockServer := &MockPlexMediaServer{Version: "1.25.0"}
	collector := NewPlexMediaServerCollector(mockServer)

	expected := `
# HELP plex_info Plex media server information
# TYPE plex_info Gauge
plex_info{version="1.25.0"} 1
# HELP plex_sessions_active_count Number of active Plex sessions
# TYPE plex_sessions_active_count Gauge
plex_sessions_active_count 0
# HELP plex_transcode_sessions_active_count Number of active Plex transcoding sessions
# TYPE plex_transcode_sessions_active_count Gauge
plex_transcode_sessions_active_count 0
	`

	if err := testutil.CollectAndCompare(collector, strings.NewReader(expected)); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}
}

func TestPlexMediaServerCollectorMetricsSessions(t *testing.T) {

	mockServer := &MockPlexMediaServer{Sessions: 17}
	collector := NewPlexMediaServerCollector(mockServer)

	expected := `
# HELP plex_info Plex media server information
# TYPE plex_info Gauge
plex_info{version=""} 1
# HELP plex_sessions_active_count Number of active Plex sessions
# TYPE plex_sessions_active_count Gauge
plex_sessions_active_count 17
# HELP plex_transcode_sessions_active_count Number of active Plex transcoding sessions
# TYPE plex_transcode_sessions_active_count Gauge
plex_transcode_sessions_active_count 0
	`

	if err := testutil.CollectAndCompare(collector, strings.NewReader(expected)); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}
}

func TestPlexMediaServerCollectorMetricsTranscodeSessions(t *testing.T) {

	mockServer := &MockPlexMediaServer{Sessions: 10, TranscodeSessions: 8}
	collector := NewPlexMediaServerCollector(mockServer)

	expected := `
# HELP plex_info Plex media server information
# TYPE plex_info Gauge
plex_info{version=""} 1
# HELP plex_sessions_active_count Number of active Plex sessions
# TYPE plex_sessions_active_count Gauge
plex_sessions_active_count 10
# HELP plex_transcode_sessions_active_count Number of active Plex transcoding sessions
# TYPE plex_transcode_sessions_active_count Gauge
plex_transcode_sessions_active_count 8
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
# HELP plex_info Plex media server information
# TYPE plex_info Gauge
plex_info{version=""} 1
# HELP plex_media_server_library_media_count Number of medias in a plex library
# TYPE plex_media_server_library_media_count Gauge
plex_media_server_library_media_count{name="mylib", type="TV Shows"} 200
plex_media_server_library_media_count{name="anotherlib", type="Movie"} 340
# HELP plex_sessions_active_count Number of active Plex sessions
# TYPE plex_sessions_active_count Gauge
plex_sessions_active_count 0
# HELP plex_transcode_sessions_active_count Number of active Plex transcoding sessions
# TYPE plex_transcode_sessions_active_count Gauge
plex_transcode_sessions_active_count 0
	`

	if err := testutil.CollectAndCompare(collector, strings.NewReader(expected)); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}
}
