package database

import (
	"errors"

	"gorm.io/gorm"
	"xcylla.io/config/pkg/database/migration"
)

func ConnectUserDatabase(databaseName string) (*gorm.DB, error) {
	udb := connect_SQLITE(databaseFolder, databaseName+".db")
	if udb == nil {
		logging.Error("Failed to connect to database")
		return nil, errors.New("Failed to connect to database")
	}

	migration.MigrateUserTables(udb)
	LoadUserData(udb)

	logging.Debug("Successfully connected to %s database", databaseName)
	return udb, nil
}
