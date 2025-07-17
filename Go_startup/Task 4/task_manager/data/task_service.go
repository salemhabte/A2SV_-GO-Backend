package data

import (
	"errors"
	"task_manager/models"
)

var tasks = make(map[string] models.Tasks)

func GetTasks()[]models.Tasks{
	taskList := [] models.Tasks{}

	for _, task := range tasks{
		taskList = append(taskList, task)
	}

	return taskList
}

func GetTaskById(id string) (models.Tasks, error){

	// for _, task := range tasks{
	// 	if id == task.ID {
	// 		return task
	// 	}
	// }
	// return models.Tasks{}
	task,exist := tasks[id]

	if !exist {
		return models.Tasks{}, errors.New("task not found")
	}
	return task, nil
}

func UpdateTask(id string, update models.Tasks) error{

	_,exist := tasks[id]

	if ! exist{
		return errors.New("task not found") 
	}
	update.ID = id
	tasks[id] = update
	return nil
}

func RemoveTask( id string) error{
	_, exist := tasks[id]

	if !exist{
		return errors.New("not found")
	}
	delete(tasks, id)
	return nil
}

func NewTask(newTask models.Tasks)error{
	if _, exits := tasks[newTask.ID]; exits{
		return errors.New("task with this ID already exists")
	}
	tasks[newTask.ID] = newTask
	return nil
}