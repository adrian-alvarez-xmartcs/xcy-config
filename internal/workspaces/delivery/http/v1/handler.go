package v1

import (
	"context"

	"xcylla.io/common/log"
	"xcylla.io/config/internal/workspaces"
	"xcylla.io/config/pkg/constant"
	"xcylla.io/config/pkg/router"
)

type handlers struct {
	uc  workspaces.Usecase
	log log.Logger
}

func NewWorkspacesHandlers(uc workspaces.Usecase) *handlers {
	return &handlers{uc, log.NewLogger("WorkspacesHandler")}
}

func (h *handlers) GetWorkspaces(c router.RouterContext) {
	h.log.Info("New GetWorkspaces handler started")
	var (
		ctx, cancel = context.WithTimeout(context.Background(), constant.DurationHandlerContext)
	)
	defer cancel()

	res, err := h.uc.GetWorkspacesList(ctx)
	if err != nil {
		h.log.Error("Error getting workspaces")
		c.ReturnInternalError(nil, err)
		return
	}

	h.log.Info("Returning GetWorkspaces result")
	c.ReturnOK(res)
}
