package main

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

type MockPlexServer struct {
	Sessions int
}

func (mps *MockPlexServer) CurrentSessionsCount() (int, error) {
	return mps.Sessions, nil
}

func TestSetPlexSessionsMetrics2(t *testing.T) {
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
