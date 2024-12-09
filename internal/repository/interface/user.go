package repository

import (
	"github.com/ViniciusDSLima/AuthSystem/internal/domain/entity"
)

type UserRepositoryInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	GetAll() ([]entity.User, error)
	UpdatePassword(id, password string) error
}
