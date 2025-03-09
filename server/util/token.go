package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GerarToken(id string) string {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": id,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return ""
	}

	return tokenString
}
