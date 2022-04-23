package main

import (
	"context"
	"github.com/drgomesp/ethos/pkg/ethoscli"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

func init() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Caller().Logger()
}

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Commands: []cli.Command{
			{
				Name:      "chain",
				ShortName: "c",
				Usage:     "Start a local Ethereum blockchain node",
				Action: func(ctx *cli.Context) error {
					return ethoscli.Chain(context.Background())
				},
			},
			{
				Name:      "test",
				ShortName: "t",
				Usage:     "start with the shenanigans...",
				Action: func(ctx *cli.Context) error {
					cfg, err := LoadConfigFromYaml()
					if err != nil {
						return err
					}

					ethoscli.TestConfig = cfg

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

func LoadConfigFromYaml() (ethoscli.EthosConfig, error) {
	f, err := ioutil.ReadFile("./ethos.yaml")
	if err != nil {
		return ethoscli.NilConfig, err
	}

	var cfg ethoscli.EthosConfig
	err = yaml.Unmarshal(f, &cfg)
	if err != nil {
		return ethoscli.NilConfig, err
	}

	return cfg, nil
}
