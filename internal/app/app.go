package app

import (
	"gorm.io/gorm"
	"xcylla.io/common/log"
	"xcylla.io/config/pkg/config"
	"xcylla.io/config/pkg/database"
	"xcylla.io/config/pkg/router"
)

type App struct {
	db     *gorm.DB
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

	return app.router.Start(app.cfg.Backend.Port)
}
