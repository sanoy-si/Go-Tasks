package main

import (
	"github.com/gin-gonic/gin"
	"Task_Management_System/router"
	"Task_Management_System/data"
)

func main(){

	r := gin.Default()
	service := data.NewInMemoryTaskManagementService()
	router.BindRouter(r, service)
	
	r.Run("localhost:8080")
}