package compiler

import "github.com/ethereum/go-ethereum/accounts/abi"

type (
	Lang            string
	SourceName      string
	ContractName    string
	LibraryName     string
	MethodSignature string
)

const (
	LangSolidity = Lang("Solidity")
)

type Source struct {
	Name    SourceName `json:"name"`
	Content string     `json:"content"`
}

type OutputSelectionMap map[SourceName]map[ContractName][]string

type Settings struct {
	Optimizer struct {
		Enabled bool `json:"enabled"`
	} `json:"optimizer"`
	OutputSelection OutputSelectionMap     `json:"output_selection"`
	EVMVersion      string                 `json:"evm_version"`
	Libraries       map[LibraryName]string `json:"libraries"`
}

type Input struct {
	Language string            `json:"language"`
	Sources  map[string]Source `json:"sources"`
}

type OutputSource struct {
	ID  int
	AST interface{}
}

type OutputByteCode struct {
	Object              string
	Opcodes             string
	SourceMap           string
	LinkReferences      map[SourceName]map[LibraryName][]string
	ImmutableReferences map[string][]string
}

type OutputContract struct {
	ABI abi.ABI
	EVM struct {
		ByteCode          OutputByteCode
		DeployedByteCode  OutputByteCode
		MethodIdentifiers map[MethodSignature]string
	}
}

type Output struct {
	Sources   map[SourceName]OutputSource
	Contracts map[SourceName]map[ContractName]OutputContract
}
