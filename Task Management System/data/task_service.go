package data

import (
	"Task_Management_System/models"
	"errors"
)

type TaskMangemetService interface{
	GetTasks() []models.Task
	GetTask(id int) (models.Task, error)
	CreateTask(newTask models.Task) models.Task
	UpdateTask(id int, updatedTask models.Task) (models.Task, error)
	DeleteTask(id int) error
}

type inMemoryTaskManagementService struct{
	tasks map[int]models.Task
	currentId int
}

func NewInMemoryTaskManagementService() *inMemoryTaskManagementService{
	return &inMemoryTaskManagementService{
		tasks: make(map[int]models.Task),
		currentId: 1,
	}
}

func (service *inMemoryTaskManagementService) GetTasks() []models.Task{
	allTasks := []models.Task{} 
	for _, task := range service.tasks{
		allTasks = append(allTasks, task) 
	} 
	return allTasks
}

func (service *inMemoryTaskManagementService) GetTask(id int) (models.Task, error){
	task, exists := service.tasks[id]

	if !exists{
		return models.Task{}, errors.New("task not found")
	}

	return task, nil
}

func (service *inMemoryTaskManagementService) CreateTask(newTask models.Task) models.Task{
	newTask.ID = service.currentId
	service.currentId += 1
	return newTask
}

func (service *inMemoryTaskManagementService) UpdateTask(id int, updatedTask models.Task) (models.Task, error){
	_, err := service.GetTask(id)
	if err != nil{
		return models.Task{}, err
	}

	updatedTask.ID = id
	service.tasks[id] = updatedTask

	return updatedTask, nil
}

func (service *inMemoryTaskManagementService) DeleteTask(id int) error{
	_, exists := service.tasks[id]

	if !exists{
		return errors.New("task not found")
	}

	delete(service.tasks, id)

	return nil
}



