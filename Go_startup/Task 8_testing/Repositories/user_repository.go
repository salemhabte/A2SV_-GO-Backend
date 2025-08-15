package repositories

import (
	"context"
	"go_clean/Domain"
	"go_clean/Usecases"

	"go.mongodb.org/mongo-driver/bson"
)

type UserRepo struct {
}

func NewUserRepo() usecases.UserRepositoryInterface {
	return &UserRepo{}
}
func (u *UserRepo) CreateUser(user *Domain.User) error {
	_, err := userCollection.InsertOne(context.TODO(), user)
	return err
}

func (u *UserRepo) GetByUsername(username string) (*Domain.User, error) {
	var user Domain.User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	return &user, err
}
