package repository

import (
	"context"
	"github.com/ViniciusDSLima/AuthSystem/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type UserRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		userCollection: collection,
	}
}

func (r *UserRepository) Create(user *entity.User) error {

	_, err := r.userCollection.InsertOne(context.Background(), user)

	if err != nil {
		log.Println("Error inserting user: ", err)
	}

	return err
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err := r.userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)

	return &user, err
}

func (r *UserRepository) GetAll() ([]entity.User, error) {
	var users []entity.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	cursor, err := r.userCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user entity.User

		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) UpdatePassword(id, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"password": password}}

	_, err := r.userCollection.UpdateOne(ctx, filter, update)
	return err
}
