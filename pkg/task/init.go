package task

import (
	"context"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"io/fs"
	"os"
	"path/filepath"
)

func Init(ctx context.Context) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(DefaultEthosConfig())
	if err != nil {
		return err
	}

	filePath := filepath.Join(dir, EthosFile)
	err = os.WriteFile(filePath, data, fs.ModePerm)
	if err != nil {
		return err
	}

	log.Info().Msgf("Ethos config file initialized (%s)", filePath)

	return err
}
