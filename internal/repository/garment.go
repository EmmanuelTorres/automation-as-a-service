package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
)

type GarmentQuery interface {
	CreateGarment(garment datastruct.Garment) (*int64, error)
	GetGarment(id int64) (*datastruct.Garment, error)
	GetGarmentByCode(code string) (*datastruct.Garment, error)
	UpdateGarment(garment dto.Garment) (*datastruct.Garment, error)
	DeleteGarment(id int64) error
}

type garmentQuery struct{}

func (d *garmentQuery) CreateGarment(garment datastruct.Garment) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.GarmentTableName).
		Columns("code", "designer_id", "brand_id").
		Values(garment.Code, garment.DesignerID, garment.BrandID).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (d *garmentQuery) GetGarment(id int64) (*datastruct.Garment, error) {
	qb := pgQb().
		Select("id", "code", "designer_id", "brand_id").
		From(datastruct.GarmentTableName).
		Where(squirrel.Eq{"id": id})

	garment := datastruct.Garment{}
	err := qb.QueryRow().Scan(&garment.ID, &garment.Code, &garment.DesignerID, &garment.BrandID)
	if err != nil {
		return nil, err
	}
	return &garment, err
}

func (d *garmentQuery) GetGarmentByCode(code string) (*datastruct.Garment, error) {
	qb := pgQb().
		Select("id", "code", "designer_id", "brand_id").
		From(datastruct.GarmentTableName).
		Where(squirrel.Eq{"code": code})

	garment := datastruct.Garment{}
	err := qb.QueryRow().Scan(&garment.ID, &garment.Code, &garment.DesignerID, &garment.BrandID)
	if err != nil {
		return nil, err
	}
	return &garment, err
}

func (d *garmentQuery) UpdateGarment(garment dto.Garment) (*datastruct.Garment, error) {
	qb := pgQb().
		Update(datastruct.GarmentTableName).
		SetMap(map[string]interface{}{
			"code":        garment.Code,
			"designer_id": garment.DesignerID,
			"brand_id":    garment.BrandID,
		}).
		Where(squirrel.Eq{"id": garment.ID}).
		Suffix("RETURNING id, name, country_id")

	var updatedGarment datastruct.Garment
	err := qb.QueryRow().Scan(&updatedGarment.ID, &updatedGarment.Code, &updatedGarment.DesignerID, &garment.BrandID)
	if err != nil {
		return nil, err
	}
	return &updatedGarment, nil
}

func (d *garmentQuery) DeleteGarment(id int64) error {
	qb := pgQb().
		Delete(datastruct.GarmentTableName).
		From(datastruct.GarmentTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
