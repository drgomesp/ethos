package ethoscli

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
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"math/big"
	"time"
)

var TestConfig EthosConfig

func Test(ctx context.Context) error {
	client, err := ethclient.Dial(TestConfig.EndpointJsonRPC)
	if err != nil {
		log.Fatal().Err(err).Msg("rpc connection failed")
	}

	address, privateKey, publicKey, err := CreateWallet(ctx, err)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create wallet")
	}
	go GetBalancePeriodically(ctx, address, client, privateKey, publicKey)
	go GetBlocksPeriodically(ctx, client)

	var ws *ethclient.Client
	ws, err = ethclient.Dial(TestConfig.EndpointWebSocket)
	if err != nil {
		log.Fatal().Err(err).Msg("websocket connection failed")

	}

	return ListenForContractEvents(ctx, ws, "0xD7bEA2b69C7a1015aAdAA134e564eEe6d34149C0")
}

func CreateWallet(ctx context.Context, err error) (
	common.Address,
	*ecdsa.PrivateKey,
	*ecdsa.PublicKey,
	error,
) {
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
	publicKeyAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	log.Debug().
		Str("private key", hexutil.Encode(privateKeyBytes)).
		Str("public key", hexutil.Encode(publicKeyBytes)).
		Str("address", publicKeyAddress.String()).
		Msg("")

	return publicKeyAddress, privateKey, publicKeyECDSA, nil
}

func GetBalancePeriodically(
	ctx context.Context,
	address common.Address,
	client *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	publicKey *ecdsa.PublicKey,
) {
	for {
		balance, err := client.BalanceAt(ctx, address, nil)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to get balance")
		}

		log.Debug().Int64("balance", balance.Int64()).Msg("")
		time.Sleep(5 * time.Second)

		if balance.Int64() > 0 {
			DeployContract(ctx, client, privateKey, publicKey)
		}
	}
}

func DeployContract(
	ctx context.Context,
	rpc *ethclient.Client,
	privateKey *ecdsa.PrivateKey,
	publicKey *ecdsa.PublicKey,
) {

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
		big.NewInt(TestConfig.ChainID),
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

func ListenForContractEvents(
	ctx context.Context,
	client *ethclient.Client,
	contactAddr string,
) error {
	contractAddress := common.HexToAddress(contactAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return err
	}

	for {
		select {
		case err := <-sub.Err():
			return err
		case vLog := <-logs:
			j, e := vLog.MarshalJSON()
			if e != nil {
				return err
			}

			log.Debug().RawJSON("event", j).Msg("")
		}
	}
}

func GetBlocksPeriodically(ctx context.Context, client *ethclient.Client) {
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
