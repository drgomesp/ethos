package ethoscli

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/alexeyco/simpletable"
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

	info := new(buildInfo)
	info.contracts = map[string]map[string]string{}
	var err error

	// TODO: try deleting ./contracts/ and see this breaking
	err = filepath.Walk(
		Config.ContractsDir,
		func(srcFilePath string, contractFileInfo fs.FileInfo, srcErr error) error {
			if !contractFileInfo.IsDir() {
				if e := compileContract(srcFilePath, stdout, stderr); e != nil {
					err = e
					return e
				}

				name := strings.TrimSuffix(
					contractFileInfo.Name(),
					filepath.Ext(contractFileInfo.Name()),
				)

				info.contracts[contractFileInfo.Name()] = map[string]string{
					".abi": filepath.Join(Config.BuildDir, fmt.Sprintf("%s.abi", name)),
					".bin": filepath.Join(Config.BuildDir, fmt.Sprintf("%s.bin", name)),
				}

				return generateBindings(stdout, stderr, contractFileInfo.Name())
			}

			return nil
		})

	if err != nil {
		return err
	}

	err = filepath.Walk(Config.ContractsDir, func(srcFilePath string, srcFileInfo fs.FileInfo, srcErr error) error {
		if srcFileInfo.IsDir() {
			return nil
		}

		return filepath.Walk(
			Config.BuildDir,
			func(genFilePath string, genFileInfo fs.FileInfo, genErr error) error {
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

				if ext := filepath.Ext(srcFileInfo.Name()); ext != ".sol" {
					return os.Rename(
						filepath.Join(Config.BuildDir, srcFileInfo.Name()),
						path,
					)
				}

				return nil
			})

	})

	if err != nil {
		return err
	}

	displayInfo(info)

	return nil
}

func displayInfo(info *buildInfo) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "src"},
			{Align: simpletable.AlignCenter, Text: "abi"},
			{Align: simpletable.AlignCenter, Text: "bin"},
		},
	}

	for name, paths := range info.contracts {
		r := []*simpletable.Cell{
			{Text: name},
		}

		for _, path := range paths {
			r = append(
				r,
				&simpletable.Cell{
					Text: path,
				},
			)
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())

	log.Info().Msg("contracts compiled successfully")
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
			log.Trace().Msg(string(output))
		}

		if d, e := ioutil.ReadAll(stderr); e == nil {
			return errors.Wrap(err, string(d))
		}
	}

	return nil
}

func generateBindings(stdout io.Writer, stderr io.ReadWriter, contractFileName string) error {
	name := strings.TrimSuffix(contractFileName, filepath.Ext(contractFileName))
	out := filepath.Join(Config.ContractBindingsDir, strings.ToLower(fmt.Sprintf("%s.go", name)))

	cmd := exec.Command(
		"abigen",
		"--bin",
		filepath.Join(Config.BuildDir, fmt.Sprintf("%s.bin", name)),
		"--abi",
		filepath.Join(Config.BuildDir, fmt.Sprintf("%s.abi", name)),
		fmt.Sprintf("--pkg=%s", Config.ContractsBindingsPkg),
		fmt.Sprintf("--type=%s", name),
		fmt.Sprintf("--out=%s", out),
	)

	log.Trace().Str("cmd", cmd.String()).Msg("compiling ")

	output, err := cmd.Output()
	if err != nil {
		if len(output) > 0 {
			log.Trace().Msg(string(output))
		}

		if d, e := ioutil.ReadAll(stderr); e == nil {
			return errors.Wrap(err, string(d))
		}
	}

	return nil
}
