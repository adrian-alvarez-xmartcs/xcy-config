package config

type (
	Config struct {
		Database DatabaseConfig
		Backend  BackendConfig
	}

	DatabaseConfig struct {
		Provider string
		Postgres struct {
			Ip       string
			Port     string
			Database string
			User     string
			Password string
		}
		Databases []string
	}
	BackendConfig struct {
		Port string
	}
)

func (c Config) Validate() error {
	return nil
}
