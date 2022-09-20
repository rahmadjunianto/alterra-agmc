package main

import (
	"day6/database"
	"day6/internal/factory"
	"day6/internal/http"
	"day6/internal/middleware"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	database.GetConnection()
}
func main() {

	f := factory.NewFactory()
	e := echo.New()

	middleware.LogMiddlewares(e)

	http.NewHttp(e, f)

	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("routes.json", data, 0644)
	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
