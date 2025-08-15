package routers

import (
	"go_clean/Delivery/Controllers"
	"go_clean/Infrastructure"

	"github.com/gin-gonic/gin"
)

func Router(uc *controllers.UserController){
	router := gin.Default()

	
	router.POST("/register", uc.RegisterUser)
	router.POST("/login", uc.Login)


	auth := router.Group("/")
	auth.Use(Infrastructure.JWTAuthMiddleware())
	{
		auth.GET("/tasks/", uc.GetTask)
		auth.GET("/tasks/details/:id", uc.GetTaskById)
		auth.PUT("/tasks/:id", uc.UpdateTask)
		auth.POST("/tasks", uc.CreatedTask)

		// Only admins can delete
		auth.DELETE("/tasks/:id", Infrastructure.RoleMiddleware("admin"), uc.RemoveTask)
	}

	router.Run()
}