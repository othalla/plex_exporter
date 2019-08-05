package collector

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type Plex interface {
	CurrentSessionsCount() (int, error)
	GetLibraries() ([]Library, error)
}

func NewPlexMediaServerCollector(server Plex) *PlexMediaServerCollector {
	return &PlexMediaServerCollector{MetricsSessions: prometheus.NewDesc("plex_sessions_active_count",
		"Number of active Plex sessions",
		[]string{}, nil,
	),
		MetricsLibraries: prometheus.NewDesc("plex_media_server_library_media_count",
			"Number of medias in a plex library",
			[]string{"name", "type"},
			nil,
		),
		Server: server,
	}
}

type PlexMediaServerCollector struct {
	MetricsSessions  *prometheus.Desc
	MetricsLibraries *prometheus.Desc
	Server           Plex
}

func (p *PlexMediaServerCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- p.MetricsSessions
	ch <- p.MetricsLibraries
}

func (p *PlexMediaServerCollector) Collect(ch chan<- prometheus.Metric) {
	sessions, err := p.Server.CurrentSessionsCount()
	if err != nil {
		log.Print(err)
	}

	ch <- prometheus.MustNewConstMetric(p.MetricsSessions, prometheus.GaugeValue, float64(sessions))

	libraries, err := p.Server.GetLibraries()
	if err != nil {
		log.Print(err)
	}

	for _, library := range libraries {
		ch <- prometheus.MustNewConstMetric(p.MetricsLibraries, prometheus.GaugeValue, float64(library.Size), library.Name, library.Type)
	}
}
