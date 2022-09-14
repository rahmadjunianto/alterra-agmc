package controllers

import (
	"day2/lib/database"
	"day2/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func GetBook(c echo.Context) error {
	books := database.GetBooks()
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": books})
}
func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	books := database.GetBooksById(id)
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": books})
}
func CreateBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	b := database.CreateBook(book)
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": b})
}
func UpdateBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{}
	c.Bind(&book)
	b := database.UpdateBookById(id, book)
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": b})
}

func DeleteBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	b := database.DeleteBooksById(id)
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": b})
}
