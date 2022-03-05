package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
)

type DesignerQuery interface {
	CreateDesigner(designer datastruct.Designer) (*int64, error)
	GetDesigner(id int64) (*datastruct.Designer, error)
	GetDesignerByName(name string) (*datastruct.Designer, error)
	UpdateDesigner(designer dto.Designer) (*datastruct.Designer, error)
	DeleteDesigner(id int64) error
}

type designerQuery struct{}

func (d *designerQuery) CreateDesigner(designer datastruct.Designer) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.DesignerTableName).
		Columns("name", "country_id").
		Values(designer.Name, designer.CountryID).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (d *designerQuery) GetDesigner(id int64) (*datastruct.Designer, error) {
	qb := pgQb().
		Select("id", "name", "country_id").
		From(datastruct.DesignerTableName).
		Where(squirrel.Eq{"id": id})

	designer := datastruct.Designer{}
	err := qb.QueryRow().Scan(&designer.ID, &designer.Name, &designer.CountryID)
	if err != nil {
		return nil, err
	}
	return &designer, err
}

func (d *designerQuery) GetDesignerByName(name string) (*datastruct.Designer, error) {
	qb := pgQb().
		Select("id", "name", "country_id").
		From(datastruct.DesignerTableName).
		Where(squirrel.Eq{"name": name})

	designer := datastruct.Designer{}
	err := qb.QueryRow().Scan(&designer.ID, &designer.Name, &designer.CountryID)
	if err != nil {
		return nil, err
	}
	return &designer, err
}

func (d *designerQuery) UpdateDesigner(designer dto.Designer) (*datastruct.Designer, error) {
	qb := pgQb().
		Update(datastruct.DesignerTableName).
		SetMap(map[string]interface{}{
			"name":       designer.Name,
			"country_id": designer.CountryID,
		}).
		Where(squirrel.Eq{"id": designer.ID}).
		Suffix("RETURNING id, name")

	var updatedDesigner datastruct.Designer
	err := qb.QueryRow().Scan(&updatedDesigner.ID, &updatedDesigner.Name, &updatedDesigner.CountryID)
	if err != nil {
		return nil, err
	}
	return &updatedDesigner, nil
}

func (d *designerQuery) DeleteDesigner(id int64) error {
	qb := pgQb().
		Delete(datastruct.DesignerTableName).
		From(datastruct.DesignerTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
