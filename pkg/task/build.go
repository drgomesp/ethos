package task

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/drgomesp/ethos/pkg/compiler"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type buildInfo struct {
	contracts map[string]map[string]string
}

func Clean(ctx context.Context) error {
	return os.RemoveAll(Config.BuildDir)
}

func Build(ctx context.Context) error {
	if err := Clean(ctx); err != nil {
		return err
	}

	if err := os.Mkdir(Config.BuildDir, fs.ModePerm); err != nil {
		return err
	}

	info := &buildInfo{map[string]map[string]string{}}

	files := make([]string, 0)
	filepath.Walk(Config.ContractsDir, func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(path) == ".sol" {
			files = append(files, path)
		}

		return nil
	})

	contracts, err := compiler.CompileSolidity(Config.Compiler, files...)
	if err != nil {
		return err
	}

	// If the entire compiler code was specified, build and bind based on that
	var (
		abis  []string
		bins  []string
		types []string
		sigs  []map[string]string
		libs  = make(map[string]string)
	)

	for name, contract := range contracts {
		abi, err := json.Marshal(contract.Info.AbiDefinition) // Flatten the compiler parse
		if err != nil {
			utils.Fatalf("Failed to parse ABIs from compiler output: %v", err)
		}
		abis = append(abis, string(abi))
		bins = append(bins, contract.Code)
		sigs = append(sigs, contract.Hashes)
		nameParts := strings.Split(name, ":")
		types = append(types, nameParts[len(nameParts)-1])

		libPattern := crypto.Keccak256Hash([]byte(name)).String()[2:36]
		libs[libPattern] = nameParts[len(nameParts)-1]

		// Generate the contract binding
		code, err := bind.Bind(types, abis, bins, sigs, "main", bind.LangGo, libs, nil)
		if err != nil {
			utils.Fatalf("Failed to generate ABI binding: %v", err)
		}

		if err := ioutil.WriteFile("test.go", []byte(code), 0600); err != nil {
			utils.Fatalf("Failed to write ABI binding: %v", err)
		}
	}

	displayInfo(info)

	return nil
}

func compileContracts(opts compiler.CompilerOptions, stdout io.Writer, stderr io.ReadWriter) error {
	//data, err := json.Marshal(opts)
	//if err != nil {
	//	return err
	//}

	//os.File.Readdir
	files, err := os.ReadDir(Config.ContractsDir)
	if err != nil {
		return err
	}

	for _, f := range files {
		log.Debug().Msg(f.Name())
	}

	panic("")
	//cmd := exec.Command(
	//	Config.Compiler,
	//	"--standard-json",
	//)

	//jsonOutput, err := ioutil.ReadFile(c.GlobalString(jsonFlag.Name))
	//if err != nil {
	//	utils.Fatalf("Failed to read combined-json from compiler: %v", err)
	//}
	//contracts, err = compiler.ParseCombinedJSON(jsonOutput, "", "", "", "")
	//if err != nil {
	//	utils.Fatalf("Failed to read contract information from json output: %v", err)
	//}

	//log.Debug().RawJSON("data", data).Msg("")
	//cmd.Stdin = strings.NewReader(string(data))
	//
	//output, err := cmd.CombinedOutput()
	//if err != nil {
	//	return err
	//}
	//
	//log.Trace().Str("cmd", cmd.String()).Msg("compiling")
	//
	//if len(output) > 0 {
	//	contracts, err := compiler.ParseCombinedJSON(output, "", "", "", "")
	//	if err != nil {
	//		return err
	//	}
	//
	//	for name, contract := range contracts {
	//		log.Trace().
	//			Str("name", name).
	//			Interface("contract", contract).
	//			Msg("compiled")
	//	}
	//}

	return nil
}

func generateBindings(basePath string, contractFileName string) error {
	name := strings.TrimSuffix(contractFileName, filepath.Ext(contractFileName))
	basePath = strings.TrimSuffix(basePath, contractFileName)
	out := filepath.Join(Config.ContractBindingsDir, strings.ToLower(fmt.Sprintf("%s.go", name)))

	if err := os.MkdirAll(Config.ContractBindingsDir, fs.ModePerm); err != nil {
		panic(err)
	}

	basePath = filepath.Join(Config.BuildDir, basePath)
	binPath := filepath.Join(basePath, fmt.Sprintf("%s.bin", name))
	abiPath := filepath.Join(basePath, fmt.Sprintf("%s.abi", name))

	if _, err := os.Stat(binPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Warn().Msg(err.Error())
			return nil
		} else {
			return err
		}
	}

	if _, err := os.Stat(abiPath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Warn().Msg(err.Error())
			return nil
		} else {
			return err
		}
	}

	cmd := exec.Command(
		"abigen",
		"--bin",
		binPath,
		"--abi",
		abiPath,
		fmt.Sprintf("--pkg=%s", Config.ContractsBindingsPkg),
		fmt.Sprintf("--type=%s", name),
		fmt.Sprintf("--out=%s", out),
	)

	log.Trace().Str("cmd", cmd.String()).Msg("generating bindings")

	output, err := cmd.Output()
	if err != nil {
		if len(output) > 0 {
			log.Trace().Msg(string(output))
		}

		return err

		//if d, e := ioutil.ReadAll(stderr); e == nil {
		//	return errors.Wrap(err, string(d))
		//}
	}

	return nil
}
