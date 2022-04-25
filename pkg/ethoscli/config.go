package ethoscli

import "github.com/rs/zerolog"

var Config = DefaultEthosConfig()

const EthosFile = "ethos.yaml"

type EthosConfig struct {
	LogLevel             zerolog.Level `yaml:"log_level"`
	Compiler             string        `yaml:"compiler"`
	ContractsDir         string        `yaml:"contracts_dir"`
	ContractBindingsDir  string        `yaml:"contracts_bindings_dir"`
	ContractsBindingsPkg string        `yaml:"contracts_bindings_pkg"`
	BuildDir             string        `yaml:"build_dir"`
	ChainID              int64         `yaml:"chain_id"`
	EndpointJsonRPC      string        `yaml:"endpoint_json_rpc"`
	EndpointWebSocket    string        `yaml:"endpoint_web_socket"`
}

func DefaultEthosConfig() *EthosConfig {
	return &EthosConfig{
		LogLevel:             zerolog.InfoLevel,
		Compiler:             "solcjs",
		ContractsDir:         "./contracts",
		ContractBindingsDir:  "./pkg/contracts",
		ContractsBindingsPkg: "main",
		BuildDir:             "./build",
		ChainID:              1337,
		EndpointJsonRPC:      "https://cloudflare-eth.com",
		EndpointWebSocket:    "wss://rinkeby.infura.io/ws",
	}
}
