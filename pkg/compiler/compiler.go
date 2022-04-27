package compiler

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/rs/zerolog/log"
)

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

type SourceFile struct {
	Content string `json:"content"`
}

type OutputSelectionMap map[SourceName]map[ContractName][]string

type Settings struct {
	Optimizer struct {
		Enabled bool `json:"enabled"`
	} `json:"optimizer"`
	OutputSelection OutputSelectionMap     `json:"outputSelection"`
	EVMVersion      string                 `json:"evmVersion"`
	Libraries       map[LibraryName]string `json:"libraries,omitempty"`
}

type Input struct {
	Language Lang                       `json:"language"`
	Sources  map[SourceName]*SourceFile `json:"sources"`
	Settings Settings                   `json:"settings"`
}

func (i *Input) JSON() []byte {
	data, err := json.Marshal(i)
	if err != nil {
		log.Fatal().Err(err)
	}

	return data
}

type OutputSource struct {
	ID  int
	AST interface{}
}

type ImmutableReference struct {
	Start  uint `json:"start"`
	Length uint `json:"length"`
}

type OutputByteCode struct {
	Object              string
	Opcodes             string
	SourceMap           string
	LinkReferences      map[SourceName]map[LibraryName][]string
	ImmutableReferences map[string][]ImmutableReference
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
