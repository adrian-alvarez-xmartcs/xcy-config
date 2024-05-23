package migration

import (
	"gorm.io/gorm"
	usersEntities "xcylla.io/config/internal/users/entities"
	workspacesEntities "xcylla.io/config/internal/workspaces/entities"
)

func MigrateMainTables(db *gorm.DB) []error {
	var err []error

	err = append(err, db.AutoMigrate(&workspacesEntities.Def_Workspace{}))

	return err
}

func MigrateUserTables(db *gorm.DB) []error {
	var err []error

	err = append(err, db.AutoMigrate(&usersEntities.Def_Roles{}))
	err = append(err, db.AutoMigrate(&usersEntities.Def_Users{}))

	return err
}
