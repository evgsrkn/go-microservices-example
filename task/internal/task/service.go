package task

import (
	"context"

	"github.com/evgsrkn/go-microservices-example/task/internal/task/model"
	"github.com/evgsrkn/go-microservices-example/task/internal/user"
	"github.com/evgsrkn/go-microservices-example/user/pkg/pb"

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
		repo       IRepository
		userClient user.Client
	}
)

func NewService(repo IRepository, userClient user.Client) IService {
	return &service{repo, userClient}
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
	if err := s.checkUserExistence(task.User_id); err != nil {
		return nil, err
	}

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
	if err := s.checkUserExistence(task.User_id); err != nil {
		return err
	}

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

func (s *service) checkUserExistence(id int) error {
	_, err := s.userClient.GetUserById(
		context.Background(),
		&pb.UserWithID{Id: int64(id)},
	)

	if err != nil {
		return errors.Wrap(err, "user doesn`t exist")
	}

	return nil
}
