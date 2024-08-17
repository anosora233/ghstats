package main

import (
	"log"
	"os"

	"github.com/anosora233/ghstats/lib"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ghstats",
		Usage: "Display repository release statistics",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "username",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "repository",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			username := c.String("username")
			repository := c.String("repository")
			lib.ShowReleases(username, repository)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
