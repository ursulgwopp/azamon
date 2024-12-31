package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// auth
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
	Username string
}

type Profile struct {
	Id        int         `json:"id"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	Balance   float64     `json:"balance"`
	ItemsList []uuid.UUID `json:"itemsList"`
}

// items
type Item struct {
	Id          uuid.UUID `json:"id"`
	Seller      string    `json:"seller"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
}
