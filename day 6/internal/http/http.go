package http

import (
	"day6/internal/app/auth"
	"day6/internal/app/users"
	"day6/internal/factory"
	"day6/pkg/util"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.GET("/status", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})
	v1 := e.Group("/api/v1")
	users.NewHandler(f).Route(v1.Group("/users"))
	auth.NewHandler(f).Route(v1.Group("/auth"))

}
