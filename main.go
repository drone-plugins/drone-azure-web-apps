package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "azure web apps plugin"
	app.Usage = "azure web apps plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "username",
			Usage:  "username",
			EnvVar: "PLUGIN_USERNAME,AZURE_WEB_APPS_USERNAME",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "password",
			EnvVar: "PLUGIN_PASSWORD,AZURE_WEB_APPS_PASSWORD",
		},
		cli.StringFlag{
			Name:   "site",
			Usage:  "site",
			EnvVar: "PLUGIN_SITE,AZURE_WEB_APPS_SITE",
		},
		cli.StringFlag{
			Name:   "slot",
			Usage:  "slot",
			EnvVar: "PLUGIN_SLOT,AZURE_WEB_APPS_SLOT",
		},
		cli.BoolFlag{
			Name:   "force",
			Usage:  "force",
			EnvVar: "PLUGIN_FORCE",
		},
		cli.BoolFlag{
			Name:   "commit",
			Usage:  "commit",
			EnvVar: "PLUGIN_COMMIT",
		},
		cli.StringFlag{
			Name:   "commit.author.name",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.author.email",
			Usage:  "git author email",
			EnvVar: "DRONE_COMMIT_AUTHOR_EMAIL",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Commit: Commit{
			Author: Author{
				Name:  c.String("commit.author.name"),
				Email: c.String("commit.author.email"),
			},
		},
		Repo: Repo{
			Name: c.String("repo.name"),
		},
		Config: Config{
			Username: c.String("username"),
			Password: c.String("password"),
			Site:     c.String("site"),
			Slot:     c.String("slot"),
			Force:    c.Bool("force"),
			Commit:   c.Bool("commit"),
		},
	}

	if plugin.Config.Username == "" {
		return errors.New("Missing username")
	}

	if plugin.Config.Password == "" {
		return errors.New("Missing password")
	}

	return plugin.Exec()
}
