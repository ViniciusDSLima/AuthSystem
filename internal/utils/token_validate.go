package utils

import (
	"errors"
	repository "github.com/ViniciusDSLima/AuthSystem/internal/repository/implementation"
	"time"
)

type ValidateRecoveryTokenUseCase struct {
	repo repository.RecoveryTokenRepository
}

func NewValidateRecoveryTokenUseCase(repo repository.RecoveryTokenRepository) *ValidateRecoveryTokenUseCase {
	return &ValidateRecoveryTokenUseCase{repo: repo}
}

func (uc *ValidateRecoveryTokenUseCase) Execute(token string) (string, error) {
	response, err := uc.repo.FindByToken(token)

	if err != nil || time.Now().After(response.ExpiresAt) {
		return "", errors.New("token inválido ou expirado")
	}

	return response.UserId.Hex(), nil
}
