package service

import (
	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
	"github.com/automation-as-a-service/internal/repository"
)

type BrandService interface {
	CreateBrand(designer dto.Brand) (*int64, error)
	GetBrand(requestedID int64) (*dto.Brand, error)
	GetBrandByName(name string) (*dto.Brand, error)
	UpdateBrand(designer dto.Brand) (*dto.Brand, error)
	DeleteBrand(id int64) error
}

type brandService struct {
	dao repository.DAO
}

func NewBrandService(dao repository.DAO) BrandService {
	return &brandService{dao: dao}
}

func (d *brandService) CreateBrand(brand dto.Brand) (*int64, error) {
	dsBrand := datastruct.Brand{
		Name:      brand.Name,
		CountryID: brand.CountryID,
	}

	id, err := d.dao.NewBrandQuery().CreateBrand(dsBrand)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (d *brandService) GetBrand(requestedID int64) (*dto.Brand, error) {
	designer, err := d.dao.NewBrandQuery().GetBrand(requestedID)
	if err != nil {
		return nil, err
	}

	dtoBrand := dto.Brand{
		ID:        designer.ID,
		Name:      designer.Name,
		CountryID: designer.CountryID,
	}

	return &dtoBrand, nil
}

func (d *brandService) GetBrandByName(name string) (*dto.Brand, error) {
	designer, err := d.dao.NewBrandQuery().GetBrandByName(name)
	if err != nil {
		return nil, err
	}

	dtoBrand := dto.Brand{
		ID:        designer.ID,
		Name:      designer.Name,
		CountryID: designer.CountryID,
	}

	return &dtoBrand, err
}

func (d *brandService) UpdateBrand(designer dto.Brand) (*dto.Brand, error) {
	updatedBrand, err := d.dao.NewBrandQuery().UpdateBrand(designer)
	if err != nil {
		return nil, err
	}

	dtoBrand := dto.Brand{
		ID:        updatedBrand.ID,
		Name:      updatedBrand.Name,
		CountryID: updatedBrand.CountryID,
	}

	return &dtoBrand, nil
}

func (d *brandService) DeleteBrand(id int64) error {
	err := d.dao.NewBrandQuery().DeleteBrand(id)
	if err != nil {
		return err
	}
	return nil
}
