package main

import (
	"day4/config"
	m "day4/middlewares"
	"day4/routes"
	"day4/util"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	var err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitDB()
	e := routes.Users()
	e.Validator = &util.CustomValidator{Validator: validator.New()}
	//implement middleware logger
	m.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":8000"))
}
