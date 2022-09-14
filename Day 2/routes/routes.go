package routes

import (
	"day2/controllers"
	"github.com/labstack/echo/v4"
)

func Users() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")
	users := v1.Group("/users")
	books := v1.Group("/books")
	users.GET("", controllers.GetUser)
	users.GET("/:id", controllers.GetUserById)
	users.POST("", controllers.CreateUser)
	users.PUT("/:id", controllers.UpdateUserById)
	users.DELETE("/:id", controllers.DeleteUserById)

	books.GET("", controllers.GetBook)
	books.GET("/:id", controllers.GetBookById)
	books.POST("", controllers.CreateBook)
	books.PUT("/:id", controllers.UpdateBookById)
	books.DELETE("/:id", controllers.DeleteBookById)
	return e
}
