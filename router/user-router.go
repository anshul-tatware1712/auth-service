package router

import (
	"auth-service/controller"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct{
	UserController *controller.UserController
}

func NewUserRouter(uc *controller.UserController) Router {
	return &UserRouter{UserController: uc}
}

func (ur *UserRouter) Register(r chi.Router) {
	r.Post("/user", ur.UserController.RegisterUser)
}