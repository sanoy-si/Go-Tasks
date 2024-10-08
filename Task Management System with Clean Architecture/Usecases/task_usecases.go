package usecases

import (
	"context"

	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/domain"
	"github.com/sanoy-si/Task_Management_System_with_Clean_Architecture/infrastructure"
)

type TaskUsecase struct {
	repository domain.TaskRepository
}

func NewTaskUsecase(taskRepository domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		repository: taskRepository,
	}
}

func (taskUsecase *TaskUsecase) GetTasks() ([]domain.Task, error) {
	return taskUsecase.repository.GetTasks(context.TODO())
}

func (taskUsecase *TaskUsecase) GetTaskByID(id string) (domain.Task, error) {
	return taskUsecase.repository.GetTaskByID(id, context.TODO())
}

func (taskUsecase *TaskUsecase) CreateTask(newTask domain.Task) (domain.Task, error) {
	newTask.ID = infrastructure.GenerateID()
	return taskUsecase.repository.CreateTask(newTask, context.TODO())
}

func (taskUsecase *TaskUsecase) UpdateTask(id string, updatedTask domain.Task) (domain.Task, error) {
	return taskUsecase.repository.UpdateTask(id, updatedTask, context.TODO())
}

func (taskUsecase *TaskUsecase) DeleteTask(id string) error {
	return taskUsecase.repository.DeleteTask(id, context.TODO())
}
