package app

import (
	workspaceV1 "xcylla.io/config/internal/workspaces/delivery/http/v1"
	workspaceEntities "xcylla.io/config/internal/workspaces/entities"
	workspaceRepository "xcylla.io/config/internal/workspaces/repository"
	workspaceUsecase "xcylla.io/config/internal/workspaces/usecase"

	userV1 "xcylla.io/config/internal/users/delivery/http/v1"
	userRepository "xcylla.io/config/internal/users/repository"
	userUsecase "xcylla.io/config/internal/users/usecase"
)

func (app *App) startService() error {
	err := app.MigrationTables()
	for _, e := range err {
		if e != nil {
			app.log.Error("Error automigrating: ", e)
		}
	}

	workspaceRepo := workspaceRepository.NewWorkspacesRepository(app.db)
	workspaceUsecase := workspaceUsecase.NewWorkspacesUsecase(workspaceRepo)
	workspaceRoute := workspaceV1.NewWorkspacesHandlers(workspaceUsecase)
	workspaceRoute.WorkspacesRoutes(app.router) //Routes

	userRepo := userRepository.NewUsersRepository(app.db)
	userUsecase := userUsecase.NewUsersUsecase(userRepo, workspaceUsecase)
	userRoute := userV1.NewUsersHandlers(userUsecase)
	userRoute.UserRoutes(app.router) //Routes

	return nil
}

func (app *App) MigrationTables() []error {
	var err []error

	err = append(err, app.db.MainDb.AutoMigrate(&workspaceEntities.Def_Workspace{}))

	return err
}
