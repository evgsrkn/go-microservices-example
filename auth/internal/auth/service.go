//go:generate mockgen -source service.go -destination ./../../mocks/auth/service.go -package auth_mock
package auth

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/evgsrkn/go-microservices-example/auth/internal/user"
	"github.com/evgsrkn/go-microservices-example/auth/pkg/jwt"
	"github.com/evgsrkn/go-microservices-example/auth/pkg/pb"
	user_pb "github.com/evgsrkn/go-microservices-example/user/pkg/pb"
	"golang.org/x/crypto/bcrypt"
)

type (
	IService interface {
		IsUserExist(username string) bool
		RegisterUser(creds *pb.BasicCredentials) (string, error)
		ValidateUser(creds *pb.Credentials) (string, error)
	}

	service struct {
		repo       Repository
		userClient user.Client
	}
)

const hashRound = 14

func NewService(repo Repository, userClient user.Client) IService {
	return &service{repo, userClient}
}

func (s *service) IsUserExist(username string) bool {
	if _, err := s.repo.GetUserByUsername(username); err != nil {
		return false
	}
	return true
}

func (s *service) RegisterUser(creds *pb.BasicCredentials) (string, error) {
	u := user_pb.User{
		Login:    creds.Username,
		Password: creds.Password,
		Name:     creds.Name,
		Role:     creds.Role,
	}

	hashed, errHash := bcrypt.GenerateFromPassword([]byte(u.Password), hashRound)
	if errHash != nil {
		return "", errHash
	}

	u.Password = string(hashed)

	if _, err := s.userClient.CreateUser(context.Background(), &u); err != nil {
		return "", errors.Wrap(err, "cannot create user")
	}

	token, err := jwt.GenJWT(fmt.Sprint(u.Id), u.Name)
	if err != nil {
		return "", errors.Wrap(err, "cannot generate JWT")
	}

	return token, nil
}

func (s *service) ValidateUser(creds *pb.Credentials) (string, error) {
	user, _ := s.repo.GetUserByUsername(creds.Username)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))

	if err != nil {
		return "", errors.Wrap(err, "wrong password")
	}

	token, err := jwt.GenJWT(fmt.Sprint(user.ID), user.Login)
	if err != nil {
		return "", errors.Wrap(err, "cannot generate JWT")
	}

	return token, nil
}
