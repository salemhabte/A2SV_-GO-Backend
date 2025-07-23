package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func Router(){
	router := gin.Default()

	
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.Login)


	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/tasks", controllers.GetTask)
		auth.GET("/tasks/:id", controllers.GetTaskById)
		auth.PUT("/tasks/:id", controllers.UpdateTask)
		auth.POST("/tasks", controllers.CreatedTask)

		// Only admins can delete
		auth.DELETE("/tasks/:id", middleware.RoleMiddleware("admin"), controllers.RemoveTask)
	}

	router.Run()
}