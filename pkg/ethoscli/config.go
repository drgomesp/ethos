package ethoscli

var Config = DefaultEthosConfig()

type EthosConfig struct {
	ContractsDir      string `yaml:"contracts_dir"`
	ChainID           int64  `yaml:"chain_id"`
	EndpointJsonRPC   string `yaml:"endpoint_json_rpc"`
	EndpointWebSocket string `yaml:"endpoint_web_socket"`
}

func DefaultEthosConfig() *EthosConfig {
	return &EthosConfig{
		ContractsDir:      "./contracts",
		ChainID:           1337,
		EndpointJsonRPC:   "https://cloudflare-eth.com",
		EndpointWebSocket: "wss://rinkeby.infura.io/ws",
	}
}
