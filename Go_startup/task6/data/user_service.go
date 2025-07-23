package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)



func InitUserCollection(db *mongo.Database) {
	userCollection = db.Collection("users")
}
func Register(user * models.User) (*models.User,error){

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	
	var existing models.User
	err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existing)
	if err == nil {
		return &models.User{},errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
	return &models.User{},err
	}

	user.Password = string(hashedPassword)
	_, err = userCollection.InsertOne(ctx, user)
	return user,err


}
var JwtSecret = []byte("your_jwt_secret")
func AuthenticateUser(username, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
  "user_id": user.ID,
  "username":   user.Username,
  "role": user.Role,
	})
	jwtToken, err := token.SignedString(JwtSecret)
	if err != nil {
	
	return "",err
	}
	return jwtToken, nil
}