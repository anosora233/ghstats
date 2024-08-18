package main

import (
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
				Aliases:  []string{"u"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "repository",
				Aliases: []string{"r"},
			},
		},
		Action: func(c *cli.Context) error {
			username := c.String("username")
			repository := c.String("repository")
			if len(repository) == 0 {
				lib.DisplayRepositories(username)
			} else {
				lib.DisplayReleases(username, repository)
			}
			return nil
		},
	}

	app.Run(os.Args)
}
