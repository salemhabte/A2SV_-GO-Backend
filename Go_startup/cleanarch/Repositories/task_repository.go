package repositories

import (
	"context"
	"go_clean/Domain"
	"go_clean/Usecases"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepo struct {
}

func NewTaskRepo() usecases.TaskRepositoryInterface {
	return &TaskRepo{}
}
func (t *TaskRepo) GetAll() ([]Domain.Task, error) {
	cursor, err := TaskCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	var tasks []Domain.Task

	err = cursor.All(context.TODO(), &tasks)

	return tasks, err
}

func (t *TaskRepo) GetByID(id string) (*Domain.Task, error) {

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return &Domain.Task{}, err
	}
	var task Domain.Task

	err = TaskCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	return &task, err
}

func (t *TaskRepo) Update(id string, UpdateTask Domain.Task) error {

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	update := bson.M{
		"$set": UpdateTask,
	}
	_, err = TaskCollection.UpdateByID(context.TODO(), objID, update)

	return err
}

func (t *TaskRepo) Delete(id string) error {
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	_, err = TaskCollection.DeleteOne(context.TODO(), bson.M{
		"_id": objId,
	})
	return err
}

func (t *TaskRepo) CreateTask(newTask *Domain.Task) error {
	_, err := TaskCollection.InsertOne(context.TODO(), newTask)
	return err

}
