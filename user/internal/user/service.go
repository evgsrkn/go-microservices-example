package user

import (
	"github.com/pkg/errors"

	"github.com/evgsrkn/go-microservices-example/user/internal/user/model"
)

type (
	IService interface {
		GetAllUsers() ([]*model.User, error)
		GetUserById(id int) (*model.User, error)
		UpdateUser(user *model.User) (*model.User, error)
		DeleteUser(id int) error
		CreateUser(creds *model.User) error
	}

	service struct {
		repo IRepository
	}
)

func NewService(repo IRepository) IService {
	return &service{repo}
}

func (s *service) DeleteUser(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetAllUsers() ([]*model.User, error) {
	users, err := s.repo.GetAll()
	if err != nil {
		return nil, errors.Wrap(err, "can't get all users")
	}

	return users, nil

}

func (s *service) UpdateUser(user *model.User) (*model.User, error) {
	if err := s.repo.Update(user); err != nil {
		return nil, errors.Wrap(err, "can't update user")
	}

	user, err := s.GetUserById(user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) CreateUser(user *model.User) error {
	if err := s.repo.Create(user); err != nil {
		return errors.Wrap(err, "can't create user")
	}

	return nil
}

func (s *service) GetUserById(id int) (*model.User, error) {
	user, err := s.repo.GetById(id)
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}

	return user, nil
}
