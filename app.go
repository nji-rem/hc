package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/panjf2000/gnet/v2"
	"github.com/rs/zerolog/log"
	"hc/api/config"
	"net/http"
	_ "net/http/pprof"
)

type App struct {
	Config     config.Reader
	GameServer gnet.EventHandler
	DB         *sqlx.DB
}

func (a *App) Run(addr string) error {
	a.checkDatabase()
	a.startProfilerIfEnabled()

	log.Info().Msgf("Starting %s on address %s", a.Config.GetString("app.name"), addr)

	if err := gnet.Run(a.GameServer, addr, gnet.WithMulticore(true)); err != nil {
		log.Error().Msgf("unable to execute game server: %s", err.Error())
	}

	return nil
}

// startProfilerIfEnabled starts the pprof profiler if enabled. It's encouraged to keep the profiler active on production
// environments in order to locate potential bottlenecks.
func (a *App) startProfilerIfEnabled() {
	profilerEnabled := a.Config.GetBool("app.profiler_enabled")
	if !profilerEnabled {
		log.Warn().Msg("Profiler is disabled! Enabling it is recommended. Enable it in config/app.yaml or set APP_PROFILER_ENABLED=true")
		return
	}

	log.Info().Msg("Profiler (pprof) enabled")

	go func() {
		if err := http.ListenAndServe("localhost:8080", nil); err != nil {
			log.Error().Msgf("Profiler stopped! Err: %s", err.Error())
		}
	}()
}

func (a *App) checkDatabase() {
	if err := a.DB.Ping(); err != nil {
		log.Fatal().Msgf("Server stopped! Database connection is not available, err: %s", err.Error())
	}

	log.Info().Msg("Database connection is established")
}
