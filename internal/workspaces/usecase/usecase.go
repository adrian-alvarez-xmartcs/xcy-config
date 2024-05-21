package usecase

import (
	"context"

	"xcylla.io/common/log"
	"xcylla.io/config/internal/workspaces"
	"xcylla.io/config/internal/workspaces/dtos"
	"xcylla.io/config/pkg/config"
)

type usecase struct {
	repo workspaces.Repository
	cfg  config.Config
	log  log.Logger
}

func NewUseCase(repo workspaces.Repository, cfg config.Config) workspaces.Usecase {
	return &usecase{repo, cfg, log.NewLogger("WorkspacesUseCase")}
}

func (uc *usecase) GetWorkspacesList(ctx context.Context) (response dtos.WorkspaceArrayResponse, err error) {
	uc.log.Debug("Getting list of workspaces")

	workspaces, err := uc.repo.GetWorkspaces(ctx)
	if err != nil {
		uc.log.Error("Error getting workspaces")
		return nil, err
	}

	var dtoWorkspace []dtos.WorkspaceResponse
	for _, w := range workspaces {
		dtoWorkspace = append(dtoWorkspace, dtos.WorkspaceResponse{Id: int(w.ID), Name: w.Name})
	}

	uc.log.Debug("Successfully retrieved workspaces list")
	return dtoWorkspace, err
}
