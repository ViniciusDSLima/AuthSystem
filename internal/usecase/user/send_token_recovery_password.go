package usecase

import (
	"errors"
	repository "github.com/ViniciusDSLima/AuthSystem/internal/repository/implementation"
	"github.com/ViniciusDSLima/AuthSystem/internal/utils"
	"github.com/ViniciusDSLima/AuthSystem/pkg/email"
	"time"
)

type SendTokenRecoveryPasswordUseCase struct {
	userRepository  repository.UserRepository
	tokenRepository repository.RecoveryTokenRepository
}

func NewSendTokenRecoveryPasswordUseCase(
	userRepository repository.UserRepository,
	tokenRepository repository.RecoveryTokenRepository,
) *SendTokenRecoveryPasswordUseCase {
	return &SendTokenRecoveryPasswordUseCase{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
	}
}

func (uc *SendTokenRecoveryPasswordUseCase) SendTokenToEmail(email string) error {
	user, err := uc.userRepository.FindByEmail(email)

	if err != nil {
		return errors.New("usuário não encontrado")
	}

	token := utils.GenerateRandomToken(32)

	expirationDateTimeToken := time.Now().Add(1 * time.Hour)

	err = uc.tokenRepository.Create(user.Id.Hex(), token, expirationDateTimeToken)

	if err != nil {
		return errors.New("erro ao salvar token de recuperação")
	}

	err = services.SendEmail(email, token)

	if err != nil {
		return errors.New("falha ao enviar o email")
	}

	return nil
}
