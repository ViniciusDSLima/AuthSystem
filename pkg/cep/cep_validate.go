package cep

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ViniciusDSLima/AuthSystem/internal/domain/entity"
)

func GetAddress(cep string) (entity.Address, error) {
	apiUrl := os.Getenv("API_VIACEP_URL")

	url := fmt.Sprintf("%s/%s/json", apiUrl, cep)

	resp, err := http.Get(url)

	if err != nil {
		return entity.Address{}, fmt.Errorf("erro ao acessar a API do ViaCEP: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entity.Address{}, fmt.Errorf("CEP inválido ou não encontrado: status %d", resp.StatusCode)
	}

	var viaCEPResponse struct {
		Cep         string `json:"cep"`
		Logradouro  string `json:"logradouro"`
		Complemento string `json:"complemento"`
		Bairro      string `json:"bairro"`
		Localidade  string `json:"localidade"`
		UF          string `json:"uf"`
		Erro        bool   `json:"erro,omitempty"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&viaCEPResponse); err != nil {
		return entity.Address{}, fmt.Errorf("erro ao decodificar resposta do ViaCEP: %w", err)
	}

	if viaCEPResponse.Erro {
		return entity.Address{}, errors.New("CEP não encontrado")
	}

	address := entity.Address{
		ZipCode:      viaCEPResponse.Cep,
		Street:       viaCEPResponse.Logradouro,
		Complement:   viaCEPResponse.Complemento,
		Neighborhood: viaCEPResponse.Bairro,
		City:         viaCEPResponse.Localidade,
		State:        viaCEPResponse.UF,
	}

	return address, nil
}
