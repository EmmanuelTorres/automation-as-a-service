package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
)

type ProjectQuery interface {
	CreateProject(project datastruct.Project) (*int64, error)
	GetProject(name string) (*datastruct.Project, error)
	UpdateProject(project dto.Project) (*datastruct.Project, error)
	DeleteProject(id int64) error
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
	qb := pgQb().
		Select("id", "name").
		From(datastruct.ProjectTableName).
		Where(squirrel.Eq{"name": name})

	project := datastruct.Project{}
	err := qb.QueryRow().Scan(&project.ID, &project.Name)
	if err != nil {
		return nil, err
	}
	return &project, err
}

func (p *projectQuery) UpdateProject(project dto.Project) (*datastruct.Project, error) {
	qb := pgQb().
		Update(datastruct.ProjectTableName).
		SetMap(map[string]interface{}{
			"name": project.Name,
		}).
		Where(squirrel.Eq{"id": project.ID}).
		Suffix("RETURNING id, name")

	var updatedProject datastruct.Project
	err := qb.QueryRow().Scan(&updatedProject.ID, &updatedProject.Name)
	if err != nil {
		return nil, err
	}
	return &updatedProject, nil
}

func (p *projectQuery) DeleteProject(id int64) error {
	qb := pgQb().
		Delete(datastruct.ProjectTableName).
		From(datastruct.ProjectTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
