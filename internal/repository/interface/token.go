package repository

import (
	"github.com/ViniciusDSLima/AuthSystem/internal/domain/entity"
	"time"
)

type TokenRepository interface {
	Create(userID string, token string, expiresAt time.Time) error
	FindByToken(token string) (*entity.RecoveryToken, error)
}
