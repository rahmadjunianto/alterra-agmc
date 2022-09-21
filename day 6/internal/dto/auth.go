package dto

import "github.com/golang-jwt/jwt/v4"

type (
	CreateUsersRequestBody struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginUsersRequestBody struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	JWTClaims struct {
		ID    uint   `json:"user_id"`
		Email string `json:"email"`
		Name  string `json:"name"`
		jwt.RegisteredClaims
	}
)
