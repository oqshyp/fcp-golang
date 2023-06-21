package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type TaskService interface {
	Store(task *model.Task) error
	Update(task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskService struct {
	taskRepository repo.TaskRepository
}

func NewTaskService(taskRepository repo.TaskRepository) TaskService {
	return &taskService{taskRepository}
}

func (c *taskService) Store(task *model.Task) error {
	err := c.taskRepository.Store(task)
	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) Update(task *model.Task) error {
	err := s.taskRepository.Update(task)
	return err // TODO: replace this
}

func (s *taskService) Delete(id int) error {
	err := s.taskRepository.Delete(id)
	return err // TODO: replace this
}

func (s *taskService) GetByID(id int) (*model.Task, error) {
	task, err := s.taskRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) GetList() ([]model.Task, error) {
	result, err := s.taskRepository.GetList()
	return result, err // TODO: replace this
}

func (s *taskService) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	result, err := s.taskRepository.GetTaskCategory(id)
	return result, err // TODO: replace this
}
