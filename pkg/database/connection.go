package database

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func connect_SQLITE(folder, databaseName string) *gorm.DB {
	cxd, err := os.Executable()
	if err != nil {
		return nil
	}
	os.Chdir(cxd + "/../..")

	queryStr := filepath.Join(folder, databaseName)
	fmt.Println(queryStr)
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
