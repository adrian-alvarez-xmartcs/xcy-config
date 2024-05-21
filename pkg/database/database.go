package database

import (
	"fmt"
	"os"
	"path/filepath"

	"xcylla.io/common/log"
	"xcylla.io/config/pkg/config"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	_sqlite   = "sqlite"
	_postgres = "postgres"
)

var (
	db               = &gorm.DB{}
	logging          = log.NewLogger("Database")
	databaseMainName = "database.db"
)

func Initialize(cfg config.DatabaseConfig) *gorm.DB {
	logging.Trace("Initializing database")

	execPath, err := os.Executable()
	if err != nil {
		logging.Error("Error getting executable path:", err)
		return nil
	}
	databaseNameFolder := filepath.Join(filepath.Dir(execPath), "..", "database")

	if _, err := os.Stat(databaseNameFolder); os.IsNotExist(err) {
		if os.Mkdir(databaseNameFolder, os.ModePerm) != nil {
			logging.Error("Error creating database folder")
			return nil
		}
	} else {
		os.Remove(filepath.Join(databaseNameFolder, databaseMainName))
	}

	switch cfg.Provider {
	case _sqlite:
		logging.Trace("Database provider SQLite")
		db = connect_SQLITE(databaseNameFolder)
	case _postgres:
		logging.Trace("Database provider POSTGRE")
		postgres := cfg.Postgres
		db = connect_POSTGRE(postgres.Ip, postgres.User, postgres.Password, postgres.Database, postgres.Port)
	}

	return db
}

func connect_SQLITE(databaseNameFolder string) *gorm.DB {
	queryStr := filepath.Join(databaseNameFolder, databaseMainName)
	db, err := gorm.Open(sqlite.Open(queryStr), &gorm.Config{Logger: logger.Discard})
	if err != nil {
		logging.Error("Error connecting to the SQLite database:", err)
		return nil
	}

	return db
}

func connect_POSTGRE(host, user, password, dbname, port string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Discard})
	if err != nil {
		logging.Error("Error connecting to the PostgreSQL database:", err)
		return nil
	}

	return db
}

func Clear() {
	if db != nil {
		dbConn, _ := db.DB()
		if err := dbConn.Close(); err != nil {
			logging.Error("Error closing database connection:", err)
			return
		}
	}

	execPath, err := os.Executable()
	if err != nil {
		logging.Error("Error getting executable path:", err)
		return
	}
	databaseNameFolder := filepath.Join(filepath.Dir(execPath), "..", "database")

	if _, err := os.Stat(databaseNameFolder); err == nil {
		if err := os.RemoveAll(databaseNameFolder); err != nil {
			logging.Error("Error deleting database folder:", err)
			return
		}
	}
}
