package main

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math/big"
	"os"
	"time"
)

const (
	EndpointLocal  = "http://localhost:8545"
	EndpointRemote = "https://cloudflare-eth.com"
)

func main() {
	cfg := &EthosConfig{
		EndpointJsonRPC: EndpointRemote,
	}

	// UNIX Time is faster and smaller than most timestamps
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123})

	client, err := ethclient.Dial(cfg.EndpointJsonRPC)
	if err != nil {
		log.Fatal().Err(err)
	}

	for {
		header := GetBlockHeaderOrPanic(client)
		block := GetBlockOrPanic(client, header)

		LogBlock(block)

		time.Sleep(1 * time.Second)
	}
}

func GetBlockHeaderOrPanic(client *ethclient.Client) *types.Header {
	header, err := client.HeaderByNumber(context.TODO(), nil)
	if err != nil {
		log.Fatal().Err(err)
	}
	return header
}

func GetBlockOrPanic(client *ethclient.Client, header *types.Header) *types.Block {
	blockNumber := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(context.TODO(), blockNumber)
	if err != nil {
		log.Fatal().Err(err)
	}

	return block
}

func LogBlock(block *types.Block) {
	log.Debug().
		Interface("height", block.Number().Uint64()).
		Interface("time", block.Time()).
		Interface("difficulty", block.Difficulty().Uint64()).
		Interface("hash", block.Hash().Hex()).
		Interface("transactions", len(block.Transactions())).
		Msg("block")
}
