package app

import (
	"os"
	"os/signal"
	"syscall"

	"xcylla.io/common/log"
	"xcylla.io/config/pkg/config"
	"xcylla.io/config/pkg/database"
	"xcylla.io/config/pkg/router"
)

type App struct {
	db     *database.Database
	router *router.Router
	cfg    config.Config
	log    log.Logger
}

func NewApp(cfg config.Config) *App {
	return &App{
		db:     database.Initialize(cfg.Database),
		router: router.NewRouter(),
		cfg:    cfg,
		log:    log.NewLogger("App"),
	}
}

func (app *App) Run() error {
	if err := app.startService(); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-quit
		app.log.Warning("Server is shutting down...")
		app.db.Close(app.cfg.Database.Folder, true)

		os.Exit(0)
	}()

	return app.router.Start(app.cfg.Backend.Port)
}
