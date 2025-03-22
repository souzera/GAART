package scripts

import "github.com/souzera/GAART/schemas"

func validatePassword(password string) bool {
	return len(password) < 8
}

func validateConfirmPassword(password string, confirmPassword string) bool {
	return password == confirmPassword
}

func validateLogin(login string) bool {
	if len(login) < 3 {
		return false
	}

	if db.Where("login = ?", login).First(&schemas.Usuario{}).RowsAffected > 0 {
		return false
	}

	return true
}