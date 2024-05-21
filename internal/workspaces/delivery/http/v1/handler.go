package v1

import (
	"context"
	"time"

	"xcylla.io/common/log"
	"xcylla.io/config/internal/workspaces"
	"xcylla.io/config/pkg/router"
)

type handlers struct {
	uc  workspaces.Usecase
	log log.Logger
}

func NewHandlers(uc workspaces.Usecase) *handlers {
	return &handlers{uc, log.NewLogger("WorkspacesHandler")}
}

func (h *handlers) GetWorkspaces(c router.RouterContext) {
	h.log.Debug("New GetWorkspaces handler")
	var (
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	)
	defer cancel()

	res, err := h.uc.GetWorkspacesList(ctx)
	if err != nil {
		h.log.Error("Error getting workspaces")
		c.ReturnInternalError(nil, err)
		return
	}

	h.log.Debug("Returning GetWorkspaces result")
	c.ReturnOK(res)
}
