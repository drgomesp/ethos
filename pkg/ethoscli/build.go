package ethoscli

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
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

	// TODO: try deleting ./contracts/ and see this breaking
	err := filepath.Walk(
		Config.ContractsDir,
		func(srcFilePath string, contractFileInfo fs.FileInfo, srcErr error) error {
			if !contractFileInfo.IsDir() {
				if err := compileContract(srcFilePath, stdout, stderr); err != nil {
					return err
				}

				return generateBindings(stdout, stderr, contractFileInfo.Name())
			}

			return nil
		})

	err = filepath.Walk(
		Config.BuildDir,
		func(genFilePath string, genFileInfo fs.FileInfo, genErr error) error {
			if genFileInfo.IsDir() {
				return nil
			}

			if ext := filepath.Ext(genFileInfo.Name()); ext != ".sol" {
				return nil
			}

			return filepath.Walk(Config.ContractsDir, func(srcFilePath string, srcFileInfo fs.FileInfo, srcErr error) error {
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

				log.Debug().Msg(path)

				return os.Rename(
					filepath.Join(Config.BuildDir, genFileInfo.Name()),
					path,
				)
			})
		})

	if err != nil {
		return err
	}

	return nil
}

func generateBindings(stdout io.Writer, stderr io.ReadWriter, contractFileName string) error {
	name := strings.TrimSuffix(contractFileName, filepath.Ext(contractFileName))
	out := filepath.Join(Config.ContractBindingsDir, fmt.Sprintf("%s.go", name))

	cmd := exec.Command(
		"abigen",
		"--bin",
		filepath.Join(Config.BuildDir, fmt.Sprintf("%s.bin", name)),
		"--abi",
		filepath.Join(Config.BuildDir, fmt.Sprintf("%s.abi", name)),
		fmt.Sprintf("--pkg=%s", Config.ContractsBindingsPkg),
		fmt.Sprintf("--out=%s", out),
	)

	log.Trace().Msg(cmd.String())

	output, err := cmd.Output()
	if err != nil {
		if len(output) > 0 {
			log.Debug().Msg(string(output))
		}

		if d, e := ioutil.ReadAll(stderr); e == nil {
			return errors.Wrap(err, string(d))
		}
	}

	log.Info().Msg("contracts compiled successfully")

	return nil
}

func compileContract(path string, stdout io.Writer, stderr io.ReadWriter) error {
	cmd := exec.Command(
		Config.Compiler,
		"--abi",
		"--bin",
		path,
		"--output-dir",
		Config.BuildDir,
	)

	cmd.Stderr = stderr

	output, err := cmd.Output()
	if err != nil {
		if len(output) > 0 {
			log.Debug().Msg(string(output))
		}

		if d, e := ioutil.ReadAll(stderr); e == nil {
			return errors.Wrap(err, string(d))
		}
	}

	log.Info().Msg("contracts compiled successfully")

	return nil
}
