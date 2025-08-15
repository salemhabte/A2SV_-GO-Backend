package usecases

import (
	"errors"
	"go_clean/Domain"
	"go_clean/Infrastructure"
	"go_clean/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseTestSuite struct {
    suite.Suite
    mockRepo *mocks.MockTaskRepository
    usecase  *UserUseCase
}

func (suite *UserUsecaseTestSuite) TestRegister_sucsess(){
	user := &Domain.User{
        Username: "testuser",
        Password: "plainpass",
    }
	suite.mockRepo.On("CreateUser", user).Return(nil)

    // Act
    err := suite.usecase.Register(user)

    // Assert
    assert.Nil(suite.T(), err)
}

func (suite *UserUsecaseTestSuite) TestLogin_Success() {
    password := "secret"
    hashed, _ := Infrastructure.Hashpassword(password)
    user := &Domain.User{ID: "123", Username: "testuser", Password: hashed, Role: "admin"}

    suite.mockRepo.On("GetByUsername", "testuser").Return(user, nil)

    token, err := suite.usecase.Login("testuser", password)

    suite.Nil(err)
    suite.NotEmpty(token)
    suite.mockRepo.AssertExpectations(suite.T())
}


func (suite *UserUsecaseTestSuite) TestLogin_UserNotFound() {
    suite.mockRepo.On("GetByUsername", "unknown").Return(nil, errors.New("not found"))

    token, err := suite.usecase.Login("unknown", "pass")

    suite.EqualError(err, "user not found")
    suite.Equal("", token)
    suite.mockRepo.AssertExpectations(suite.T())
}


func (suite *UserUsecaseTestSuite) TestLogin_InvalidPassword() {
    password := "secret"
    hashed, _ := Infrastructure.Hashpassword(password)
    user := &Domain.User{ID: "123", Username: "testuser", Password: hashed, Role: "admin"}

    suite.mockRepo.On("GetByUsername", "testuser").Return(user, nil).Once()

    token, err := suite.usecase.Login("testuser", "wrongpass")

    suite.EqualError(err, "invalid credentials")
    suite.Equal("", token)
    suite.mockRepo.AssertExpectations(suite.T())
}