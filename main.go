package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	fmt.Println("")
	// TODO: Move to dependency injection container? We'll probably have to in case we need to make this configurable.
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	log.Info().Msg("Configured logger")

	addr := fmt.Sprintf("tcp://:%d", 30001)

	application := InitializeApp()

	panic(application.Run(addr))
}
