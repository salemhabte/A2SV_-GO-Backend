package controllers

import (
	"fmt"
	"go_clean/Domain"
	usecases "go_clean/Usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskDTO struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	DueDate     string             `bson:"due_date" json:"due_date"` // Or use time.Time
	Status      string             `bson:"status" json:"status"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
}
type UserController struct {
	Taskusecase *usecases.TaskUsecase
	Userusecase *usecases.UserUseCase
}

type UserDTO struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"` // Omit in response
	Role     string             `json:"role" bson:"role"`
}

func NewController(task usecases.TaskUsecase, user usecases.UserUseCase) *UserController {
	return &UserController{
		Taskusecase: &task,
		Userusecase: &user,
	}
}
func (ts *UserController) GetTask(ctx *gin.Context) {
	
	tasks, err := ts.Taskusecase.GetTasks()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (ts *UserController) GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := ts.Taskusecase.GetByID(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, task)
}
func (ts *UserController) UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask TaskDTO
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ts.Taskusecase.UpdateTask(id, *ts.ChangeToDomainTask(&updatedTask)); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task updated Successfully",
	})
}

func (ts *UserController) RemoveTask(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := ts.Taskusecase.DeleteTask(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task deleted Successfuly",
	})
}

func (ts *UserController) CreatedTask(ctx *gin.Context) {
	var newTask TaskDTO

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := ts.Taskusecase.CreateTask(*ts.ChangeToDomainTask(&newTask))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "successuly created",
	})
}
func (uc *UserController) RegisterUser(ctx *gin.Context) {
	var user UserDTO

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid request payload",
		})
		return
	}
	fmt.Println(user)
	err := uc.Userusecase.Register(uc.ChangeToDomain(&user))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "successuly created",
	})
}

func (uc *UserController) Login(ctx *gin.Context) {

	var user UserDTO

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}
	jwtToken, err := uc.Userusecase.Login(user.Username, user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})
}

func (uc *UserController) ChangeToDomain(userDTO *UserDTO) *Domain.User {
	return &Domain.User{
		ID:       userDTO.ID.Hex(),
		Username: userDTO.Username,
		Password: userDTO.Password,
		Role:     userDTO.Role,
	}
}

func (tc *UserController) ChangeToDomainTask(taskDTO *TaskDTO) *Domain.Task {
	return &Domain.Task{
		ID:          taskDTO.ID.Hex(),
		Title:       taskDTO.Title,
		Description: taskDTO.Description,
		DueDate:     taskDTO.DueDate,
		Status:      taskDTO.Status,
		UserID:      taskDTO.UserID.Hex(),
	}
}
