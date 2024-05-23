package dtos

type (
	LoginRequest struct {
		Workspace string `json:"workspace"`
		Username  string `json:"username"`
		Password  string `json:"password"`
	}

	LoginResponse struct {
		ID       string       `json:"id"`
		Name     string       `json:"name"`
		Subname  string       `json:"subname"`
		Username string       `json:"username"`
		Email    string       `json:"email"`
		Role     RoleResponse `json:"role"`
	}
)

func (lr LoginRequest) Validate() error {
	return nil
}
