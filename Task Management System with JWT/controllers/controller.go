package controllers

import (
	"Task_Management_System_with_JWT/data"
	"Task_Management_System_with_JWT/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()
func GetTasks(service *data.PersistentTaskManagementService) gin.HandlerFunc{
	return func(c *gin.Context){
		c.IndentedJSON(http.StatusOK, service.GetTasks())
	}
}

func GetTask(service *data.PersistentTaskManagementService) gin.HandlerFunc{
	return func(c *gin.Context){
			id := c.Param("id")
			task, err := service.GetTask(id)

			if err != nil{
				c.IndentedJSON(http.StatusNotFound, gin.H{"message: ": "Task Not Found"})
				return 
			}

			c.IndentedJSON(http.StatusOK, task)
			}
}

func CreateTask(service *data.PersistentTaskManagementService) gin.HandlerFunc{
	return func(c *gin.Context){
			var newTask models.Task
			if err := c.BindJSON(&newTask); err != nil{
				return
			}

			newTask = service.CreateTask(newTask)
			
			c.IndentedJSON(http.StatusCreated, newTask)
		}
	
}

func UpdateTask(service *data.PersistentTaskManagementService) gin.HandlerFunc{
	return func(c *gin.Context){
			var updatedTask models.Task
			
			if err := c.BindJSON(&updatedTask); err != nil{
				return
			}
			
			updatedTaskid := c.Param("id")

			updatedTask, err := service.UpdateTask(updatedTaskid, updatedTask)
			if err != nil{
				c.IndentedJSON(http.StatusNotFound, gin.H{"message: ": "Task Not Found"})
				return
			}

			c.IndentedJSON(http.StatusOK, updatedTask)
		}


}


func DeleteTask(service *data.PersistentTaskManagementService) gin.HandlerFunc{
	return func (c *gin.Context){
			id := c.Param("id")
			if err := service.DeleteTask(id); err != nil{
				c.IndentedJSON(http.StatusNotFound, gin.H{"message: ": "Task Not Found"})
				return
			}

			c.IndentedJSON(http.StatusOK, gin.H{"message: ": "Task Deleted Successfully."})
		}
}


func Register(service *data.PersistentTaskManagementService) gin.HandlerFunc{
	return func(c *gin.Context){

		var newUser models.User
		if err := c.BindJSON(&newUser); err != nil{
			return
		}
		
		if err := validate.Struct(newUser); err != nil{
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}


		insertedId, err :=  service.Register(newUser)
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

type credentials struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}
func Login(service *data.PersistentTaskManagementService) gin.HandlerFunc{
	return func(c *gin.Context){
		var inputCredentials credentials
		if err := c.BindJSON(&inputCredentials); err != nil{
			return
		}

		token, err := service.Login(inputCredentials.Username, inputCredentials.Password)
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