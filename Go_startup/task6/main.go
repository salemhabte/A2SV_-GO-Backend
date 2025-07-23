package main

import (
	"task_manager/data"
	"task_manager/router"
)

func main(){
	data.ConnectDB()
	router.Router()
}