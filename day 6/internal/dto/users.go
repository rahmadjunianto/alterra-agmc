package dto

import (
	"time"

	"gorm.io/gorm"
)

type (
	UpdateUsersRequestBody struct {
		ID       *uint   `param:"id" validate:"required"`
		Name     *string `json:"name" validate:"omitempty"`
		Email    *string `json:"email" validate:"omitempty"`
		Password *string `json:"password" validate:"omitempty"`
	}
	LoginUsersResponse struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Token string `json:"token"`
	}
	UsersWithJWTResponse struct {
		UsersResponse
		JWT string `json:"jwt"`
	}
	UsersResponse struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	UsersWithCUDResponse struct {
		UsersResponse
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		DeletedAt *gorm.DeletedAt `json:"deleted_at"`
	}
)
