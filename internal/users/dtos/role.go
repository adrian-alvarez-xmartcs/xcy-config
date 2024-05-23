package dtos

type (
	RoleResponse struct {
		Write  bool `json:"write"`
		Read   bool `jsno:"read"`
		Delete bool `json:"delete"`
	}
)
