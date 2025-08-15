package Infrastructure

import "golang.org/x/crypto/bcrypt"

func Hashpassword(password string) (string, error) {

	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashpassword), err
}

func ComparePassword(userPassword, password string)  error{
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	return err
}