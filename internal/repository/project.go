package repository

import (
	"github.com/automation-as-a-service/internal/datastruct"
)

type ProjectQuery interface {
	CreateProject(project datastruct.Project) (*int64, error)
	// GetProject(id int64) (*datastruct.Project, error)
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
