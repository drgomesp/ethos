package ethoscli

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"io/fs"
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

	stdout := bufio.NewWriter(bytes.NewBuffer(nil))
	stderr := bufio.NewWriter(bytes.NewBuffer(nil))
	err = filepath.Walk(Config.ContractsDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if err != nil {
			return err
		}

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

func compileContract(path string, stdout, stderr io.Writer) error {
	cmd := exec.Command(
		"solc",
		"--abi",
		"--bin",
		path,
		"-o",
		Config.BuildDir,
		"--overwrite",
	)

	cmd.Stdout = stdout
	cmd.Stderr = stderr

	return cmd.Run();
}
