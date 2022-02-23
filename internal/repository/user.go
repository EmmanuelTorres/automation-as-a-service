package repository

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
)

type UserQuery interface {
	CreateUser(user datastruct.Person) (*int64, error)
	GetUser(id int64) (*datastruct.Person, error)
	DeleteUser(userID int64) error
	UpdateUser(person dto.Person) (*datastruct.Person, error)
	GetUserPasswordByEmail(email string) (*string, error)
	GetUserIdByEmail(email string) (*int64, error)
}

type userQuery struct{}

func (u *userQuery) CreateUser(user datastruct.Person) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.PersonTableName).
		Columns("first_name", "last_name", "email", "password", "phone_number", "role").
		Values(user.FirstName, user.LastName, user.Email, user.Password, user.PhoneNumber, user.Role).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (u *userQuery) GetUser(id int64) (*datastruct.Person, error) {
	qb := pgQb().
		Select("id", "first_name", "last_name", "email", "password", "phone_number", "role").
		From(datastruct.PersonTableName).
		Where(squirrel.Eq{"id": id})

	user := datastruct.Person{}
	err := qb.QueryRow().Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.PhoneNumber, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (u *userQuery) DeleteUser(userID int64) error {
	qb := pgQb().
		Delete(datastruct.PersonTableName).
		From(datastruct.PersonTableName).
		Where(squirrel.Eq{"id": userID})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (u *userQuery) UpdateUser(person dto.Person) (*datastruct.Person, error) {
	qb := pgQb().
		Update(datastruct.PersonTableName).
		SetMap(map[string]interface{}{
			"first_name":   person.FirstName,
			"last_name":    person.LastName,
			"email":        person.Email,
			"phone_number": person.PhoneNumber,
		}).
		Where(squirrel.Eq{"id": person.ID}).
		Suffix("RETURNING id, first_name, last_name, email, phone_number")

	var updatedPerson datastruct.Person
	err := qb.QueryRow().Scan(&updatedPerson.ID, &updatedPerson.FirstName, &updatedPerson.LastName, &updatedPerson.Email, &updatedPerson.PhoneNumber)
	if err != nil {
		return nil, err
	}
	return &updatedPerson, nil
}

func (u *userQuery) GetUserPasswordByEmail(email string) (*string, error) {
	qb := pgQb().
		Select("password").
		From(datastruct.PersonTableName).
		Where(squirrel.Eq{"email": email})

	var password string
	err := qb.QueryRow().Scan(&password)
	if err != nil {
		return nil, fmt.Errorf("email and password don't match %v", err)
	}
	return &password, nil
}

func (u *userQuery) GetUserIdByEmail(email string) (*int64, error) {
	qb := pgQb().
		Select("id").
		From(datastruct.PersonTableName).
		Where(squirrel.Eq{"email": email})

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("cannot get user id %v", err)
	}
	return &id, nil
}
