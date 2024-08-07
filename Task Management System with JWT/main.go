package main

import (
	"github.com/gin-gonic/gin"
	"Task_Management_System_with_JWT/router"
	"Task_Management_System_with_JWT/data"
)

func main(){

	r := gin.Default()
	taskService := data.NewPersistentTaskManagementService()
	userService := data.NewMongoDBService()
	router.BindRouter(r, taskService, userService)
	
	r.Run("localhost:8080")
}