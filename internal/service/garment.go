package service

import (
	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
	"github.com/automation-as-a-service/internal/repository"
)

type GarmentService interface {
	CreateGarment(garment dto.Garment) (*int64, error)
	GetGarment(requestedID int64) (*dto.Garment, error)
	GetGarmentByCode(code string) (*dto.Garment, error)
	UpdateGarment(garment dto.Garment) (*dto.Garment, error)
	DeleteGarment(id int64) error
}

type garmentService struct {
	dao repository.DAO
}

func NewGarmentService(dao repository.DAO) GarmentService {
	return &garmentService{dao: dao}
}

func (d *garmentService) CreateGarment(garment dto.Garment) (*int64, error) {
	dsGarment := datastruct.Garment{
		Code:       garment.Code,
		DesignerID: garment.DesignerID,
		BrandID:    garment.BrandID,
	}

	id, err := d.dao.NewGarmentQuery().CreateGarment(dsGarment)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (d *garmentService) GetGarment(requestedID int64) (*dto.Garment, error) {
	garment, err := d.dao.NewGarmentQuery().GetGarment(requestedID)
	if err != nil {
		return nil, err
	}

	dtoGarment := dto.Garment{
		ID:         garment.ID,
		Code:       garment.Code,
		DesignerID: garment.DesignerID,
		BrandID:    garment.BrandID,
	}

	return &dtoGarment, nil
}

func (d *garmentService) GetGarmentByCode(code string) (*dto.Garment, error) {
	garment, err := d.dao.NewGarmentQuery().GetGarmentByCode(code)
	if err != nil {
		return nil, err
	}

	dtoGarment := dto.Garment{
		ID:         garment.ID,
		Code:       garment.Code,
		DesignerID: garment.DesignerID,
		BrandID:    garment.BrandID,
	}

	return &dtoGarment, err
}

func (d *garmentService) UpdateGarment(garment dto.Garment) (*dto.Garment, error) {
	updatedGarment, err := d.dao.NewGarmentQuery().UpdateGarment(garment)
	if err != nil {
		return nil, err
	}

	dtoGarment := dto.Garment{
		ID:         updatedGarment.ID,
		Code:       garment.Code,
		DesignerID: garment.DesignerID,
		BrandID:    garment.BrandID,
	}

	return &dtoGarment, nil
}

func (d *garmentService) DeleteGarment(id int64) error {
	err := d.dao.NewGarmentQuery().DeleteGarment(id)
	if err != nil {
		return err
	}
	return nil
}
