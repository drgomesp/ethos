package task

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/drgomesp/ethos/pkg/build"
	"github.com/drgomesp/ethos/pkg/compiler"
	"github.com/rs/zerolog/log"
	"io/fs"
	"os"
	"os/exec"
)

type buildInfo struct {
	contracts map[string]map[string]string
}

func Clean(ctx context.Context) error {
	return os.RemoveAll(Config.BuildDir)
}

func Build(ctx context.Context) error {
	var err error
	if err = Clean(ctx); err != nil {
		return err
	}

	if err = os.Mkdir(Config.BuildDir, fs.ModePerm); err != nil {
		return err
	}

	b, err := build.NewFromSourceDir(Config.SourcesDir)
	if err != nil {
		return err
	}

	compilerIn, err := b.CompilerInput()
	if err != nil {
		return err
	}

	//log.Debug().
	//	Interface("compilerIn", string(data)).
	//	Msg("build prepared")

	cmd := exec.Command(
		Config.Compiler,
		"--standard-json",
	)

	if err = os.WriteFile("compilerIn.json", compilerIn.JSON(), fs.ModePerm); err != nil {
		return err
	}

	cmd.Stdin = bytes.NewReader(compilerIn.JSON())

	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	var compilerOut compiler.Output
	if err = json.Unmarshal(out, &compilerOut); err != nil {
		return err
	}

	log.Debug().
		Interface("file", compilerOut.Contracts["contracts/access/AccessControl.sol"]).
		Msg("")

	return nil
}
