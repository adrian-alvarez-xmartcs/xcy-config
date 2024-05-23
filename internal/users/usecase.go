package users

import (
	"context"

	"xcylla.io/config/internal/users/dtos"
)

type Usecase interface {
	Login(ctx context.Context, payload dtos.LoginRequest) (response dtos.LoginResponse, err error)
}
