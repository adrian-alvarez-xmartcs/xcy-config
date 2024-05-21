package dtos

type (
	WorkspaceArrayResponse []WorkspaceResponse

	WorkspaceResponse struct {
		Id   int
		Name string
	}
)
