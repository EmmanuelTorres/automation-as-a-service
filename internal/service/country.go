package service

import (
	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
	"github.com/automation-as-a-service/internal/repository"
)

type CountryService interface {
	CreateCountry(country dto.Country) (*int64, error)
	GetCountry(requestedCountryID int64) (*dto.Country, error)
	GetCountryByName(name string) (*dto.Country, error)
	UpdateCountry(country dto.Country) (*dto.Country, error)
	DeleteCountry(id int64) error
}

type countryService struct {
	dao repository.DAO
}

func NewCountryService(dao repository.DAO) CountryService {
	return &countryService{dao: dao}
}

func (c *countryService) CreateCountry(country dto.Country) (*int64, error) {
	countryInfo := datastruct.Country{
		Name: country.Name,
	}

	id, err := c.dao.NewCountryQuery().CreateCountry(countryInfo)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (c *countryService) GetCountry(requestedCountryID int64) (*dto.Country, error) {
	country, err := c.dao.NewCountryQuery().GetCountry(requestedCountryID)
	if err != nil {
		return nil, err
	}

	fullCountry := dto.Country{
		ID:   country.ID,
		Name: country.Name,
	}

	return &fullCountry, err
}

func (c *countryService) GetCountryByName(name string) (*dto.Country, error) {
	country, err := c.dao.NewCountryQuery().GetCountryByName(name)
	if err != nil {
		return nil, err
	}

	fullCountry := dto.Country{
		ID:   country.ID,
		Name: country.Name,
	}

	return &fullCountry, err
}

func (c *countryService) UpdateCountry(country dto.Country) (*dto.Country, error) {
	updatedCountry, err := c.dao.NewCountryQuery().UpdateCountry(country)
	if err != nil {
		return nil, err
	}

	fullCountry := dto.Country{
		ID:   updatedCountry.ID,
		Name: updatedCountry.Name,
	}

	return &fullCountry, nil
}

func (c *countryService) DeleteCountry(id int64) error {
	err := c.dao.NewCountryQuery().DeleteCountry(id)
	if err != nil {
		return err
	}
	return nil
}
