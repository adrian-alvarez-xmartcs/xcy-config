package v1

import "xcylla.io/config/pkg/router"

func (h *handlers) WorkspacesRoutes(r *router.Router) {
	r.GET("/api/workspace/get", h.GetWorkspaces)
}
