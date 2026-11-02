package services

import (
	db "auth-service/db/repository"
	"fmt"
)

type UserService interface {
	CreateUser() error
}
type UserServiceImpl struct{
	UserRepository db.UserRepository
}

func (s *UserServiceImpl) CreateUser() error {
	fmt.Println("creating user")
	s.UserRepository.CreateUser()
	return nil

}

func NewUserService(_userRepo db.UserRepository) UserService {
	return &UserServiceImpl{UserRepository: _userRepo}
}

