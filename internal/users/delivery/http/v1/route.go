package v1

import "xcylla.io/config/pkg/router"

func (h *handlers) UserRoutes(r *router.Router) {
	r.POST("/api/user/login", h.Login)
}
