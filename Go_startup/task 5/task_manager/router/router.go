package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func Router(){
	router := gin.Default()

	router.GET("/tasks", controllers.GetTask)
	router.GET("/tasks/:id", controllers.GetTaskById)

	router.PUT("/tasks/:id", controllers.UpdateTask)

	router.DELETE("/tasks/:id", controllers.RemoveTask)

	router.POST("/tasks", controllers.CreatedTask)

	router.Run()

}