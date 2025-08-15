package usecases
import "go_clean/Domain"

type TaskRepositoryInterface interface {
	CreateTask(task *Domain.Task) error
	GetByID(id string) (*Domain.Task, error)
	GetAll() ([]Domain.Task, error)
	Update(id string, UpdateTask Domain.Task) error
	Delete(id string) error
}

type UserRepositoryInterface interface {
	CreateUser(user *Domain.User) error
	GetByUsername(username string) (*Domain.User, error)
}