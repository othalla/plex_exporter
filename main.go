package main

import (
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/othalla/plex_exporter/collector"
	"github.com/othalla/plex_exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load json configuration from `FILE`",
		},
	}

	app.Action = runExporter

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func runExporter(c *cli.Context) error {
	name := c.String("config")

	config, err := config.Load(name)
	if err != nil {
		return err
	}

	client := &http.Client{}

	server := &collector.PlexMediaServer{Address: config.Server.Address,
		Port:       config.Server.Port,
		Token:      config.Server.Token,
		HTTPClient: client,
	}
	collector := collector.NewPlexMediaServerCollector(server)

	prometheus.MustRegister(collector)

	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Starting exporter on port %d ...", config.Exporter.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Exporter.Port), nil)

	return nil
}
