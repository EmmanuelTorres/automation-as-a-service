package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
)

type BrandQuery interface {
	CreateBrand(Brand datastruct.Brand) (*int64, error)
	GetBrand(id int64) (*datastruct.Brand, error)
	GetBrandByName(name string) (*datastruct.Brand, error)
	UpdateBrand(Brand dto.Brand) (*datastruct.Brand, error)
	DeleteBrand(id int64) error
}

type brandQuery struct{}

func (b *brandQuery) CreateBrand(Brand datastruct.Brand) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.BrandTableName).
		Columns("name", "country_id").
		Values(Brand.Name, Brand.CountryID).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (b *brandQuery) GetBrand(id int64) (*datastruct.Brand, error) {
	qb := pgQb().
		Select("id", "name", "country_id").
		From(datastruct.BrandTableName).
		Where(squirrel.Eq{"id": id})

	Brand := datastruct.Brand{}
	err := qb.QueryRow().Scan(&Brand.ID, &Brand.Name, &Brand.CountryID)
	if err != nil {
		return nil, err
	}
	return &Brand, err
}

func (b *brandQuery) GetBrandByName(name string) (*datastruct.Brand, error) {
	qb := pgQb().
		Select("id", "name", "country_id").
		From(datastruct.BrandTableName).
		Where(squirrel.Eq{"name": name})

	Brand := datastruct.Brand{}
	err := qb.QueryRow().Scan(&Brand.ID, &Brand.Name, &Brand.CountryID)
	if err != nil {
		return nil, err
	}
	return &Brand, err
}

func (b *brandQuery) UpdateBrand(Brand dto.Brand) (*datastruct.Brand, error) {
	qb := pgQb().
		Update(datastruct.BrandTableName).
		SetMap(map[string]interface{}{
			"name":       Brand.Name,
			"country_id": Brand.CountryID,
		}).
		Where(squirrel.Eq{"id": Brand.ID}).
		Suffix("RETURNING id, name, country_id")

	var updatedBrand datastruct.Brand
	err := qb.QueryRow().Scan(&updatedBrand.ID, &updatedBrand.Name, &updatedBrand.CountryID)
	if err != nil {
		return nil, err
	}
	return &updatedBrand, nil
}

func (b *brandQuery) DeleteBrand(id int64) error {
	qb := pgQb().
		Delete(datastruct.BrandTableName).
		From(datastruct.BrandTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
