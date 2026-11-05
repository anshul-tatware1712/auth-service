package controller

import (
	"auth-service/dto"
	"auth-service/services"
	"auth-service/utils"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(us services.UserService) *UserController {
	return &UserController{UserService: us}
}

func (uc *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	req := &dto.AuthRequest{}
	if err := utils.Validator.Struct(req); err != nil {
		utils.WriteJsonToErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	user := uc.UserService.CreateUser(req.Email, req.Password)
	utils.WriteJsonToSuccessResponse(w, user)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	req := &dto.AuthRequest{}
	if err := utils.Validator.Struct(req); err != nil {
		utils.WriteJsonToErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	accessToken, err := uc.UserService.LoginUser(req.Email, req.Password)
	if err != nil {
		utils.WriteJsonToErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}
	utils.WriteJsonToSuccessResponse(w, accessToken)
}
