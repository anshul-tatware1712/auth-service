package services

import (
	"auth-service/db/models"
	db "auth-service/db/repository"
	"fmt"

	"github.com/golang-jwt/jwt/v5"

	"auth-service/utils"
)

type UserService interface {
	CreateUser(email, password string) *models.User
	LoginUser( email, password string) (string, error)
}
type UserService1 struct{
	UserRepository db.UserRepository
}
func NewUserService(_userRepo db.UserRepository) UserService {
	return &UserService1{UserRepository: _userRepo}
}


func (s *UserService1) CreateUser(email string, password string) *models.User {
	hashed, _ := utils.HashPassword(password)
	u := s.UserRepository.CreateUser(email, hashed)
	return u
}
func (s *UserService1) LoginUser(email string, password string) (string, error) {
	u := s.UserRepository.FindByEmail(email)
	if u == nil {
		fmt.Println("user not found")
		return "", fmt.Errorf("User not found")
	}
	if !utils.CheckPassword(u.Password, password) {
		return "", fmt.Errorf("Wrong password")
	}
	
	payload := jwt.MapClaims{"user": u.Id}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, _ := token.SignedString([]byte("secret"))

	return tokenString, nil
}


