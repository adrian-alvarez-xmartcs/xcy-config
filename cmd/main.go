package main

import (
	"xcylla.io/common/log"
	"xcylla.io/config/internal/app"
	"xcylla.io/config/pkg/config"
)

func main() {
	log.Init()
	config.Load()

	app := app.NewApp(config.Cfg)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
