package usecases

import (
	"go_clean/Domain"
)

type TaskUsecase struct {
	TaskRepo TaskRepositoryInterface
}

func NewTaskUsecase(ts TaskRepositoryInterface) *TaskUsecase {
	return &TaskUsecase{
		TaskRepo: ts,
	}
}
func (ts *TaskUsecase) GetTasks() ([]Domain.Task, error) {
	return ts.TaskRepo.GetAll()
}

func (ts *TaskUsecase) CreateTask(task Domain.Task) error {
	return ts.TaskRepo.CreateTask(&task)
}

func (ts *TaskUsecase) GetByID(id string) (Domain.Task, error) {
	taskPtr, err := ts.TaskRepo.GetByID(id)
	if err != nil {
		return Domain.Task{}, err
	}
	return *taskPtr, nil
}

func (ts *TaskUsecase) UpdateTask(id string, UpdateTask Domain.Task) error {
	return ts.TaskRepo.Update(id, UpdateTask)
}

func (ts *TaskUsecase) DeleteTask(id string) error {
	return ts.TaskRepo.Delete(id)
}
