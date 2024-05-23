package database

import (
	"gorm.io/gorm"
	"xcylla.io/config/internal/workspaces/entities"
	"xcylla.io/config/pkg/config"
	"xcylla.io/config/pkg/mock"
)

func LoadWorkspaceData(db *gorm.DB, cfg config.DatabaseConfig) {
	for k, v := range cfg.Db {
		tx := db.Create(&entities.Def_Workspace{Name: k, Description: v.Description})
		if tx.Error != nil {
			logging.Error("error %s", tx.Error.Error())
		}
	}
}

func LoadUserData(db *gorm.DB, cfg config.DatabaseConfig) {
	mock.FillUser(db)
}
