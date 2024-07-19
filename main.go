package main

import (
	"embed"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

//go:embed internal/account/infrastructure/migrations
var embedMigrations embed.FS

func main() {
	// inefficient, but pretty.
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	log.Info().Msg("Configured logger")

	addr := fmt.Sprintf("tcp://:%d", 30001)

	application := InitializeApp()

	panic(application.Bootstrap(addr))
}
