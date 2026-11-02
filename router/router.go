package router

import (
	"github.com/go-chi/chi/v5"
)


type Router interface {
	Register(r chi.Router)
}

func SetupRouter(router Router) *chi.Mux {
	chirouter := chi.NewRouter()
    router.Register(chirouter)
	return chirouter
}