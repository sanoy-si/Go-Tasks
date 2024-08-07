package router

import (
	"github.com/gin-gonic/gin"
	"Task_Management_System_with_JWT/controllers"
	"Task_Management_System_with_JWT/data"
	"Task_Management_System_with_JWT/middleware"


)

func BindRouter(r *gin.Engine, taskService *data.PersistentTaskManagementService, userService data.UserService){
	r.GET("/tasks", middleware.AuthMiddleware(), controllers.GetTasks(taskService))
	r.GET("/tasks/:id", middleware.AuthMiddleware(), controllers.GetTask(taskService))
	r.POST("/tasks", middleware.AuthMiddleware(), controllers.CreateTask(taskService))
	r.PUT("/tasks/:id", middleware.AuthMiddleware(), controllers.UpdateTask(taskService))
	r.DELETE("/tasks/:id", middleware.AuthMiddleware(), controllers.DeleteTask(taskService))
	r.POST("/promote/:username", middleware.AuthMiddleware(), controllers.PromoteUser(userService))
	r.POST("/register", controllers.Register(userService))
	r.POST("/login", controllers.Login(userService))

}