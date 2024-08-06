package main

import (
	"github.com/gin-gonic/gin"
	"Task_Management_System_with_JWT/router"
	"Task_Management_System_with_JWT/data"
)

func main(){

	r := gin.Default()
	service := data.NewPersistentTaskManagementService()
	router.BindRouter(r, service)
	
	r.Run("localhost:8080")
}