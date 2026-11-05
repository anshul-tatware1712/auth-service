package controller

import (
	"auth-service/services"
	"net/http"
)

type UserController struct{
	UserService services.UserService
}

func NewUserController(us services.UserService) *UserController {	
	return &UserController{UserService: us}
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	uc.UserService.CreateUser("name", "email", "password")
	w.Write([]byte("user created"))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	uc.UserService.LoginUser("email", "password")
	w.Write([]byte("user logged in"))
}