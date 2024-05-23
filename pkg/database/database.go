package database

import (
	"fmt"
	"os"
	"path/filepath"

	"xcylla.io/common/log"
	"xcylla.io/config/pkg/config"
	"xcylla.io/config/pkg/database/migration"

	"gorm.io/gorm"
)

// const (
// 	_sqlite   = "sqlite"
// 	_postgres = "postgres"
// )

type Database struct {
	MainDb   *gorm.DB
	SystemDb *gorm.DB
	Database *gorm.DB
}

var (
	logging      = log.NewLogger("Database")
	mainDbName   = "main"
	systemDbName = "system"
)

func Initialize(cfg config.DatabaseConfig) *Database {
	logging.Trace("Initializing databases")
	var db Database

	db.MainDb = connect_SQLITE(cfg.Folder, mainDbName+".db")
	if db.MainDb == nil {
		logging.Error("Failed to connect to main database")
		panic("Failed to connect to main database")
	}

	db.SystemDb = connect_SQLITE(cfg.Folder, systemDbName+".db")
	if db.SystemDb == nil {
		logging.Error("Failed to connect to system database")
		panic("Failed to connect to system database")
	}

	migration.MigrateMainTables(db.MainDb)
	LoadWorkspaceData(db.MainDb, cfg)

	return &db
}

func ConnectDatabase(databaseName string, cfg config.DatabaseConfig) *gorm.DB {
	db := connect_SQLITE(cfg.Folder, databaseName+".db")
	if db == nil {
		logging.Error("Failed to connect to database")
		panic("Failed to connect to database")
	}

	migration.MigrateUserTables(db)
	LoadUserData(db, cfg)

	return db
}

func (db *Database) Close(folder string, delete bool) {
	if main, err := db.MainDb.DB(); err != nil {
		return
	} else {
		main.Close()
	}

	if system, err := db.SystemDb.DB(); err != nil {
		return
	} else {
		system.Close()
	}

	if db.Database != nil {
		if database, err := db.Database.DB(); err != nil {
			return
		} else {
			database.Close()
		}
	}

	if delete {
		dbFiles := []string{mainDbName, systemDbName}
		for _, dbFile := range dbFiles {
			dbPath := filepath.Join(folder, dbFile+".db")
			if err := os.Remove(dbPath); err != nil {
				fmt.Printf("Error removing %s: %v\n", dbPath, err)
			}
		}
	}
}
