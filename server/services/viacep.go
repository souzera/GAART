package services

import (
	"encoding/json"
	"fmt"

	"net/http"
)

type ViaCEPResponse struct {
	CEP string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade string `json:"unidade"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF string `json:"uf"`
	Estado string `json:"estado"`
	Regiao string `json:"regiao"`
	IBGE string `json:"ibge"`
	GIA string `json:"gia"`
	DDD string `json:"ddd"`
	SIAFI string `json:"siafi"`
}

func ConsultarCEP(cep string) (ViaCEPResponse, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	response, err := http.Get(url)
	if err != nil {
		return ViaCEPResponse{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return ViaCEPResponse{}, fmt.Errorf("CEP não encontrado")
	}

	var result ViaCEPResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return ViaCEPResponse{}, err
	}

	return result, nil
}