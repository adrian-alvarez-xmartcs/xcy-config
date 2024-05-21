package workspaces

import (
	"context"

	"xcylla.io/config/internal/workspaces/entities"
)

type Repository interface {
	Atomic(ctx context.Context, repo func(tx Repository) error) error

	GetWorkspaces(context.Context) ([]entities.Workspace, error)
}
