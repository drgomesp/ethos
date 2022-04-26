package ethereum

type CompilerSettings struct {
	EVMVersion      string                         `json:"evmVersion"`
	Optimizer       OptimizerOptions               `json:"optimizer"`
	OutputSelection map[string]map[string][]string `json:"outputSelection"`
}
type OptimizerOptions struct {
	Enabled bool `json:"enabled"`
}
type CompilerSource struct {
	URLs []string `json:"urls"`
}

type CompilerOptions struct {
	Language string                    `json:"language"`
	Sources  map[string]CompilerSource `json:"sources"`
	Settings CompilerSettings          `json:"settings"`
}
