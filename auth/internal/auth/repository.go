//go:generate mockgen -source repository.go -destination ./../../mocks/auth/repository.go -package auth_mock
package auth

import (
	"context"

	"github.com/evgsrkn/go-microservices-example/auth/internal/auth/model"
	"github.com/pkg/errors"
)

type Repository interface {
	GetUserByUsername(username string) (*model.User, error)
}

func (db *storage) GetUserByUsername(username string) (*model.User, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	var user model.User

	err := db.QueryRow(
		context.Background(),
		"SELECT id,login,password FROM users WHERE login=$1",
		username,
	).Scan(&user.ID, &user.Login, &user.Password)

	db.log.Debug(user.Login)
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}

	return &user, err
}
