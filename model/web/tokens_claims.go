package web

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	UserId    string `json:"user_id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
	jwt.StandardClaims
}
