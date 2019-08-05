package collector

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
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

func TestSetPlexLibrariesMetrics(t *testing.T) {
	mockServer := &MockPlexMediaServer{Libraries: []Library{
		{Name: "mylib", Type: "TV Shows", Size: 200},
		{Name: "anotherlib", Type: "Movie", Size: 500},
	}}
	collector := &PlexMediaServerCollector{PlexServer: mockServer}

	ch := make(chan prometheus.Metric)
	go func() {
		collector.Collect(ch)
		close(ch)
	}()

	for range ch {
	}

	gaugeOne, _ := plexLibrariesGauge.GetMetricWithLabelValues("mylib")
	assert.Equal(t, float64(200), testutil.ToFloat64(gaugeOne))
	gaugeTwo, _ := plexLibrariesGauge.GetMetricWithLabelValues("anotherlib")
	assert.Equal(t, float64(500), testutil.ToFloat64(gaugeTwo))
}

func TestExporterGetSessions(t *testing.T) {
	const metadata = `
# HELP plex_sessions_active_count Number of active Plex sessions
# TYPE plex_sessions_active_count Gauge
	`

	mockServer := &MockPlexMediaServer{Sessions: 17}
	collector := &PlexMediaServerCollector{PlexServer: mockServer}

	expected := `
plex_sessions_active_count 17
	`

	if err := testutil.CollectAndCompare(collector, strings.NewReader(metadata+expected)); err != nil {
		t.Errorf("unexpected collecting result:\n%s", err)
	}
}
