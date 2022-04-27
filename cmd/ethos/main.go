package main

import (
	"context"
	"github.com/drgomesp/ethos/pkg/task"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var Version string
var Build string

func init() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	if Version != "" {
		log.Info().Msgf("build.Version: %s\t", Version)
	}

	if Build != "" {
		log.Info().Msgf("build.Build: %s\t", Build)
	}

	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Initialize the Ethos config file",
				Action: func(ctx *cli.Context) error {
					return task.Init(context.Background())
				},
			},
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Build your Solidity contract source files",
				Action: func(ctx *cli.Context) error {
					MustLoadConfig()
					return task.Build(context.Background())
				},
			},
			{
				Name:    "chain",
				Aliases: []string{"c"},
				Usage:   "Start a local Ethereum blockchain node",
				Action: func(ctx *cli.Context) error {
					MustLoadConfig()
					return task.Chain(context.Background())
				},
			},
			{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "Some shenanigans...",
				Action: func(ctx *cli.Context) error {
					MustLoadConfig()
					return task.Test(context.Background())
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run")
	}
}

func MustLoadConfig() {
	f, err := ioutil.ReadFile("./ethos.yaml")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	err = yaml.Unmarshal(f, &task.Config)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}
}
