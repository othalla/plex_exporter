package plex

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	plexSessionsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "plex_media_server",
			Subsystem: "sessions",
			Name:      "current_active",
			Help:      "Total of actives sessions on remote plex media server",
		})
)

func SetPlexSessionsMetrics(sessions float64) {
	plexSessionsGauge.Set(sessions)
}
