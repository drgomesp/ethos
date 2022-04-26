package ethoscli

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/drgomesp/ethos/pkg/ethereum"
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

	stdout := bufio.NewWriter(new(bytes.Buffer))
	buf := new(bytes.Buffer)
	stderr := bufio.NewReadWriter(bufio.NewReader(buf), bufio.NewWriter(buf))

	info := &buildInfo{
		map[string]map[string]string{},
	}

	compilerOptions := ethereum.CompilerOptions{
		Language: "Solidity",
		Sources:  map[string]ethereum.CompilerSource{},
		Settings: ethereum.CompilerSettings{
			EVMVersion: "istanbul",
			Optimizer: ethereum.OptimizerOptions{
				Enabled: false,
			},
			OutputSelection: map[string]map[string][]string{
				"*": {
					"*": []string{
						"abi",
						"metadata",
						"evm.bytecode",
					},
				},
			},
		},
	}

	var err error
	err = filepath.Walk(
		// TODO: try deleting ./contracts/ and see this breaking
		Config.ContractsDir,
		func(srcFilePath string, contractFileInfo fs.FileInfo, srcErr error) error {
			if !contractFileInfo.IsDir() && filepath.Ext(contractFileInfo.Name()) == ".sol" {
				//if err = compileContract(compilerOptions, contractFileInfo.Name(), srcFilePath, stdout,
				//	stderr); err != nil {
				//	return err
				//}

				basePath := strings.TrimPrefix(strings.TrimPrefix(srcFilePath,
					Config.ContractsDir), string(filepath.Separator))
				name := strings.TrimSuffix(
					contractFileInfo.Name(),
					filepath.Ext(contractFileInfo.Name()),
				)
				path := filepath.Join(Config.BuildDir, basePath)

				compilerOptions.Sources[basePath] = ethereum.CompilerSource{
					URLs: []string{srcFilePath},
				}

				info.contracts[srcFilePath] = map[string]string{
					".abi": filepath.Join(path, fmt.Sprintf("%s.abi", name)),
					".bin": filepath.Join(path, fmt.Sprintf("%s.bin", name)),
				}

				log.Debug().
					Interface("contract", info.contracts[srcFilePath]).
					Msg("done")

				return nil
			}

			return nil
		})

	if err = compileContracts(compilerOptions, stdout, stderr); err != nil {
		return err
	}

	err = filepath.Walk(Config.ContractsDir, func(srcFilePath string, srcFileInfo fs.FileInfo, srcErr error) error {
		if srcFileInfo.IsDir() {
			return nil
		}

		return filepath.Walk(Config.BuildDir, func(genFilePath string, genFileInfo fs.FileInfo, genErr error) error {
			if srcFileInfo.IsDir() {
				return nil
			}

			contractFileName := srcFileInfo.Name()
			genFileExtension := filepath.Ext(genFileInfo.Name())

			path := filepath.Join(Config.BuildDir, fmt.Sprintf(
				"%s%s",
				strings.TrimSuffix(
					contractFileName,
					filepath.Ext(contractFileName),
				),
				genFileExtension,
			))

			genFiles := info.contracts[srcFilePath]
			basePath := filepath.Dir(genFiles[".abi"])

			ext := filepath.Ext(srcFileInfo.Name())
			if ext == ".bin" || ext == ".abi" {
				if err = os.Rename(
					filepath.Join(basePath, srcFileInfo.Name()),
					path,
				); err != nil {
					return err
				}
			}

			return nil
		})
	})

	err = filepath.Walk(Config.ContractsDir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(info.Name()) == ".sol" {
			return generateBindings(path, info.Name())
		}

		return nil
	})

	if err != nil {
		return err
	}

	displayInfo(info)

	return nil
}

func compileContracts(opts ethereum.CompilerOptions, stdout io.Writer, stderr io.ReadWriter) error {
	data, err := json.Marshal(opts)
	if err != nil {
		return err
	}

	cmd := exec.Command(
		Config.Compiler,
		"--standard-json",
	)
	cmd.Stdin = bytes.NewReader(data)

	// open the out file for writing
	outfile, err := os.Create("./compile.json")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	log.Trace().Str("cmd", cmd.String()).Msg("compiling")
	//cmd.Stderr = stderr

	err = cmd.Run()
	if err != nil {
		if d, e := ioutil.ReadAll(stderr); e == nil {
			return errors.Wrap(err, string(d))
		}
	}

	//if len(output) > 0 {
	//	log.Trace().Msg(string(output))
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
