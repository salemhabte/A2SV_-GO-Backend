package mocks

import (
	 "go_clean/Domain"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *Domain.User) error{
	args := m.Called(user)

	return  args.Error(0)
}

func (m *MockUserRepository)GetByUsername(username string) (*Domain.User, error) {
	args := m.Called(username)

	return  args.Get(0).(*Domain.User),args.Error(1)
}