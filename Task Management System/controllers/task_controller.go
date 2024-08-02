package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Task_Management_System/data"
	"Task_Management_System/models"
)

func GetTasks(service data.TaskMangemetService) gin.HandlerFunc{
	return func(c *gin.Context){
		c.IndentedJSON(http.StatusOK, service.GetTasks())
	}
}

func GetTask(service data.TaskMangemetService) gin.HandlerFunc{
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

func CreateTask(service data.TaskMangemetService) gin.HandlerFunc{
	return func(c *gin.Context){
			var newTask models.Task
			if err := c.BindJSON(&newTask); err != nil{
				return
			}

			newTask = service.CreateTask(newTask)
			
			c.IndentedJSON(http.StatusOK, newTask)
		}
	
}

func UpdateTask(service data.TaskMangemetService) gin.HandlerFunc{
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


func DeleteTask(service data.TaskMangemetService) gin.HandlerFunc{
	return func (c *gin.Context){
			id := c.Param("id")
			if err := service.DeleteTask(id); err != nil{
				c.IndentedJSON(http.StatusNotFound, gin.H{"message: ": "Task Not Found"})
				return
			}

			c.IndentedJSON(http.StatusOK, gin.H{"message: ": "Task Deleted Successfully."})
		}
}


