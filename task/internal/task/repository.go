package task

import (
	"context"
	"task/internal/task/model"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type RepositoryI interface {
	GetById(id int) (*model.Task, error)
	GetAll() ([]*model.Task, error)
	Update(task *model.Task) error
	Delete(id int) error
	Create(task *model.Task) error
}

func (db *storage) GetById(id int) (*model.Task, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	var task model.Task

	err := db.QueryRow(
		context.Background(),
		"SELECT * FROM tasks WHERE id=$1",
		id,
	).Scan(
		&task.ID,
		&task.Name,
		&task.Description,
		&task.Project_id,
		&task.User_id,
		&task.Status,
	)

	if err != nil {
		return nil, errors.Wrap(err, "task not found")
	}

	db.log.Debug("Found task", zap.Int("id", id))

	return &task, nil
}

func (db *storage) GetAll() ([]*model.Task, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	var tasks []*model.Task

	rows, err := db.Query(context.Background(), "SELECT * FROM tasks")
	defer rows.Close()
	if err != nil {
		return nil, errors.Wrap(err, "can't get all tasks")
	}

	for rows.Next() {
		var task *model.Task
		err := rows.Scan(
			&task.ID,
			&task.Name,
			&task.Description,
			&task.Project_id,
			&task.User_id,
			&task.Status,
		)
		if err != nil {
			return nil, errors.Wrap(err, "can't get all tasks")
		}

		tasks = append(tasks, task)
	}

	db.log.Debug("Found all tasks")

	return tasks, nil
}

func (db *storage) Update(task *model.Task) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	rows, err := db.Query(
		context.Background(),
		"UPDATE tasks set name=$1, description=$2, project_id=$3, user_id=$4, status=$5 WHERE id=$6",
		task.Name, task.Description, task.Project_id, task.User_id, task.Status, task.ID,
	)
	rows.Close()

	if err != nil {
		return errors.Wrap(err, "can't update task")
	}

	if err := rows.Err(); err != nil {
		return errors.Wrap(err, "can't update task")
	}

	return nil
}

func (db *storage) Create(task *model.Task) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	rows, err := db.Query(
		context.Background(),
		"INSERT INTO tasks(name, description, project_id, user_id, status) VALUES($1,$2,$3,$4) RETURNING id",
		task.Name,
		task.Description,
		task.Project_id,
		task.User_id,
		task.Status,
	)
	rows.Close()

	if err != nil {
		return errors.Wrap(err, "can't create task")
	}

	return nil
}

func (db *storage) Delete(id int) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	err := db.QueryRow(
		context.Background(),
		"DELETE FROM tasks WHERE id=$1",
		id,
	).Scan()

	if err != nil {
		return errors.Wrap(err, "can't delete task")
	}

	return nil
}
