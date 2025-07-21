package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)
func GetTask(ctx *gin.Context) {
	
	tasks,err := data.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func GetTaskById ( ctx *gin.Context){
	id := ctx.Param("id")
	task, err := data.GetTaskById(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound,gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, task)
}
func UpdateTask(ctx *gin.Context){
	id := ctx.Param("id")
	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return 
	}

	if err := data.UpdateTask(id, updatedTask); err != nil{
		ctx.JSON(http.StatusOK, gin.H{
			"error" : err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "Task updated Successfully",
	})
}

func RemoveTask(ctx *gin.Context){
	id := ctx.Param("id")

	if err := data.RemoveTask(id); err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{
			"error" : err.Error(),
			
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "Task deleted Successfuly",
	})
}

func CreatedTask(ctx *gin.Context){
	var newTask models.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	result, err := data.NewTask(newTask);
	if  err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, result)
}