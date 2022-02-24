package service

import (
	"errors"

	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
	"github.com/automation-as-a-service/internal/repository"
)

type ProjectService interface {
	CreateProject(project datastruct.Project, userID int64) (*int64, error)
	GetProject(name string) (*datastruct.Project, error)
	UpdateProject(project dto.Project) (*datastruct.Project, error)
	DeleteProject(id, userID int64) error
}

type projectService struct {
	dao repository.DAO
}

func NewProjectService(dao repository.DAO) ProjectService {
	return &projectService{dao: dao}
}

// Creates a project on the database
func (p *projectService) CreateProject(project datastruct.Project, userID int64) (*int64, error) {
	// Retrieve the user
	user, err := p.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		return nil, err
	}

	// Only admins can create projects
	if user.Role == datastruct.ADMIN {
		// Create the project
		id, err := p.dao.NewProjectQuery().CreateProject(project)
		if err != nil {
			return nil, err
		}
		return id, nil
	}
	return nil, errors.New("you don't have access")
}

func (p *projectService) GetProject(name string) (*datastruct.Project, error) {
	// Retrieve the project from the database
	project, err := p.dao.NewProjectQuery().GetProject(name)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p *projectService) UpdateProject(project dto.Project) (*datastruct.Project, error) {
	updatedProject, err := p.dao.NewProjectQuery().UpdateProject(project)
	if err != nil {
		return nil, err
	}
	return updatedProject, nil
}

func (p *projectService) DeleteProject(id, userID int64) error {
	// Retrieve the user
	user, err := p.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		return err
	}

	// Only admins can delete projects
	if user.Role == datastruct.ADMIN {
		// Delete the project
		err := p.dao.NewProjectQuery().DeleteProject(id)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("you don't have access")
}
