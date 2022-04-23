package main

type EthosConfig struct {
	ChainID           int64
	EndpointJsonRPC   string
	EndpointWebSocket string
}

func DefaultEthosConfig() EthosConfig {
	return EthosConfig{
		EndpointJsonRPC:   "https://cloudflare-eth.com",
		EndpointWebSocket: "wss://rinkeby.infura.io/ws",
	}
}
