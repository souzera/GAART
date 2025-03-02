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

func errorItemNotFound(name string) error {
	return fmt.Errorf("item %s não foi encontrado ou não existe", name)
}

func errorCustomMessage(messagem string) error {
	return fmt.Errorf("ERROR: %s", messagem)
}
