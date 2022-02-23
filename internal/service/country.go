package service

import (
	"errors"

	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
	"github.com/automation-as-a-service/internal/repository"
)

type CountryService interface {
	CreateCountry(country datastruct.Country, userID int64) (*int64, error)
	GetCountry(requestedCountryID int64) (*datastruct.Country, error)
	GetCountryByName(name string) (*datastruct.Country, error)
	UpdateCountry(country dto.Country, user *datastruct.Person) (*datastruct.Country, error)
	DeleteCountry(name string, userID int64) error
}

type countryService struct {
	dao repository.DAO
}

func NewCountryService(dao repository.DAO) CountryService {
	return &countryService{dao: dao}
}

func (c *countryService) CreateCountry(country datastruct.Country, userID int64) (*int64, error) {
	user, err := c.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		return nil, err
	}

	if user.Role == datastruct.ADMIN {
		id, err := c.dao.NewCountryQuery().CreateCountry(country)
		if err != nil {
			return nil, err
		}
		return id, nil
	}

	return nil, errors.New("you don't have access")
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

func (c *countryService) UpdateCountry(country dto.Country, user *datastruct.Person) (*datastruct.Country, error) {
	if user.Role == datastruct.ADMIN {
		updatedCountry, err := c.dao.NewCountryQuery().UpdateCountry(country)
		if err != nil {
			return nil, err
		}
		return updatedCountry, nil
	}
	return nil, errors.New("you don't have access")
}

func (c *countryService) DeleteCountry(name string, userID int64) error {
	user, err := c.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		return err
	}

	if user.Role == datastruct.ADMIN {
		err := c.dao.NewCountryQuery().DeleteCountry(name)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("you don't have access")
}
