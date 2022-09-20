package controllers

import (
	"day4/lib/database"
	"day4/middlewares"
	"day4/models"
	"fmt"
	"github.com/go-playground/validator/v10"
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
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"success": false, "message": err.Error()})
	}
	users, e := database.CreateUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{"success": false, "message": e.Error()})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": users})
}
func UpdateUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	authorization := middlewares.CheckAuthorization(id, c)
	if !authorization {
		return echo.NewHTTPError(http.StatusUnauthorized, map[string]interface{}{"success": false, "message": "unauthorized"})
	}
	user := models.Users{}
	c.Bind(&user)
	if err := c.Validate(user); err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"success": false, "message": errorMessages})
	}
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

func Login(c echo.Context) error {
	user := models.Users{}
	c.Bind(&user)
	users, e := database.Login(user)
	if e != nil {
		if e.Error() == "record not found" {
			return echo.NewHTTPError(http.StatusBadRequest, "these credentials do not match our records")
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
		}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": users})
}
