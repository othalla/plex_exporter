package plex

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
)

func TestSetPlexSessionsMetrics(t *testing.T) {
	SetPlexSessionsMetrics(10)

	assert.Equal(t, testutil.ToFloat64(plexSessionsGauge), float64(10))
}
