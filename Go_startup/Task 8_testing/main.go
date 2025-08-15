package main

import (
	 "go_clean/Delivery/Controllers"
	"go_clean/Delivery/routers"
	"go_clean/Repositories"
	 "go_clean/Usecases"
)

func main() {
	repositories.ConnectDB()
	userRepo := repositories.NewUserRepo()
	taskRepo := repositories.NewTaskRepo()
	taskusecase := usecases.NewTaskUsecase(taskRepo)
	userusecase := usecases.NewUserusecase(userRepo)
	usercontroller := controllers.NewController(*taskusecase,*userusecase)
	routers.Router(usercontroller)
}