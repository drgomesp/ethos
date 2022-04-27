package task

var Config = DefaultEthosConfig()

const EthosFile = "ethos.yaml"

type EthosConfig struct {
	Compiler             string `yaml:"compiler"`
	SourcesDir           string `yaml:"contracts_dir"`
	ContractBindingsDir  string `yaml:"contracts_bindings_dir"`
	ContractsBindingsPkg string `yaml:"contracts_bindings_pkg"`
	BuildDir             string `yaml:"build_dir"`
	ChainID              int64  `yaml:"chain_id"`
	EndpointJsonRPC      string `yaml:"endpoint_json_rpc"`
	EndpointWebSocket    string `yaml:"endpoint_web_socket"`
}

func DefaultEthosConfig() *EthosConfig {
	return &EthosConfig{
		Compiler:             "solc",
		SourcesDir:           "./contracts",
		ContractBindingsDir:  "./pkg/contracts",
		ContractsBindingsPkg: "main",
		BuildDir:             "./build",
		ChainID:              1337,
		EndpointJsonRPC:      "https://cloudflare-eth.com",
		EndpointWebSocket:    "wss://rinkeby.infura.io/ws",
	}
}
