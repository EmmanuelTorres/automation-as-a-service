package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
)

type CountryQuery interface {
	CreateCountry(country datastruct.Country) (*int64, error)
	GetCountry(id int64) (*datastruct.Country, error)
	GetCountryByName(name string) (*datastruct.Country, error)
	UpdateCountry(country dto.Country) (*datastruct.Country, error)
	DeleteCountry(id int64) error
}

type countryQuery struct{}

func (c *countryQuery) CreateCountry(country datastruct.Country) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.CountryTableName).
		Columns("name").
		Values(country.Name).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (c *countryQuery) GetCountry(id int64) (*datastruct.Country, error) {
	qb := pgQb().
		Select("id", "name").
		From(datastruct.CountryTableName).
		Where(squirrel.Eq{"id": id})

	country := datastruct.Country{}
	err := qb.QueryRow().Scan(&country.ID, &country.Name)
	if err != nil {
		return nil, err
	}
	return &country, err
}

func (c *countryQuery) GetCountryByName(name string) (*datastruct.Country, error) {
	qb := pgQb().
		Select("id", "name").
		From(datastruct.CountryTableName).
		Where(squirrel.Eq{"name": name})

	country := datastruct.Country{}
	err := qb.QueryRow().Scan(&country.ID, &country.Name)
	if err != nil {
		return nil, err
	}
	return &country, err
}

func (c *countryQuery) UpdateCountry(country dto.Country) (*datastruct.Country, error) {
	qb := pgQb().
		Update(datastruct.CountryTableName).
		SetMap(map[string]interface{}{
			"name": country.Name,
		}).
		Where(squirrel.Eq{"id": country.ID}).
		Suffix("RETURNING id, name")

	var updatedCountry datastruct.Country
	err := qb.QueryRow().Scan(&updatedCountry.ID, &updatedCountry.Name)
	if err != nil {
		return nil, err
	}
	return &updatedCountry, nil
}

func (c *countryQuery) DeleteCountry(id int64) error {
	qb := pgQb().
		Delete(datastruct.CountryTableName).
		From(datastruct.CountryTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
