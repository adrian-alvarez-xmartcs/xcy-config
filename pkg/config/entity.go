package config

type (
	Config struct {
		Database DatabaseConfig
		Backend  BackendConfig
	}

	DatabaseConfig struct {
		Folder string
		Mock   bool
		Db     map[string]dbConfig
	}

	dbConfig struct {
		Provider    string
		Description string
		Data        Provider //Interface
	}

	ProviderSqlite struct {
		Provider
		DatabaseName string
	}

	ProviderPostgre struct {
		Provider
		Database string
		Ip       string
		Port     string
		User     string
		Password string
	}

	BackendConfig struct {
		Port string
	}
)

type (
	Provider interface {
	}
)

func (c Config) Validate() error {
	return nil
}
