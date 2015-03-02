package main

import (
	"github.com/codegangsta/cli"
	"github.com/jmervine/exec/v2"
	"log"
	"os"
	"strings"
)

func CliHandler(args []string) (port string, action []string, conditions Conditions) {
	app := cli.NewApp()
	app.Name = "go-git-fish"
	app.Version = "0.0.1"
	app.Author = "Joshua Mervine"
	app.Email = "joshua@mervine.net"
	app.Usage = "http listener and handler for github post commit hooks"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "command,c",
			Usage:  "command to execute *required*",
			EnvVar: "FISH_COMMAND",
		},
		cli.StringFlag{
			Name:   "port,p",
			Usage:  "http listener port",
			Value:  "8888",
			EnvVar: "FISH_PORT",
		},
		cli.BoolFlag{
			Name:   "verify",
			Usage:  "run command on startup to verify",
			EnvVar: "FISH_VERIFY",
		},
		cli.StringFlag{
			Name:   "secret,s",
			Usage:  "require a secret to be passed by github",
			EnvVar: "FISH_SECRET",
		},
		cli.StringFlag{
			Name:   "branches,b",
			Usage:  "filter on branch names, comma delim",
			EnvVar: "FISH_BRANCH",
		},
		cli.BoolFlag{
			Name:   "owner",
			Usage:  "filer, require repo owner push",
			EnvVar: "FISH_OWNER",
		},
		cli.BoolFlag{
			Name:   "admin",
			Usage:  "filer, require repo admin push",
			EnvVar: "FISH_ADMIN",
		},
		cli.BoolFlag{
			Name:   "master",
			Usage:  "filer, require branch be assigned as master branch",
			EnvVar: "FISH_MASTER",
		},
	}

	app.Action = func(c *cli.Context) {
		if c.String("command") == "" {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}

		action = safeSplit(c.String("command"))
		if c.Bool("verify") {
			if _, _, err := exec.ExecTee2(os.Stdout, os.Stderr, action[0], action[1:]...); err != nil {
				log.Fatal(err)
			}
		}

		if c.String("branches") != "" {
			branches := strings.Split(c.String("branches"), ",")
			for i, b := range branches {
				branches[i] = strings.TrimSpace(b)
			}

			conditions.Branches = branches
		}

		port = c.String("port")
		conditions.Secret = c.String("secret")
		conditions.Owner = c.Bool("owner")
		conditions.Admin = c.Bool("admin")
		conditions.Master = c.Bool("master")
	}

	app.Run(args)

	return
}
