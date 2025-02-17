package handler

import "fmt"

func errorParamRequired(param string) error {
	return fmt.Errorf("parametro %s é obrigatório", param)
}

func errorUniqueViolation(param string) error {
	return fmt.Errorf("parametro %s já existe", param)
}

func errorNotFound(param string) error {
	return fmt.Errorf("parametro %s não encontrado", param)
}

func errorInvalid(param string) error {
	return fmt.Errorf("parametro %s inválido", param)
}

func errorCustomMessage(message string) error {
	return fmt.Errorf(message)
}
