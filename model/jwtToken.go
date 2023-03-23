package model

import (
	"github.com/golang-jwt/jwt/v5"
)

// JWTClient struct
type JWTClient struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

// Claims struct
type Claims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
