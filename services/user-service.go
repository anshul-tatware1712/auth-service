package services

import (
	db "auth-service/db/repository"
	"fmt"
)

type UserService interface {
	CreateUser() error
}
type UserService1 struct{
	UserRepository db.UserRepository
}

func (s *UserService1) CreateUser() error {
	fmt.Println("creating user")
	s.UserRepository.CreateUser()
	return nil
}

func NewUserService(_userRepo db.UserRepository) UserService {
	return &UserService1{UserRepository: _userRepo}
}

