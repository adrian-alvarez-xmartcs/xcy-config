package database

import (
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"xcylla.io/common/log"
	"xcylla.io/config/pkg/config"
	"xcylla.io/config/pkg/database/migration"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	logging        = log.NewLogger("Database")
	mainDbName     = "main"
	systemDbName   = "system"
	database       = &Database{}
	UserDatabase   = &gorm.DB{}
	databaseFolder = ""
)

func Initialize(cfg config.DatabaseConfig) *Database {
	logging.Trace("Initializing databases")
	databaseFolder = cfg.Folder

	database.MainDb = connect_SQLITE(databaseFolder, mainDbName+".db")
	if database.MainDb == nil {
		logging.Error("Failed to connect to main database")
		panic("Failed to connect to main database")
	}

	database.SystemDb = connect_SQLITE(databaseFolder, systemDbName+".db")
	if database.SystemDb == nil {
		logging.Error("Failed to connect to system database")
		panic("Failed to connect to system database")
	}

	migration.MigrateMainTables(database.MainDb)
	LoadWorkspaceData(database.MainDb, cfg)

	return database
}

func connect_SQLITE(folder, databaseName string) *gorm.DB {
	cxd, err := os.Executable()
	if err != nil {
		return nil
	}
	os.Chdir(cxd + "/../..")

	queryStr := filepath.Join(folder, databaseName)
	db, err := gorm.Open(sqlite.Open(queryStr), &gorm.Config{Logger: logger.Discard})
	if err != nil {
		logging.Error("Error connecting to the SQLite database:", err)
		return nil
	}

	return db
}

// func connect_POSTGRE(host, user, password, dbname, port string) *gorm.DB {
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, user, password, dbname, port)
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Discard})
// 	if err != nil {
// 		logging.Error("Error connecting to the PostgreSQL database:", err)
// 		return nil
// 	}

// 	return db
// }

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
				logging.Error("Error removing %s: %v\n", dbPath, err)
				return
			}
		}
	}
}
