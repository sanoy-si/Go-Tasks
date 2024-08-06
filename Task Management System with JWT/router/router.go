package router

import (
	"github.com/gin-gonic/gin"
	"Task_Management_System_with_JWT/controllers"
	"Task_Management_System_with_JWT/data"

)

func BindRouter(r *gin.Engine, service *data.PersistentTaskManagementService){
	r.GET("/tasks", controllers.GetTasks(service))
	r.GET("tasks/:id", controllers.GetTask(service))
	r.POST("/tasks", controllers.CreateTask(service))
	r.PUT("tasks/:id", controllers.UpdateTask(service))
	r.DELETE("tasks/:id", controllers.DeleteTask(service))
	r.POST("register", controllers.Register(service))

}