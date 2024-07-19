package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"hc/pkg/database"
	"os"
)

func main() {
	// inefficient, but pretty.
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()

	log.Info().Msg("Configured logger")

	// TODO: Move to dependency injection container? Not sure
	connection, err := database.NewMySQLConnection(database.ConnectionInfo{
		Host:     "localhost",
		User:     "root",
		Password: os.Getenv("HC_MYSQL_PASSWORD"),
		Port:     3306,
		DBName:   "hcdb",
	})

	if err != nil {
		log.Fatal().Msgf("unable to establish database connection, err: %s", err.Error())
	}

	defer connection.Close()

	log.Info().Msg("Database connection established with driver 'mysql'")

	addr := fmt.Sprintf("tcp://:%d", 30001)

	application := InitializeApp()

	panic(application.Bootstrap(addr))
}
