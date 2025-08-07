package service

import (
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

	s.repo.Create(task)
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



