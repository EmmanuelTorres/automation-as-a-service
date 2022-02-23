package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/automation-as-a-service/internal/datastruct"
)

type ProjectQuery interface {
	CreateProject(project datastruct.Project) (*int64, error)
	GetProject(name string) (*datastruct.Project, error)
	// UpdateProject(project dto.Project) (*datastruct.Project, error)
	// DeleteProject(id int64) error
}

type projectQuery struct{}

func (p *projectQuery) CreateProject(project datastruct.Project) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.ProjectTableName).
		Columns("name").
		Values(project.Name).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (p *projectQuery) GetProject(name string) (*datastruct.Project, error) {
	qb := pgQb().Select("id", "name").From(datastruct.ProjectTableName).Where(squirrel.Eq{"name": name})

	project := datastruct.Project{}
	err := qb.QueryRow().Scan(&project.ID, &project.Name)
	if err != nil {
		return nil, err
	}
	return &project, err
}
