package usecase

import (
	"context"

	"xcylla.io/common/log"
	"xcylla.io/config/internal/users"
	"xcylla.io/config/internal/users/dtos"
	"xcylla.io/config/internal/workspaces"
	"xcylla.io/config/pkg/database"
	"xcylla.io/config/pkg/utils"
)

type usecase struct {
	repo             users.Repository
	log              log.Logger
	workspaceUsecase workspaces.Usecase
}

func NewUsersUsecase(repo users.Repository, workspaceUsecase workspaces.Usecase) users.Usecase {
	return &usecase{repo, log.NewLogger("UsersUsecase"), workspaceUsecase}
}

func (uu *usecase) Login(ctx context.Context, payload dtos.LoginRequest) (response dtos.LoginResponse, err error) {
	uu.log.Debug("Attemting to log in user")

	_, err = uu.workspaceUsecase.GetWorkspace(ctx, payload.Workspace)
	if err != nil {
		return dtos.LoginResponse{}, err
	}

	db, err := database.ConnectUserDatabase(payload.Workspace)
	if err != nil {
		return dtos.LoginResponse{}, err
	}
	uu.repo.SetDatabase(db)

	hashPw := utils.EncryptPassword(payload.Password)

	user, err := uu.repo.Login(ctx, payload.Username, hashPw)
	if err != nil {
		uu.log.Error("Error logging in user: %s", err)
		return dtos.LoginResponse{}, err
	}

	response = dtos.LoginResponse{
		ID:       user.ID,
		Name:     user.Name,
		Subname:  user.Subname,
		Username: user.Username,
		Email:    user.Email,
		Role: dtos.RoleResponse{
			Write:  user.Role.Write,
			Read:   user.Role.Read,
			Delete: user.Role.Delete,
		},
	}

	return response, nil
}
