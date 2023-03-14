package cmd

import (
	"github.com/urfave/cli/v2"
	"os"
)

var App = &cli.App{
	Name:     "golang-skeleton",
	Usage:    "golang-skeleton",
	Commands: []*cli.Command{},
	Action: func(context *cli.Context) error {
		return cli.ShowAppHelp(context)
	},
}

func Execute() error {
	return App.Run(os.Args)
}
