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

	err := filepath.Walk(
		Config.ContractsDir,
		func(srcFilePath string, srcFileInfo fs.FileInfo, srcErr error) error {
			if !srcFileInfo.IsDir() {
				return compileContract(srcFilePath, stdout, stderr)
			}

			return nil
		})

	err = filepath.Walk(
		Config.BuildDir,
		func(genFilePath string, genFileInfo fs.FileInfo, genErr error) error {
			if genFileInfo.IsDir() {
				return nil
			}

			return filepath.Walk(Config.ContractsDir, func(srcFilePath string, srcFileInfo fs.FileInfo, srcErr error) error {
				if srcFileInfo.IsDir() {
					return nil
				}

				srcFileName := srcFileInfo.Name()
				genFileExtension := filepath.Ext(genFileInfo.Name())

				newFileName := fmt.Sprintf(
					"%s%s",
					strings.TrimSuffix(
						srcFileName,
						filepath.Ext(srcFileName),
					),
					genFileExtension,
				)

				log.Debug().
					Str("genFilePath", genFilePath).
					Str("srcFilePath", srcFilePath).
					Str("srcFileName", srcFileName).
					Str("genFileExtension", genFileExtension).
					Str("newFileName", newFileName).
					Msg("")

				log.Info().Msg("transformed file names")

				return os.Rename(
					filepath.Join(Config.BuildDir, genFileInfo.Name()),
					filepath.Join(Config.BuildDir, newFileName),
				)
			})
		})

	if err != nil {
		return err
	}

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
