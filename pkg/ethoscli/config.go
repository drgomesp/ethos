package ethoscli

var Config = DefaultEthosConfig()

const EthosFile = "ethos.yaml"

type EthosConfig struct {
	Compiler          string `yaml:"compiler"`
	ContractsDir      string `yaml:"contracts_dir"`
	BuildDir          string `yaml:"build_dir"`
	ChainID           int64  `yaml:"chain_id"`
	EndpointJsonRPC   string `yaml:"endpoint_json_rpc"`
	EndpointWebSocket string `yaml:"endpoint_web_socket"`
}

func DefaultEthosConfig() *EthosConfig {
	return &EthosConfig{
		Compiler:          "solcjs",
		ContractsDir:      "./contracts",
		BuildDir:          "./build",
		ChainID:           1337,
		EndpointJsonRPC:   "https://cloudflare-eth.com",
		EndpointWebSocket: "wss://rinkeby.infura.io/ws",
	}
}
