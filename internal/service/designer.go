package service

import (
	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
	"github.com/automation-as-a-service/internal/repository"
)

type DesignerService interface {
	CreateDesigner(designer dto.Designer) (*int64, error)
	GetDesigner(requestedID int64) (*dto.Designer, error)
	GetDesignerByName(name string) (*dto.Designer, error)
	UpdateDesigner(designer dto.Designer) (*dto.Designer, error)
	DeleteDesigner(id int64) error
}

type designerService struct {
	dao repository.DAO
}

func NewDesignerService(dao repository.DAO) DesignerService {
	return &designerService{dao: dao}
}

func (d *designerService) CreateDesigner(designer dto.Designer) (*int64, error) {
	dsDesigner := datastruct.Designer{
		Name:      designer.Name,
		CountryID: designer.CountryID,
	}

	id, err := d.dao.NewDesignerQuery().CreateDesigner(dsDesigner)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (d *designerService) GetDesigner(requestedID int64) (*dto.Designer, error) {
	designer, err := d.dao.NewDesignerQuery().GetDesigner(requestedID)
	if err != nil {
		return nil, err
	}

	dtoDesigner := dto.Designer{
		ID:        designer.ID,
		Name:      designer.Name,
		CountryID: designer.CountryID,
	}

	return &dtoDesigner, nil
}

func (d *designerService) GetDesignerByName(name string) (*dto.Designer, error) {
	designer, err := d.dao.NewDesignerQuery().GetDesignerByName(name)
	if err != nil {
		return nil, err
	}

	dtoDesigner := dto.Designer{
		ID:        designer.ID,
		Name:      designer.Name,
		CountryID: designer.CountryID,
	}

	return &dtoDesigner, err
}

func (d *designerService) UpdateDesigner(designer dto.Designer) (*dto.Designer, error) {
	updatedDesigner, err := d.dao.NewDesignerQuery().UpdateDesigner(designer)
	if err != nil {
		return nil, err
	}

	dtoDesigner := dto.Designer{
		ID:        updatedDesigner.ID,
		Name:      updatedDesigner.Name,
		CountryID: updatedDesigner.CountryID,
	}

	return &dtoDesigner, nil
}

func (d *designerService) DeleteDesigner(id int64) error {
	err := d.dao.NewDesignerQuery().DeleteDesigner(id)
	if err != nil {
		return err
	}
	return nil
}
