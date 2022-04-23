package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/drgomesp/ethos/contracts"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/joho/godotenv/autoload"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"math/big"
	"os"
	"time"
)

var cfg EthosConfig

func init() {
	cfg = EthosConfig{
		EndpointJsonRPC:   os.Getenv("RPC_URL"),
		EndpointWebSocket: os.Getenv("WS_URL"),
	}

	// UNIX Time is faster and smaller than most timestamps
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC1123})
}

func main() {
	ctx := context.Background()
	rpc, err := ethclient.Dial(cfg.EndpointJsonRPC)
	if err != nil {
		log.Fatal().Err(err)
	}

	//privateKey, publicKey, _ := CreateWallet(ctx, err)
	//DeployContract(ctx, publicKey, err, rpc, privateKey)

	go GetBlocksPeriodically(ctx, rpc, cfg)

	ws, err := ethclient.Dial(cfg.EndpointWebSocket)
	if err != nil {
		log.Fatal().Err(err)
	}

	ListenForContractEvents(ctx, cfg, ws, "0xD7bEA2b69C7a1015aAdAA134e564eEe6d34149C0")
}

func DeployContract(ctx context.Context, publicKey *ecdsa.PublicKey, err error,
	rpc *ethclient.Client,
	privateKey *ecdsa.PrivateKey) {

	fromAddress := crypto.PubkeyToAddress(*publicKey)
	nonce, err := rpc.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal().Err(err)
	}

	gasPrice, err := rpc.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal().Err(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(
		privateKey,
		big.NewInt(cfg.ChainID),
	)
	if err != nil {
		log.Fatal().Err(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, _, err := contracts.DeployMain(auth, rpc)
	if err != nil {
		log.Error().Err(err).Msg("oops")
	}

	log.Debug().
		Str("address", address.Hex()).
		Interface("txn", tx).
		Msg("contract deployed")
}

func CreateWallet(ctx context.Context, err error) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal().Err(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal().Err(errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey"))
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	log.Debug().
		Str("private key", hexutil.Encode(privateKeyBytes)[2:]).
		Str("public key", hexutil.Encode(publicKeyBytes)[4:]).
		Msg("")

	return privateKey, publicKeyECDSA, nil
}

func ListenForContractEvents(ctx context.Context, cfg EthosConfig, client *ethclient.Client,
	contactAddr string) {
	contractAddress := common.HexToAddress(contactAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(ctx, query, logs)
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

func GetBlocksPeriodically(ctx context.Context, client *ethclient.Client, cfg EthosConfig) {
	for {
		if header := GetBlockHeaderOrError(ctx, client); header != nil {
			if block := GetBlockOrError(ctx, client, header); block != nil {
				LogBlock(block)
			}
		}

		time.Sleep(5 * time.Second)
	}
}

func GetBlockHeaderOrError(ctx context.Context, client *ethclient.Client) *types.Header {
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Error().Err(err)
	}
	return header
}

func GetBlockOrError(ctx context.Context, client *ethclient.Client,
	header *types.Header) *types.Block {
	blockNumber := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(ctx, blockNumber)
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
