package main

import (
	"context"
	"github.com/drgomesp/ethos/pkg/ethoscli"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()
}

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Commands: []*cli.Command{
			{
				Name:    "init",
				Aliases: []string{"i"},
				Usage:   "Initialize the Ethos config file",
				Action: func(ctx *cli.Context) error {
					return ethoscli.Init(context.Background())
				},
			},
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Build your Solitidy contract source files",
				Action: func(ctx *cli.Context) error {
					MustLoadConfig()
					return ethoscli.Build(context.Background())
				},
			},
			{
				Name:    "chain",
				Aliases: []string{"c"},
				Usage:   "Start a local Ethereum blockchain node",
				Action: func(ctx *cli.Context) error {
					MustLoadConfig()
					return ethoscli.Chain(context.Background())
				},
			},
			{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "Some shenanigans...",
				Action: func(ctx *cli.Context) error {
					MustLoadConfig()
					return ethoscli.Test(context.Background())
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to run")
	}
}

func LoadConfigFromYaml() (*ethoscli.EthosConfig, error) {
	f, err := ioutil.ReadFile("./ethos.yaml")
	if err != nil {
		return nil, err
	}

	var cfg ethoscli.EthosConfig
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func MustLoadConfig() {
	if loadedCfg, err := LoadConfigFromYaml(); err != nil {
		log.Fatal().Err(err).Msg("failed to load config file (ethos.yaml)")
	} else {
		ethoscli.Config = loadedCfg
	}
}
