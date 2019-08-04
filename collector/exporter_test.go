package collector

import (
	"fmt"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

type MockPlexServer struct {
	Sessions  int
	Libraries []Library
}

func (mps *MockPlexServer) CurrentSessionsCount() (int, error) {
	return mps.Sessions, nil
}

func (mps *MockPlexServer) GetLibraries() ([]Library, error) {
	return mps.Libraries, nil
}

func TestSetPlexSessionsMetrics(t *testing.T) {
	ps := &MockPlexServer{Sessions: 17}
	pe := &PlexExporter{PlexServer: ps}

	ch := make(chan prometheus.Metric)
	go func() {
		pe.Collect(ch)
		close(ch)
	}()

	for range ch {
	}

	assert.Equal(t, float64(17), testutil.ToFloat64(plexSessionsGauge))
}

func TestSetPlexLibrariesMetrics(t *testing.T) {
	ps := &MockPlexServer{Libraries: []Library{
		{Name: "mylib", Type: "TV Shows", Size: 200},
		{Name: "anotherlib", Type: "Movie", Size: 500},
	}}
	pe := &PlexExporter{PlexServer: ps}

	ch := make(chan prometheus.Metric)
	go func() {
		pe.Collect(ch)
		close(ch)
	}()

	for range ch {
	}

	gauge, _ := plexLibrariesGauge.GetMetricWithLabelValues("mylib")
	fmt.Println(testutil.ToFloat64(gauge))
}
