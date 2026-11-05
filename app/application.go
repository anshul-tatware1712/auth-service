package app

import (
	"auth-service/config"
	db "auth-service/config/db"
	"auth-service/controller"
	repo "auth-service/db/repository"
	"auth-service/router"
	"auth-service/services"
	"fmt"
	"net/http"
	"time"
)

type Application struct {
	Config Config
	Store repo.Storage
}

type Config struct {
	Address string
}

func NewConfig() *Config {
	port:= config.GetString("PORT")

	return &Config{
		Address: port,
	}
}

func NewApplication(cfg Config) *Application {
	return &Application{Config: cfg,Store: *repo.NewStorage()}
}

func (app *Application) Run() error {

	db, _ := db.CreateDB()
	ur := repo.NewUserRepo(db)
	us:=services.NewUserService(ur)
	uc := controller.NewUserController(us)
	_ur:=router.NewUserRouter(uc)

	server := http.Server{
		Addr: app.Config.Address,
		Handler: router.SetupRouter(_ur),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout: 5 * time.Second,
	}
	fmt.Println("Starting server on", app.Config.Address)
	return server.ListenAndServe()
}