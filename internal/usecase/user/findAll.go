package usecase

import (
	"errors"
	"github.com/ViniciusDSLima/AuthSystem/internal/domain/entity"
	repository "github.com/ViniciusDSLima/AuthSystem/internal/repository/implementation"
)

type ListUsersUseCase struct {
	repo repository.UserRepository
}

func NewListUsersUseCase(
	repo repository.UserRepository,
) *ListUsersUseCase {
	return &ListUsersUseCase{
		repo: repo,
	}
}

func (uc *ListUsersUseCase) ListUsers() ([]entity.User, error) {
	users, err := uc.repo.GetAll()

	if err != nil {
		return nil, errors.New("Erro ao buscar usuários")
	}

	return users, nil
}
