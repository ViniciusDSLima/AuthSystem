package usecase

import (
	"errors"
	repository "github.com/ViniciusDSLima/AuthSystem/internal/repository/implementation"
	"github.com/ViniciusDSLima/AuthSystem/internal/utils"
)

type LoginUseCase struct {
	repo repository.UserRepository
}

func NewLoginUseCase(
	repo repository.UserRepository,
) *LoginUseCase {
	return &LoginUseCase{
		repo: repo,
	}
}

func (uc *LoginUseCase) Login(email, password string) (string, error) {
	user, err := uc.repo.FindByEmail(email)

	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}

	if !utils.VerifyPassword(password, user.Password) {
		return "", errors.New("credenciais inválidas")
	}

	return utils.GenerateJWT(user.Id.Hex())
}
