package ethoscli

import (
	"bufio"
	"bytes"
	"context"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func Clean(ctx context.Context) error {
	return os.RemoveAll(Config.BuildDir)
}

func Build(ctx context.Context) error {
	_, err := os.Getwd()
	if err != nil {
		return err
	}

	if err := Clean(ctx); err != nil {
		return err
	}

	if err := os.Mkdir(Config.BuildDir, fs.ModePerm); err != nil {
		return err
	}

	stdout := bufio.NewWriter(new(bytes.Buffer))
	buf := new(bytes.Buffer)
	stderr := bufio.NewReadWriter(bufio.NewReader(buf), bufio.NewWriter(buf))

	err = filepath.Walk(Config.ContractsDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if err != nil {
			return err
		}

		var cfg = Config
		_ = cfg
		err = compileContract(path, stdout, stderr)
		if err != nil {
			return err
		}

		return nil
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
		"-v",
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
