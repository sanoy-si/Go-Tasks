package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/delivery/controllers"
	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/domain"
	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/infrastructure"
)

func BindRouter(r *gin.Engine, taskUsecase domain.TaskUsecase, userUsecase domain.UserUsecase) {
	r.GET("/tasks", infrastructure.AuthMiddleware(), controllers.GetTasks(taskUsecase))
	r.GET("/tasks/:id", infrastructure.AuthMiddleware(), controllers.GetTask(taskUsecase))
	r.POST("/tasks", infrastructure.AuthMiddleware(), infrastructure.AdminMiddleware(), controllers.CreateTask(taskUsecase))
	r.PUT("/tasks/:id", infrastructure.AuthMiddleware(), infrastructure.AdminMiddleware(), controllers.UpdateTask(taskUsecase))
	r.DELETE("/tasks/:id", infrastructure.AuthMiddleware(), infrastructure.AdminMiddleware(), controllers.DeleteTask(taskUsecase))
	r.POST("/users/:username/promote", infrastructure.AuthMiddleware(), infrastructure.AdminMiddleware(), controllers.PromoteUser(userUsecase))
	r.POST("/register", controllers.Register(userUsecase))
	r.POST("/login", controllers.Login(userUsecase))

}