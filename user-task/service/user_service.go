package service

import (
	"fmt"
	"usermgmt/models"
)

type UserService struct {
	users []models.User
}

func NewUserService() *UserService {
	return &UserService{users: []models.User{}}
}

func (s *UserService) AddUser(user models.User) error {
	for _, u := range s.users {
		if u.ID == user.ID {
			return fmt.Errorf("user with ID %d already exists", user.ID)
		}
	}
	s.users = append(s.users, user)
	return nil
}

func (s *UserService) GetUserByID(id models.UserID) (models.User, error) {
	for _, user := range s.users {
		if user.ID == id {
			return user, nil
		}
	}

	return models.User{}, fmt.Errorf("user with ID %d not found", id)
}

func (s *UserService) UpdateUserName(id models.UserID, newName string) error {
	for i := range s.users {
		if s.users[i].ID == id {
			s.users[i].Name = newName
			return nil
		}
	}

	return fmt.Errorf("Could not find user with ID %v to update", id)
}

func (s *UserService) DeleteUser(id models.UserID) error {
	for i := range s.users {
		if s.users[i].ID == id {
			s.users = append(s.users[:i], s.users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID %v not found", id)
}

func (s UserService) PrintAllUsers() {
	for _, user := range s.users {
		fmt.Printf("ID: %v, Name: %v, Age: %v, isAdult: %v\n", user.ID, user.Name, user.Age, user.IsAdult())
	}
}
