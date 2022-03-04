package service

import (
	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
	"github.com/automation-as-a-service/internal/repository"
)

type CountryService interface {
	CreateCountry(country datastruct.Country) (*int64, error)
	GetCountry(requestedCountryID int64) (*datastruct.Country, error)
	GetCountryByName(name string) (*datastruct.Country, error)
	UpdateCountry(country dto.Country) (*datastruct.Country, error)
	DeleteCountry(id int64) error
}

type countryService struct {
	dao repository.DAO
}

func NewCountryService(dao repository.DAO) CountryService {
	return &countryService{dao: dao}
}

func (c *countryService) CreateCountry(country datastruct.Country) (*int64, error) {
	id, err := c.dao.NewCountryQuery().CreateCountry(country)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (c *countryService) GetCountry(requestedCountryID int64) (*datastruct.Country, error) {
	country, err := c.dao.NewCountryQuery().GetCountry(requestedCountryID)
	if err != nil {
		return nil, err
	}
	return country, err
}

func (c *countryService) GetCountryByName(name string) (*datastruct.Country, error) {
	country, err := c.dao.NewCountryQuery().GetCountryByName(name)
	if err != nil {
		return nil, err
	}
	return country, err
}

func (c *countryService) UpdateCountry(country dto.Country) (*datastruct.Country, error) {
	updatedCountry, err := c.dao.NewCountryQuery().UpdateCountry(country)
	if err != nil {
		return nil, err
	}
	return updatedCountry, nil
}

func (c *countryService) DeleteCountry(id int64) error {
	err := c.dao.NewCountryQuery().DeleteCountry(id)
	if err != nil {
		return err
	}
	return nil
}
