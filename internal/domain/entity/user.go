package entity

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Address   Address            `json:"address"`
	CreatedAt *time.Time         `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time         `bson:"updated_at" json:"updated_at"`
}

func (u *User) UserValidate() error {

	if u.Name == "" || u.Email == "" || u.Password == "" {
		return errors.New("invalid user data")
	}

	if u.Address.ZipCode == "" || u.Address.Number == "" {
		return errors.New("invalid user data")
	}

	return nil
}

func (u *User) CreateAddress(addr Address) {
	u.Address = addr
}

func (u *User) SetCreatedAt() {
	createdAt := time.Now()
	u.CreatedAt = &createdAt
}

func (u *User) UpdatePassword(password string) {
	u.Password = password
	u.SetUpdatedAt()
}

func (u *User) SetUpdatedAt() {
	updatedAt := time.Now()
	u.UpdatedAt = &updatedAt
}

func (u *User) EncryptPassword(hashedPassword string) {
	u.Password = hashedPassword
}
