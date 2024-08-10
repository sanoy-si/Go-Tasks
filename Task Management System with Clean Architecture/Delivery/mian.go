package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/delivery/routers"
	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/repository"
	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/usecases"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getDatabase() (*mongo.Database, error){
	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil{
		return &mongo.Database{}, err
	}

	return client.Database(os.Getenv("DATABASE_NAME")), err
}

func getEnvPath() string{
	cwd, err := os.Getwd()
	if err != nil{
		log.Fatal(err.Error())
	}

	envPath := filepath.Join(filepath.Dir(cwd), ".env.template")

	return envPath
}


func main() {
	envPath := getEnvPath()
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Couldn't load .env file.")
	}

	r := gin.Default()

	db, err := getDatabase()
	if err != nil{
		log.Fatal(err.Error())
	}
	
	taskRepository := repositroy.NewTaskRepository(db, "tasks")
	taskUsecase := usecases.NewTaskUsecase(taskRepository)
	

	userRepository := repositroy.NewUserRepository(db, "users")
	userUsecase := usecases.NewUserUsecase(userRepository)

	routers.BindRouter(r, taskUsecase, userUsecase)

	r.Run("localhost:8080")
}