package middlewares

import (
	"day4/config"
	"day4/models"
	"github.com/labstack/echo/v4"
)

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.Users
	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
