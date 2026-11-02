package main

import (
	"fmt"

	"auth-service/app"
	"auth-service/config"
)


func main() {
	fmt.Println("Hello")
	config.Load()

	cfg := app.NewConfig()
	app := app.NewApplication(*cfg)
	 
	app.Run()
}