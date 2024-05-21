package config

import (
	"xcylla.io/common/config"
)

var Cfg Config

func Load() {
	err := config.LoadConfig("app", &Cfg)
	if err != nil {
		panic("unable to load config")
	}
}
