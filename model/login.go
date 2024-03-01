package model

import "github.com/golang-jwt/jwt/v5"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Claim data that the token payload will have
type Claim struct {
	Email string `json:"email"`
	jwt.MapClaims
}
