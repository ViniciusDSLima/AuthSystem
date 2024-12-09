package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RecoveryToken struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	UserId    primitive.ObjectID `bson:"userId"`
	Token     string             `bson:"token"`
	ExpiresAt time.Time          `bson:"expiresAt"`
}
