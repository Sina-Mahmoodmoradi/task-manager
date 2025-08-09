package service

import (
	"errors"
	"github.com/Sina-Mahmoodmoradi/task-manager/internal/domain"
)

type taskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) domain.TaskService {
	return &taskService{
		repo: repo,
	}
}


func (s *taskService)CreateTask(userID uint, title string, description string ) (*domain.Task, error){
	
	task := &domain.Task{
		Title:title,
		Description: description,
		UserID: userID,
	}

	if err:=s.repo.Create(task);err!=nil{
		return nil,err
	}

	return task, nil
}


func (s *taskService)GetTaskByID(taskID uint) (*domain.Task,error){
	return s.repo.GetByID(taskID)

}

func (s *taskService)DeleteTask(taskID uint) error {
	return s.repo.Delete(taskID)
}


func (s *taskService)ListTasksByUser(UserID uint) ([]domain.Task,error){
	return s.repo.GetByUserId(UserID)
}


func (s *taskService)UpdateTask(taskID uint, userID uint, updates *domain.TaskUpdate ) (*domain.Task, error){
	task,err := s.repo.GetByID(taskID)
	if err != nil {
		return nil, err
	}
	if task.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	if updates.Title!=nil{
		task.Title = *updates.Title
	}

	if updates.Description!=nil{
		task.Description = *updates.Description
	}

	if updates.Done!=nil{
		task.Done = *updates.Done
	}

	if err := s.repo.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}
