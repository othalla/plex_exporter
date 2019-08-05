package collector

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	pmsSessions = prometheus.NewDesc("plex_sessions_active_count",
		"Number of active Plex sessions",
		[]string{}, nil,
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

type Plex interface {
	CurrentSessionsCount() (int, error)
	GetLibraries() ([]Library, error)
}

type PlexMediaServerCollector struct {
	PlexServer Plex
}

func (p *PlexMediaServerCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- pmsSessions
}

func (p *PlexMediaServerCollector) Collect(ch chan<- prometheus.Metric) {
	sessions, err := p.PlexServer.CurrentSessionsCount()
	if err != nil {
		log.Print(err)
	}

	ch <- prometheus.MustNewConstMetric(pmsSessions, prometheus.GaugeValue, float64(sessions))

	libraries, _ := p.PlexServer.GetLibraries()

	for _, library := range libraries {
		plexLibrariesGauge.With(prometheus.Labels{"name": library.Name}).Set(float64(library.Size))
	}
}
