package web

type RegisterRequest struct {
	Username string `validate:"required,min=3,max=100" json:"username"`
	Password string `validate:"required,min=8,max=100" json:"password"`
	Email    string `validate:"required,email" json:"email"`
}

type LoginRequest struct {
	Username string `validate:"required,min=3,max=100" json:"username"`
	Password string `validate:"required,min=8,max=100" json:"password"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
