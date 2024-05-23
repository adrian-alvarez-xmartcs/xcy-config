package users

import (
	"context"

	"gorm.io/gorm"
	"xcylla.io/config/internal/users/entities"
)

type Repository interface {
	Atomic(ctx context.Context, repo func(tx Repository) error) error
	SetDatabase(database *gorm.DB)

	Login(ctx context.Context, user, password string) (*entities.Def_Users, error)
}
