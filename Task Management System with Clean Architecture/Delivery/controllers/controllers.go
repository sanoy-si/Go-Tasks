package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/domain"
	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/infrastructure"
)

func CreateTask(taskUsecase domain.TaskUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newTask domain.Task
		if err := c.BindJSON(&newTask); err != nil{
			return
		}
		
		if err := infrastructure.ValidateTask(&newTask); err != nil{
			c.IndentedJSON(http.StatusBadRequest, err.Error())
			return 
		}

		newTask, err := taskUsecase.CreateTask(newTask)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(http.StatusOK, newTask)
	}
}


func GetTasks(taskUsecase domain.TaskUsecase) gin.HandlerFunc{
	return func(c *gin.Context){
		tasks, err := taskUsecase.GetTasks()
		if err != nil{
			c.IndentedJSON(http.StatusInternalServerError, err)
		}

		c.IndentedJSON(http.StatusOK, tasks)
	}
}


func GetTask(taskUsecase domain.TaskUsecase) gin.HandlerFunc{
	return func(c *gin.Context){
			id := c.Param("id")
			task, err := taskUsecase.GetTaskByID(id)

			if err != nil{
				c.IndentedJSON(http.StatusNotFound, gin.H{"message: ": "Task Not Found"})
				return 
			}

			c.IndentedJSON(http.StatusOK, task)
			}
}

func UpdateTask(taskUsecase domain.TaskUsecase) gin.HandlerFunc{
	return func(c *gin.Context){
			var updatedTask domain.Task
			
			if err := c.BindJSON(&updatedTask); err != nil{
				return
			}
			
			updatedTaskid := c.Param("id")

			updatedTask, err := taskUsecase.UpdateTask(updatedTaskid, updatedTask)
			if err != nil{
				switch err.Error(){
				case "mongo: no documents in result":
					c.IndentedJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
					
				default:
					c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
				}
				return
			}

			c.IndentedJSON(http.StatusOK, updatedTask)
		}

}


func DeleteTask(taskUsecase domain.TaskUsecase) gin.HandlerFunc{
	return func (c *gin.Context){
			id := c.Param("id")
			if err := taskUsecase.DeleteTask(id); err != nil{
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task Not Found"})
				return
			}

			c.IndentedJSON(http.StatusNoContent, gin.H{"message":"Task Deleted Successfully."})
		}
}


func Register(userUsecase domain.UserUsecase) gin.HandlerFunc{
	return func(c *gin.Context){
		var newUser domain.User
		if err := c.BindJSON(&newUser); err != nil{
			return
		}
		
		if err := infrastructure.ValidateUser(&newUser); err != nil{
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}

		insertedId, err :=  userUsecase.Register(newUser)
		if err != nil{
			if err.Error() == "email already exists"{
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
				return
			}

			if err.Error() == "username already exists"{
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
				return
			}

			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}

		c.IndentedJSON(http.StatusCreated, gin.H{"insertedID":insertedId})
	}

}


func Login(userUsecase domain.UserUsecase) gin.HandlerFunc{
	return func(c *gin.Context){
		var userCredentials domain.UserCredentials
		if err := c.BindJSON(&userCredentials); err != nil{
			return
		}

		token, err := userUsecase.Login(userCredentials)
		if err != nil{
			switch err.Error(){
			case "mongo: no documents in result", "crypto/bcrypt: hashedPassword is not the hash of the given password":
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"invalid credentials"})
				
			default:
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			}
			return
			
		}
		
		c.IndentedJSON(http.StatusOK, gin.H{"token":token})
	}
}

func PromoteUser(userUsecase domain.UserUsecase) gin.HandlerFunc{
	return func (c *gin.Context){
		userName := c.Param("username")
		err := userUsecase.PromoteUser(userName)

		if err != nil{
			switch err.Error(){
			case "user not found", "user is already an admin":
				c.IndentedJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
				
			default:
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			}
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"message":"success"})
}
}