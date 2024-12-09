package repository

import (
	"context"
	"github.com/ViniciusDSLima/AuthSystem/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type RecoveryTokenRepository struct {
	Collection *mongo.Collection
}

func NewRecoveryTokenRepository(collection *mongo.Collection) *RecoveryTokenRepository {
	return &RecoveryTokenRepository{
		Collection: collection,
	}
}

func (repo *RecoveryTokenRepository) Create(userID string, token string, expiresAt time.Time) error {
	doc := bson.M{
		"userId":    userID,
		"token":     token,
		"expiresAt": expiresAt,
	}

	_, err := repo.Collection.InsertOne(context.TODO(), doc)

	return err
}

func (repo *RecoveryTokenRepository) FindByToken(token string) (*entity.RecoveryToken, error) {
	filter := bson.M{"token": token}

	var recoveryToken entity.RecoveryToken

	err := repo.Collection.FindOne(context.TODO(), filter).Decode(&recoveryToken)

	if err != nil {
		return nil, err
	}

	return &recoveryToken, nil
}
