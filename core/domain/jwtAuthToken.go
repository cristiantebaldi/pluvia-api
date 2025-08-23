package domain

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtAuthToken struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
	Type    string `json:"type"`
}

type JwtToken struct {
	Administrador Administrador `json:"account,omitempty"`
	Exp           int64         `json:"exp,omitempty"`
	jwt.Claims
}

func NewJwtToken(administrador Administrador, exp int64) JwtToken {
	jwtToken := JwtToken{
		Administrador: administrador,
		Exp:           exp,
	}

	return jwtToken
}

func NewJwtAuthToken(token string, hash string) JwtAuthToken {
	jwtAuthToken := JwtAuthToken{
		Token:   token,
		Refresh: hash,
		Type:    "refreshToken",
	}

	return jwtAuthToken
}
