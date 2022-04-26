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

	info := &buildInfo{
		map[string]map[string]string{},
	}

	var err error
	err = filepath.Walk(
		// TODO: try deleting ./contracts/ and see this breaking
		Config.ContractsDir,
		func(srcFilePath string, contractFileInfo fs.FileInfo, srcErr error) error {
			if !contractFileInfo.IsDir() && filepath.Ext(contractFileInfo.Name()) == ".sol" {
				if err = compileContract(srcFilePath, stdout, stderr); err != nil {
					return err
				}

				basePath := filepath.Dir(srcFilePath)

				name := strings.TrimSuffix(
					contractFileInfo.Name(),
					filepath.Ext(contractFileInfo.Name()),
				)

				info.contracts[contractFileInfo.Name()] = map[string]string{
					".abi": filepath.Join(Config.BuildDir, basePath, fmt.Sprintf("%s.abi", name)),
					".bin": filepath.Join(Config.BuildDir, basePath, fmt.Sprintf("%s.bin", name)),
				}

				log.Debug().
					Interface("contract", info.contracts[contractFileInfo.Name()]).
					Msg("done")

				return nil
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

			genFiles := info.contracts[srcFileInfo.Name()]
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
	targetPath := filepath.Join(
		Config.BuildDir,
		filepath.Dir(path),
	)

	cmd := exec.Command(
		Config.Compiler,
		"--abi",
		"--bin",
		path,
		"--output-dir",
		targetPath,
		"--overwrite",
		"--ignore-missing",
	)

	log.Trace().Str("cmd", cmd.String()).Msg("compiling")
	cmd.Stderr = stderr

	output, err := cmd.Output()
	if err != nil {
		if d, e := ioutil.ReadAll(stderr); e == nil {
			return errors.Wrap(err, string(d))
		}
	}

	if len(output) > 0 {
		log.Trace().Msg(string(output))
	}

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
