package util

import "golang.org/x/crypto/bcrypt"


func VerificarSenha(senha, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(senha))
}

func HashSenha(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	return string(bytes), err
}