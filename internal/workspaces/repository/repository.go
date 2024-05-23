package repository

import (
	"context"

	"gorm.io/gorm"
	"xcylla.io/common/log"
	"xcylla.io/config/internal/workspaces"
	"xcylla.io/config/internal/workspaces/entities"
	"xcylla.io/config/pkg/database"
)

type repository struct {
	db  *gorm.DB
	log log.Logger
}

func NewWorkspacesRepository(db *database.Database) workspaces.Repository {
	return &repository{
		db:  db.MainDb,
		log: log.NewLogger("WorkspacesRepository"),
	}
}

func (r *repository) Atomic(ctx context.Context, repo func(tx workspaces.Repository) error) error {
	txConn := r.db.Begin()
	if txConn.Error != nil {
		return txConn.Error
	}

	newRepository := &repository{db: txConn}

	err := repo(newRepository)
	if err != nil {
		return err
	}

	if newRepository.db.Error != nil {
		return newRepository.db.Error
	}

	return nil
}

func (r *repository) GetWorkspaces(ctx context.Context) ([]entities.Def_Workspace, error) {
	r.log.Trace("Getting Workspaces from repository")
	var workspaces []entities.Def_Workspace

	tx := r.db.Model(&entities.Def_Workspace{}).Find(&workspaces)
	if tx.Error != nil {
		r.log.Error("Unable to get workspaces")
		return nil, tx.Error
	}

	r.log.Trace("Returning Workspaces from repository")
	return workspaces, nil
}

func (r *repository) GetWorkspace(ctx context.Context, workspaceName string) (entities.Def_Workspace, error) {
	r.log.Trace("Getting Workspace from repository")
	var workspace entities.Def_Workspace

	tx := r.db.Model(&entities.Def_Workspace{}).Where("name = ?", workspaceName).First(&workspace)
	if tx.Error != nil {
		r.log.Error("Unable to get workspace")
		return entities.Def_Workspace{}, tx.Error
	}

	r.log.Trace("Returning Workspace from repository")
	return workspace, nil
}
