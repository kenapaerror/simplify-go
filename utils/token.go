package utils

import (
	"os"
	"simplify-go/exception"
	"simplify-go/helper"
	"simplify-go/model/web"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(request web.TokenCreateRequest, value time.Duration) string {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))

	expirationTime := time.Now().Add(time.Minute * value)

	claims := &web.TokenClaims{
		UserId:    request.UserId,
		Email:     request.Email,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Role:      request.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtTokenSecret)
	helper.PanicIfError(err)

	return tokenString
}

func ClaimsToken(refreshToken string) web.TokenClaims {
	var jwtTokenSecret = []byte(os.Getenv("JWT_TOKEN_SECRET"))

	claims := &web.TokenClaims{}

	token, err := jwt.ParseWithClaims(refreshToken, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtTokenSecret, nil
		},
	)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			panic(exception.NewUnauthorizedError(err.Error()))
		}
	}

	if !token.Valid {
		panic(exception.NewUnauthorizedError(err.Error()))
	}

	return *claims
}
