package payload

import uuid "github.com/satori/go.uuid"

type UserResponse struct {
	ID          uuid.UUID
	Email       string `json:"email" form:"email"`
	Role        string `json:"role" form:"role"`
	Name        string `json:"name" form:"name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	Address     string `json:"address" form:"address"`
}

type UserRequest struct {
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	RetypePassword string `json:"retype_password" form:"retype_password"`
	Role           string `json:"role" form:"role"`
	Name           string `json:"name" form:"name"`
	PhoneNumber    string `json:"phone_number" form:"phone_number"`
	Address        string `json:"address" form:"address"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UserParam struct {
	Keyword string
	Email   string
}
