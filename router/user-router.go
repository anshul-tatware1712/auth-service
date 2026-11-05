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
	r.Post("/signup", ur.UserController.RegisterUser)
	r.Post("/login", ur.UserController.LoginUser)
}