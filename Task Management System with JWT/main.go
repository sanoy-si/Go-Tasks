package main

import (
	"Task_Management_System_with_JWT/data"
	"Task_Management_System_with_JWT/router"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load(".env.template")
	if err != nil{
		log.Fatal("Couldn't load .env file.")
	}

	r := gin.Default()
	taskService, err := data.NewPersistentTaskManagementService()
	if err != nil{
		log.Fatal(err.Error())
	}

	userService, err := data.NewMongoDBUserService()
	if err != nil{
		log.Fatal(err.Error())
	}
	
	router.BindRouter(r, taskService, userService)
	
	r.Run("localhost:8080")
}