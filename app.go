package main

import (
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog/log"
	"hc/api/config"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

type App struct {
	Config     config.Reader
	GameServer gnet.EventHandler
}

func (a *App) Run(addr string) error {
	a.startProfilerIfEnabled()

	log.Info().Msgf("Starting %s on address %s", a.Config.GetString("app.name"), addr)

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

// startProfilerIfEnabled starts the pprof profiler if enabled. It's encouraged to keep the profiler active on production
// environments in order to locate potential bottlenecks.
func (a *App) startProfilerIfEnabled() {
	profilerEnabled := a.Config.Get("app.profiler_enabled")
	if profilerEnabled == nil || !profilerEnabled.(bool) {
		log.Warn().Msg("Profiler is disabled! It's recommended to enable it on production environments.")
		return
	}

	go func() {
		if err := http.ListenAndServe("localhost:8080", nil); err != nil {
			log.Error().Msgf("Profiler stopped! Err: %s", err.Error())
		}
	}()
}
