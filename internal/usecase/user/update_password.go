package usecase

import (
	"errors"
	"github.com/ViniciusDSLima/AuthSystem/internal/dto"
	repository "github.com/ViniciusDSLima/AuthSystem/internal/repository/implementation"
	"github.com/ViniciusDSLima/AuthSystem/internal/utils"
)

type UpdatePasswordUseCase struct {
	userRepo       repository.UserRepository
	tokenValidator utils.ValidateRecoveryTokenUseCase
}

func NewUpdatePasswordUseCase(userRepo repository.UserRepository, tokenValidator utils.ValidateRecoveryTokenUseCase) *UpdatePasswordUseCase {
	return &UpdatePasswordUseCase{
		userRepo:       userRepo,
		tokenValidator: tokenValidator,
	}
}

func (uc *UpdatePasswordUseCase) Execute(request dto.UpdatePasswordRequest) error {

	userId, err := uc.tokenValidator.Execute(request.Token)

	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(request.Password)

	if err != nil {
		return errors.New("erro ao hashear a senha")
	}

	return uc.userRepo.UpdatePassword(userId, hashedPassword)
}
