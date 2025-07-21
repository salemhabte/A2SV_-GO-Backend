package main

import (
	"task_manager/router"
	"task_manager/data"

)

func main(){
	data.ConnectDB()
	router.Router()
}