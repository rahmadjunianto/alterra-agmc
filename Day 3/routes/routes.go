package routes

import (
	"day3/controllers"
	"day3/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

type skipEndPoint struct {
	Path   string
	Method string
}

func Users() *echo.Echo {
	e := echo.New()
	//logger middleware
	middlewares.LogMiddleware(e)
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("SECRET_JWT")),
		Skipper: func(c echo.Context) bool {
			// Skip middleware if path is equal 'login'
			id := c.Param("id")
			path := []skipEndPoint{
				{Path: "/v1/login", Method: "POST"},
				{Path: "/v1/users", Method: "POST"},
				{Path: "/v1/books", Method: "GET"},
				{Path: "/v1/books/" + string(id), Method: "GET"},
			}
			for _, p := range path {
				if c.Request().URL.Path == p.Path && c.Request().Method == p.Method {
					return true
				}
			}
			return false
		},
	}))

	v1 := e.Group("/v1")
	v1.POST("/login", controllers.Login)
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
