package main

import (
	"context"
	"github.com/drgomesp/ethos/pkg/ethoscli"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strconv"
	"time"
)

func init() {
	chainId, err := strconv.Atoi(os.Getenv("CHAIN_ID"))
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// UNIX Time is faster and smaller than most timestamps
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123})
	log.Logger = logger

	ethoscli.TestConfig = ethoscli.EthosConfig{
		ChainID:           int64(chainId),
		EndpointJsonRPC:   os.Getenv("RPC_URL"),
		EndpointWebSocket: os.Getenv("WS_URL"),
	}
}

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		Commands: []cli.Command{
			{
				Name:        "chain",
				ShortName:   "c",
				Usage:       "Start a local Ethereum blockchain node",
				UsageText:   "Usage text",
				Description: "DESC: Start a local Ethereum blockchain node",
				Action: func(ctx *cli.Context) error {
					return ethoscli.Chain(context.Background())
				},
			},
			{
				Name:      "test",
				ShortName: "t",
				Action: func(ctx *cli.Context) error {
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
