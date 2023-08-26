package auth

import (
	"sync"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type storage struct {
	*pgx.Conn
	log  *zap.Logger
	lock *sync.Mutex
}

func NewStorage(db *pgx.Conn, logger *zap.Logger, lock *sync.Mutex) Repository {
	return &storage{db, logger, lock}
}
