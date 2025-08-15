package mocks

import (
	 "go_clean/Domain"

	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) GetAll() ([]Domain.Task, error) {
	args := m.Called()

	return args.Get(0).([]Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) GetByID (id string) (*Domain.Task, error){
	args := m.Called(id)

	return args.Get(0).(*Domain.Task), args.Error(1)
}

func (m *MockTaskRepository) Update(id string, UpdateTask Domain.Task) error {
	args := m.Called(id,UpdateTask)

	return args.Error(0)
}

func (m * MockTaskRepository) Delete(id string) error {
	args := m.Called(id)

	return args.Error(0)
}

func (m *MockTaskRepository) CreateTask(newTask *Domain.Task) error {
	args := m.Called(newTask)

	return args.Error(0)
}