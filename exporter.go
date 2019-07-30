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
		},
	)
)

type CollectorPlex interface {
	CurrentSessionsCount() int
}

type PlexExporter struct {
	PlexServer CollectorPlex
}

func (pe *PlexExporter) Describe(ch chan<- *prometheus.Desc) {
}

func (pe *PlexExporter) Collect(ch chan<- prometheus.Metric) {
	sessions := pe.PlexServer.CurrentSessionsCount()
	plexSessionsGauge.Set(float64(sessions))
	ch <- plexSessionsGauge

}
