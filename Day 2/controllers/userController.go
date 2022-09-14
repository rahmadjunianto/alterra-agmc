package controllers

import (
	"day2/lib/database"
	"day2/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetUser(c echo.Context) error {
	users, e := database.GetUsers()
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{"success": false, "message": e.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": users})
}
func GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users, e := database.GetUserById(&id)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{"success": false, "message": e.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": users})
}
func CreateUser(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)
	users, e := database.CreateUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{"success": false, "message": e.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": users})
}
func UpdateUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.Users{}
	c.Bind(&user)
	users, e := database.UpdateUserById(&id, &user)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{"success": false, "message": e.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": users})
}

func DeleteUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	e := database.DeleteUserById(&id)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{"success": false, "message": e.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
}
