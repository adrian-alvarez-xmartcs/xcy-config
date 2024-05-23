package workspaces

import (
	"context"

	"xcylla.io/config/internal/workspaces/dtos"
)

type Usecase interface {
	GetWorkspacesList(ctx context.Context) (response dtos.WorkspaceArrayResponse, err error)
	GetWorkspace(ctx context.Context, workspaceName string) (response dtos.WorkspaceResponse, err error)
}
