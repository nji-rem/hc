package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io/fs"
	"path/filepath"
	"strings"
)

type OptionFunc func(v *viper.Viper) error

func WithConfigDirectory(directory string) OptionFunc {
	return func(v *viper.Viper) error {
		return filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !strings.HasSuffix(info.Name(), ".yaml") && !strings.HasSuffix(info.Name(), ".yml") {
				return nil
			}

			v.SetConfigFile(path)
			if err := v.ReadInConfig(); err != nil {
				return err
			}

			log.Debug().Msgf("Loaded config file %s", info.Name())

			return nil
		})
	}
}

func WithEnvFile(envFile string) OptionFunc {
	return func(v *viper.Viper) error {
		if err := godotenv.Load(envFile); err != nil {
			return err
		}

		v.AutomaticEnv()

		return nil
	}
}

func Build(opts ...OptionFunc) (*viper.Viper, error) {
	v := viper.New()
	for _, opt := range opts {
		if err := opt(v); err != nil {
			return nil, err
		}
	}

	return v, nil
}
