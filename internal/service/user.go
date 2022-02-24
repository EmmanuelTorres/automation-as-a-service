package service

import (
	"errors"
	"log"

	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
	"github.com/automation-as-a-service/internal/repository"
)

type UserService interface {
	GetUsers() ([]datastruct.Person, error)
	GetUser(requestedUserID int64, user *datastruct.Person) (*datastruct.Person, error)
	GetUserSingle(userID int64) (*datastruct.Person, error)
	DeleteUser(id int64, user *datastruct.Person) error
	UpdateUser(person dto.Person) (*datastruct.Person, error)
}

type userService struct {
	dao repository.DAO
}

func NewUserService(dao repository.DAO) UserService {
	return &userService{dao: dao}
}

func (u *userService) GetUsers() ([]datastruct.Person, error) {
	users, err := u.dao.NewUserQuery().GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userService) GetUser(requestedUserID int64, user *datastruct.Person) (*datastruct.Person, error) {
	requestedUser, err := u.dao.NewUserQuery().GetUser(requestedUserID)
	if err != nil {
		return nil, err
	}

	if user.ID == requestedUser.ID || user.Role == datastruct.ADMIN {
		return requestedUser, nil
	} else {
		return &datastruct.Person{ID: requestedUser.ID, FirstName: requestedUser.FirstName, LastName: requestedUser.LastName}, nil
	}
}

func (u *userService) GetUserSingle(userID int64) (*datastruct.Person, error) {
	user, err := u.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		log.Printf("User does not exist")
		return nil, err
	}
	return &datastruct.Person{ID: user.ID, Role: user.Role}, nil
}

func (u *userService) DeleteUser(id int64, user *datastruct.Person) error {
	if user.ID == id || user.Role == datastruct.ADMIN {
		err := u.dao.NewUserQuery().DeleteUser(id)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("you don't have access")
}

func (u *userService) UpdateUser(person dto.Person) (*datastruct.Person, error) {
	user, err := u.dao.NewUserQuery().GetUser(person.ID)
	if err != nil {
		return nil, err
	}

	if user.Role == datastruct.ADMIN || user.ID == person.ID {
		updatedUser, err := u.dao.NewUserQuery().UpdateUser(person)
		if err != nil {
			return nil, err
		}
		return updatedUser, nil
	}
	return nil, errors.New("you don't have access")
}
