package v1

import (
	"context"

	"xcylla.io/common/log"
	"xcylla.io/config/internal/users"
	"xcylla.io/config/internal/users/dtos"
	"xcylla.io/config/pkg/constant"
	"xcylla.io/config/pkg/router"
)

type handlers struct {
	uu  users.Usecase
	log log.Logger
}

func NewUsersHandlers(uu users.Usecase) *handlers {
	return &handlers{uu, log.NewLogger("UsersHandlers")}
}

func (h *handlers) Login(c router.RouterContext) {
	h.log.Info("New Login handler started")
	var (
		ctx, cancel = context.WithTimeout(context.Background(), constant.DurationHandlerContext)
		req         dtos.LoginRequest
	)
	defer cancel()

	err := c.BindBody(&req)
	if err != nil {
		h.log.Error("Failed to bind request body")
		c.ReturnBadRequest(nil, err)
		return
	}

	res, err := h.uu.Login(ctx, req)
	if err != nil {
		h.log.Error("Error logging in user")
		c.ReturnInternalError(nil, err)
		return
	}

	h.log.Info("Login successful")
	c.ReturnOK(res)
}
