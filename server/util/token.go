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

func removerPrefixoToken(token string) string {
	if len(token) < 7 {
		return token
	}
	if token[:7] == "Bearer " {
		return token[7:]
	}
	return token
}

func ValidarToken(tokenString string) (string, error) {

	tokenString = removerPrefixoToken(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", err
	}

	return claims["sub"].(string), nil
}
