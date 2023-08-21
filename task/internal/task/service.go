package task

import (
	"task/internal/task/model"

	"github.com/pkg/errors"
)

type (
	IService interface {
		GetAllTasks() ([]*model.Task, error)
		GetTaskById(id int) (*model.Task, error)
		UpdateTask(project *model.Task) (*model.Task, error)
		DeleteTask(id int) error
		CreateTask(creds *model.Task) error
	}

	service struct {
		repo IRepository
	}
)

func NewService(repo IRepository) IService {
	return &service{repo}
}

func (s *service) DeleteTask(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAllTasks() ([]*model.Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "can't get all tasks")
	}

	return tasks, nil

}

func (s *service) UpdateTask(task *model.Task) (*model.Task, error) {
	if err := s.repo.Update(task); err != nil {
		return nil, errors.Wrap(err, "can't update task")
	}

	task, err := s.GetTaskById(task.ID)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *service) CreateTask(task *model.Task) error {
	err := s.repo.Create(task)
	if err != nil {
		return errors.Wrap(err, "can't create task")
	}

	return nil
}

func (s *service) GetTaskById(id int) (*model.Task, error) {
	task, err := s.repo.GetById(id)
	if err != nil {
		return nil, errors.Wrap(err, "task not found")
	}

	return task, nil
}
