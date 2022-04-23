package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math/big"
	"os"
	"time"
)

const (
	Endpoint       = EndpointRemote
	EndpointLocal  = "http://localhost:8545"
	EndpointRemote = "https://cloudflare-eth.com"
)

func main() {
	// UNIX Time is faster and smaller than most timestamps
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123})

	client, err := ethclient.Dial(Endpoint)
	if err != nil {
		log.Fatal().Err(err)
	}

	ctx := context.Background()

	for {
		header, err := client.HeaderByNumber(ctx, nil)
		if err != nil {
			log.Fatal().Err(err)
		}

		blockNumber := big.NewInt(header.Number.Int64())
		block, err := client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			log.Fatal().Err(err)
		}

		log.Debug().Interface("number", block.Number().Uint64()).
			Interface("time", block.Time()).
			Interface("difficulty", block.Difficulty().Uint64()).
			Interface("hash", block.Hash().Hex()).
			Interface("transactions", len(block.Transactions())).
			Msg("received block")

		time.Sleep(1 * time.Second)
	}
}
