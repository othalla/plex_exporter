package plex

import (
	"fmt"
	"log"
	"os"

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
	fmt.Printf("param: %s\n", name)
	return nil
}
