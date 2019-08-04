package collector

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	plexSessionsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "plex_media_server",
			Subsystem: "sessions",
			Name:      "active_total",
			Help:      "Total of actives sessions on remote plex media server",
		},
	)
	plexLibrariesGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "plex_media_server",
			Subsystem: "libraries",
			Name:      "media_count",
			Help:      "Total of media in a given library",
		},
		[]string{"name"},
	)
)

type CollectorPlex interface {
	CurrentSessionsCount() (int, error)
	GetLibraries() ([]Library, error)
}

type PlexExporter struct {
	PlexServer CollectorPlex
}

func (pe *PlexExporter) Describe(ch chan<- *prometheus.Desc) {
}

func (pe *PlexExporter) Collect(ch chan<- prometheus.Metric) {
	sessions, err := pe.PlexServer.CurrentSessionsCount()
	if err != nil {
		log.Print(err)
	}

	plexSessionsGauge.Set(float64(sessions))

	ch <- plexSessionsGauge

	libraries, _ := pe.PlexServer.GetLibraries()

	for _, library := range libraries {
		plexLibrariesGauge.With(prometheus.Labels{"name": library.Name}).Set(float64(library.Size))
	}
}
