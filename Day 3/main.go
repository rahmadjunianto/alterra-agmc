package main

import (
	"day3/config"
	m "day3/middlewares"
	"day3/routes"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func main() {

	var err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.InitDB()
	e := routes.Users()
	e.Validator = &CustomValidator{validator: validator.New()}
	//implement middleware logger
	m.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":8000"))
}
