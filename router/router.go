package router

import (
	"github.com/go-chi/chi/v5"
)


type Router interface {
	Register(r chi.Router)
}

func SetupRouter(userRouter Router) *chi.Mux {

	chirouter := chi.NewRouter()
 
    userRouter.Register(chirouter)
	return chirouter
}