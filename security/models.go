package security

import (
	"github.com/golang-jwt/jwt/v5"
)

// TokenType represents the type of JWT token.
type TokenType int

const (
	AccessToken TokenType = iota
	RefreshToken
)

type Claims struct {
	Email string `json:"username"`
	jwt.RegisteredClaims
}
