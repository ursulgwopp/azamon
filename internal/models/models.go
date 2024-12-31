package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type SignUpRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenClaims struct {
	jwt.StandardClaims
	Id int
}

type Profile struct {
	Id        int         `json:"id"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Balance   float64     `json:"balance"`
	ItemsList []uuid.UUID `json:"itemsList"`
}
