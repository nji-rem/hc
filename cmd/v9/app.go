package main

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog/log"
	"sync"
)

type App struct {
	// GameServer contains an instance to the game server socket runner.
	GameServer gnet.EventHandler
}

func (a *App) Bootstrap(addr string) error {
	log.Info().Msgf("Starting game server on address %s", addr)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := gnet.Run(a.GameServer, addr, gnet.WithMulticore(true)); err != nil {
			log.Error().Msgf("unable to execute game server: %s", err.Error())
		}
	}()

	wg.Wait()

	return nil
}
