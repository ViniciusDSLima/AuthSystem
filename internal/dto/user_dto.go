package dto

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RecoveryPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type UpdatePasswordRequest struct {
	Token    string `json:"token" validate:"required,token"`
	Password string `json:"password" validate:"required,min=6"`
}
