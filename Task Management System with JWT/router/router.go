package router

import (
	"github.com/gin-gonic/gin"
	"Task_Management_System/controllers"
	"Task_Management_System/data"

)

func BindRouter(r *gin.Engine, service data.TaskMangemetService){
	r.GET("/tasks", controllers.GetTasks(service))
	r.GET("tasks/:id", controllers.GetTask(service))
	r.POST("/tasks", controllers.CreateTask(service))
	r.PUT("tasks/:id", controllers.UpdateTask(service))
	r.DELETE("tasks/:id", controllers.DeleteTask(service))
}