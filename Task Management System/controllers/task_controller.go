package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Task_Management_System/data"
	"Task_Management_System/models"
)

func GetTasks(c gin.Context, service data.TaskMangemetService){
	c.IndentedJSON(http.StatusOK, service.GetTasks())
}

func GetTask(c gin.Context, service data.TaskMangemetService){
	id := c.Param("id")
	task, err := service.GetTask(id)

	if err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message: ": "Task Not Found"})
		return 
	}

	c.IndentedJSON(http.StatusOK, task)
}

func CreateTask(c gin.Context, service data.TaskMangemetService){
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil{
		return
	}

	newTask = service.CreateTask(newTask)
	
	c.IndentedJSON(http.StatusOK, newTask)
}

func UpdateTask(c gin.Context, service data.TaskMangemetService){
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


func DeleteTask(c gin.Context, service data.TaskMangemetService){
	id := c.Param("id")
	if err := service.DeleteTask(id); err != nil{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message: ": "Task Not Found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message: ": "Task Deleted Successfully."})
}


