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
	r.POST("/tasks", middleware.AuthMiddleware(), middleware.AdminMiddleware(),controllers.CreateTask(taskService))
	r.PUT("/tasks/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.UpdateTask(taskService))
	r.DELETE("/tasks/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.DeleteTask(taskService))
	r.POST("/users/:username/promote", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controllers.PromoteUser(userService))
	r.POST("/register", controllers.Register(userService))
	r.POST("/login", controllers.Login(userService))

}