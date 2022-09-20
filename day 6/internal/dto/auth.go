package dto

import "github.com/golang-jwt/jwt/v4"

type (
	RegisterEmployeeRequestBody struct {
		Name     string `json:"name" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	ByEmailAndPasswordRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	JWTClaims struct {
		UserID uint   `json:"user_id"`
		Email  string `json:"email"`
		Name   string `json:"name"`
		jwt.RegisteredClaims
	}
)
