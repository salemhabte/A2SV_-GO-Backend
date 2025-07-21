package data

import (
	"context"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollection * mongo.Collection
func ConnectDB(){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10* time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil{
		panic(err)
	}

	TaskCollection = client.Database("taskdb").Collection("tasks")
}

func GetTasks()([]models.Task,error){
	cursor, err := TaskCollection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}

	var tasks []models.Task

	err = cursor.All(context.TODO(), &tasks)

	return tasks, err
	
}

func GetTaskById(id string) (models.Task, error){

	objID, err := primitive.ObjectIDFromHex(id)

	if err !=nil {
		return models.Task{}, err
	}
	var task models.Task

	err = TaskCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	return task, err
}

func UpdateTask(id string, UpdateTask models.Task) error{

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}
	update := bson.M{
		"$set": UpdateTask,
	}
	_, err = TaskCollection.UpdateByID(context.TODO(), objID,update)

	return err
}

func RemoveTask( id string) error{
	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil{
		return err
	}
	_, err = TaskCollection.DeleteOne(context.TODO(),bson.M{
		"_id": objId,
	})
	return err
}

func NewTask(newTask models.Task)(*mongo.InsertOneResult,error){
	return TaskCollection.InsertOne(context.TODO(), newTask)
	
}

