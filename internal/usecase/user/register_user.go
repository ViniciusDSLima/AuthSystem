package usecase

import (
	"errors"
	"github.com/ViniciusDSLima/AuthSystem/internal/domain/entity"
	repository "github.com/ViniciusDSLima/AuthSystem/internal/repository/implementation"
	"github.com/ViniciusDSLima/AuthSystem/internal/utils"
	"github.com/ViniciusDSLima/AuthSystem/pkg/cep"
	"log"
)

type RegisterUserUseCase struct {
	repo repository.UserRepository
}

func NewRegisterUserUseCase(
	repo repository.UserRepository,
) *RegisterUserUseCase {
	return &RegisterUserUseCase{
		repo: repo,
	}
}

func (uc *RegisterUserUseCase) Register(user *entity.User) error {
	if err := user.UserValidate(); err != nil {
		return err
	}

	u, err := uc.userAlreadyExists(user.Email)

	if err != nil {
		log.Printf("%v", err)
	}

	if u != nil {
		return errors.New("user already exists")
	}

	address, err := cep.GetAddress(user.Address.ZipCode)

	if err != nil {
		return errors.New("falha ao buscar endereco")
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	user.CreateAddress(address)
	user.EncryptPassword(hashedPassword)
	user.SetCreatedAt()

	return uc.repo.Create(user)
}

func (uc *RegisterUserUseCase) userAlreadyExists(email string) (*entity.User, error) {

	user, err := uc.repo.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return user, nil
	}

	return user, nil
}
