package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	// inefficient, but pretty.
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	addr := fmt.Sprintf("tcp://:%d", 30001)

	application := InitializeApp()

	panic(application.Bootstrap(addr))
}
