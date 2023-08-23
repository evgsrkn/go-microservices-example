package project

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

func NewStorage(db *pgx.Conn, log *zap.Logger, lock *sync.Mutex) *storage {
	return &storage{
		db,
		log,
		lock,
	}
}
