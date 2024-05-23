package usecase

import (
	"context"

	"xcylla.io/common/log"
	"xcylla.io/config/internal/workspaces"
	"xcylla.io/config/internal/workspaces/dtos"
)

type usecase struct {
	repo workspaces.Repository
	log  log.Logger
}

func NewWorkspacesUsecase(repo workspaces.Repository) workspaces.Usecase {
	return &usecase{repo, log.NewLogger("WorkspacesUsecase")}
}

func (wu *usecase) GetWorkspacesList(ctx context.Context) (response dtos.WorkspaceArrayResponse, err error) {
	wu.log.Debug("Getting list of workspaces")

	workspaces, err := wu.repo.GetWorkspaces(ctx)
	if err != nil {
		wu.log.Error("Error getting workspaces")
		return nil, err
	}

	var dtoWorkspace []dtos.WorkspaceResponse
	for _, w := range workspaces {
		dtoWorkspace = append(dtoWorkspace, dtos.WorkspaceResponse{Id: int(w.ID), Name: w.Name, Description: w.Description})
	}

	wu.log.Debug("Successfully retrieved workspaces list")
	return dtoWorkspace, err
}

func (wu *usecase) GetWorkspace(ctx context.Context, workspaceName string) (response dtos.WorkspaceResponse, err error) {
	wu.log.Debug("Getting workspace")

	workspace, err := wu.repo.GetWorkspace(ctx, workspaceName)
	if err != nil {
		return dtos.WorkspaceResponse{}, err
	}

	dtoWorkspace := dtos.WorkspaceResponse{Id: int(workspace.ID), Name: workspace.Name, Description: workspace.Description}

	wu.log.Debug("Successfully retrieved workspace")
	return dtoWorkspace, nil
}
