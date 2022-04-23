package main

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math/big"
	"os"
	"time"
)

func main() {
	cfg := EthosConfig{
		EndpointJsonRPC:   os.Getenv("RPC_URL"),
		EndpointWebSocket: os.Getenv("WS_URL"),
	}

	// UNIX Time is faster and smaller than most timestamps
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123})

	rpc, err := ethclient.Dial(cfg.EndpointJsonRPC)
	if err != nil {
		log.Fatal().Err(err)
	}

	go GetBlocksPeriodically(rpc, cfg)

	ws, err := ethclient.Dial(cfg.EndpointWebSocket)
	if err != nil {
		log.Fatal().Err(err)
	}

	ListenForContractEvents(cfg, ws, "0xD7bEA2b69C7a1015aAdAA134e564eEe6d34149C0")
}

func ListenForContractEvents(cfg EthosConfig, client *ethclient.Client, contactAddr string) {
	contractAddress := common.HexToAddress(contactAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.TODO(), query, logs)
	if err != nil {
		log.Fatal().Err(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal().Err(err)
		case vLog := <-logs:
			j, e := vLog.MarshalJSON()
			if e != nil {
				log.Fatal().Err(err)
			}
			log.Debug().RawJSON("event", j).Msg("")
		}
	}
}

func GetBlocksPeriodically(client *ethclient.Client, cfg EthosConfig) {
	for {
		if header := GetBlockHeaderOrError(client); header != nil {
			if block := GetBlockOrError(client, header); block != nil {
				LogBlock(block)
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func GetBlockHeaderOrError(client *ethclient.Client) *types.Header {
	header, err := client.HeaderByNumber(context.TODO(), nil)
	if err != nil {
		log.Error().Err(err)
	}
	return header
}

func GetBlockOrError(client *ethclient.Client, header *types.Header) *types.Block {
	blockNumber := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(context.TODO(), blockNumber)
	if err != nil {
		log.Error().Err(err)
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
