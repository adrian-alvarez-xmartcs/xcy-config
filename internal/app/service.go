package app

import (
	v1 "xcylla.io/config/internal/workspaces/delivery/http/v1"
	"xcylla.io/config/internal/workspaces/entities"
	"xcylla.io/config/internal/workspaces/repository"
	"xcylla.io/config/internal/workspaces/usecase"
)

func (app *App) startService() error {
	err := app.MigrationTables()
	for _, e := range err {
		if e != nil {
			app.log.Error("Error automigrating: ", e)
		}
	}

	workspaceRepo := repository.NewWorkspacesRepository(app.db)
	workspaceUsecase := usecase.NewUseCase(workspaceRepo, app.cfg)
	workspaceRoute := v1.NewHandlers(workspaceUsecase)

	workspaceRoute.WorkspacesRoutes(app.router)

	return nil
}

func (app *App) MigrationTables() []error {
	var err []error

	err = append(err, app.db.MainDb.AutoMigrate(&entities.Def_Workspace{}))

	return err
}
