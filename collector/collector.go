package collector

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type Plex interface {
	GetVersion() (string, error)
	CurrentSessionsCount() (int, error)
	GetTranscodeSessions() (int, error)
	GetLibraries() ([]Library, error)
}

func NewPlexMediaServerCollector(server Plex) *PlexMediaServerCollector {
	return &PlexMediaServerCollector{
		MetricsInfo: prometheus.NewDesc("plex_info",
			"Plex media server information",
			[]string{"version"}, nil,
		),
		MetricsSessions: prometheus.NewDesc("plex_sessions_active_count",
			"Number of active Plex sessions",
			[]string{}, nil,
		),
		MetricsTranscodeSessions: prometheus.NewDesc("plex_transcode_sessions_active_count",
			"Number of active Plex transcoding sessions",
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
	MetricsInfo              *prometheus.Desc
	MetricsSessions          *prometheus.Desc
	MetricsTranscodeSessions *prometheus.Desc
	MetricsLibraries         *prometheus.Desc
	Server                   Plex
}

func (p *PlexMediaServerCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- p.MetricsInfo
	ch <- p.MetricsSessions
	ch <- p.MetricsTranscodeSessions
	ch <- p.MetricsLibraries
}

func (p *PlexMediaServerCollector) Collect(ch chan<- prometheus.Metric) {
	version, err := p.Server.GetVersion()
	if err != nil {
		log.Print(err)
	}

	ch <- prometheus.MustNewConstMetric(p.MetricsInfo, prometheus.GaugeValue, float64(1), version)

	sessions, err := p.Server.CurrentSessionsCount()
	if err != nil {
		log.Print(err)
	}

	ch <- prometheus.MustNewConstMetric(p.MetricsSessions, prometheus.GaugeValue, float64(sessions))

	transcodeSessions, err := p.Server.GetTranscodeSessions()
	if err != nil {
		log.Print(err)
	}

	ch <- prometheus.MustNewConstMetric(p.MetricsTranscodeSessions, prometheus.GaugeValue, float64(transcodeSessions))

	libraries, err := p.Server.GetLibraries()
	if err != nil {
		log.Print(err)
	}

	for _, library := range libraries {
		ch <- prometheus.MustNewConstMetric(p.MetricsLibraries, prometheus.GaugeValue, float64(library.Size), library.Name, library.Type)
	}
}
