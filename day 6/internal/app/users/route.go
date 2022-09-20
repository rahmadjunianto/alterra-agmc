package users

import (
	"day6/internal/dto"
	"day6/internal/middleware"
	"day6/internal/pkg/util"

	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	//id := c.Param("id")
	//route := []middleware.SkipRoute{
	//	{Route: "/api/v1/users/login", Method: "POST"},
	//	{Route: "/api/v1/users", Method: "POST"},
	//	//{Route: "/v1/books", Method: "GET"},
	//	//{Route: "/v1/books/" + string(id), Method: "GET"},
	//}
	g.Use(middleware.JWTMiddleware(dto.JWTClaims{}, util.JWT_SECRET))
	g.GET("", h.Get)
	g.POST("/login", h.Login)
	//g.GET("/:id", h.GetById)
	//g.PUT("/:id", h.UpdateById)
	//g.DELETE("/:id", h.DeleteById)
	//g.POST("", h.Create)
}
