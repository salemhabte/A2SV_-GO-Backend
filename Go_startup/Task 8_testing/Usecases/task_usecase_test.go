package usecases

import (
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "go_clean/Domain"
    "go_clean/mocks"
)

type TaskUsecaseTestSuite struct {
    suite.Suite
    mockRepo *mocks.MockTaskRepository
    usecase  *TaskUsecase
}

func (suite *TaskUsecaseTestSuite) SetupTest() {
    suite.mockRepo = new(mocks.MockTaskRepository)
    suite.usecase = NewTaskUsecase(suite.mockRepo)
}

func (suite *TaskUsecaseTestSuite)TestGetTasks_Sucess(){
	task := Domain.Task{ID: "1", Title: "Alice", Description: "hello tittle", Status: "active"}
	tasksList := []Domain.Task{task}
    suite.mockRepo.On("GetAll").Return(tasksList , nil)

    tasks,err := suite.usecase.GetTasks()
	assert.Equal(suite.T(),task , tasks)
    assert.Nil(suite.T(), err)
    suite.mockRepo.AssertExpectations(suite.T())

}

func (suite *TaskUsecaseTestSuite) TestCreateTask_Sucess() {
    task := Domain.Task{ID: "1", Title: "Alice", Description: "hello tittle", Status: "active"}
    suite.mockRepo.On("CreateTask",1).Return(nil)

    err := suite.usecase.CreateTask(task)
    assert.Nil(suite.T(), err)
    suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseTestSuite) TestGetByID_Sucess() {
    id := "1"
    expectedTask := &Domain.Task{ID: id, Title: "Alice", Description: "hello tittle", Status: "active"}
    suite.mockRepo.On("GetByID",id).Return(expectedTask,nil)

    task,err := suite.usecase.GetByID(id)
    assert.Equal(suite.T(),expectedTask,task)
    assert.Nil(suite.T(), err)
    suite.mockRepo.AssertExpectations(suite.T())
}

func (suite * TaskUsecaseTestSuite) TestUpdate_success(){
    id := "1"
    expectedTask := &Domain.Task{ID: id, Title: "Math", Description: "hello tittle", Status: "active"}
    suite.mockRepo.On("Update",id,expectedTask).Return(nil)

    err := suite.usecase.UpdateTask(id,*expectedTask)
    assert.Nil(suite.T(), err)
    suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseTestSuite) TestDelete_success(){
    id := "1"
    // expectedTask := &Domain.Task{ID: id, Title: "Alice", Description: "hello tittle", Status: "active"}
    suite.mockRepo.On("Delete",id).Return(nil)

    err := suite.usecase.DeleteTask(id)
    // assert.Equal(suite.T(),expectedTask,task)
    assert.Nil(suite.T(), err)
    suite.mockRepo.AssertExpectations(suite.T())
}

