package ethoscli

import (
	"context"
	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io/fs"
	"os"
	"path/filepath"
)

func Build(ctx context.Context) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	err = filepath.Walk(Config.ContractsDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		contracts, err := compiler.CompileSolidity("solc", filepath.Join(dir, path))
		if err != nil {
			return errors.Wrap(err, "failed to compile Solidity source files")
		}

		log.Info().Interface("contract", contracts[path]).Msg("contract built")

		return nil
	})

	return err
}
