package project

import (
	"context"

	"github.com/evgsrkn/go-microservices-example/project/internal/project/model"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type IRepository interface {
	GetById(id int) (*model.Project, error)
	GetAll() ([]*model.Project, error)
	Update(project *model.Project) error
	Delete(id int) error
	Create(project *model.Project) error
}

func (db *storage) GetById(id int) (*model.Project, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	var project model.Project

	err := db.QueryRow(
		context.Background(),
		"SELECT * FROM projects WHERE id=$1",
		id,
	).Scan(
		&project.ID,
		&project.Name,
		&project.Description,
	)

	if err != nil {
		return nil, errors.Wrap(err, "project not found")
	}

	db.log.Debug("Found project", zap.Int("id", id))

	return &project, nil
}

func (db *storage) GetAll() ([]*model.Project, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	var projects []*model.Project

	rows, err := db.Query(context.Background(), "SELECT * FROM projects")
	if err != nil {
		return nil, errors.Wrap(err, "can't get all projects")
	}
	defer rows.Close()

	for rows.Next() {
		var project model.Project
		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Description,
		)
		if err != nil {
			return nil, errors.Wrap(err, "can't get all projects")
		}

		projects = append(projects, &project)
	}

	db.log.Debug("Found all projects")

	return projects, nil
}

func (db *storage) Update(project *model.Project) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	rows, err := db.Query(
		context.Background(),
		"UPDATE projects set name=$1, description=$2 WHERE id=$3",
		project.Name, project.Description, project.ID,
	)
	rows.Close()

	if err != nil {
		return errors.Wrap(err, "can't update project")
	}

	if err := rows.Err(); err != nil {
		return errors.Wrap(err, "can't update project")
	}

	return nil
}

func (db *storage) Create(project *model.Project) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	rows, err := db.Query(
		context.Background(),
		"INSERT INTO projects(name, description) VALUES($1,$2) RETURNING id",
		project.Name,
		project.Description,
	)
	rows.Close()

	if err != nil {
		return errors.Wrap(err, "can't create project")
	}

	return nil
}

func (db *storage) Delete(id int) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	err := db.QueryRow(
		context.Background(),
		"DELETE FROM projects WHERE id=$1",
		id,
	).Scan()

	if err != nil {
		return errors.Wrap(err, "can't delete project")
	}

	return nil
}
