package main

import (
	"day2/config"
	"day2/routes"
)

func main() {
	config.InitDB()
	e := routes.Users()
	e.Logger.Fatal(e.Start(":8000"))
}
