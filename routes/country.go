package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func CountryRoutes(e *echo.Group) {
	nano := repositories.NewCountryRepository(mysql.DB)
	h := handlers.NewCountryHandler(nano)

	// e.GET("/countries", h.GetCountry)
	e.GET("/countries", middleware.Auth(h.FindCountry))
	e.GET("/country/:id", middleware.Auth(h.GetCountry))
	e.POST("/country", middleware.Auth(h.CreateCountry))
	e.DELETE("/country/:id", middleware.Auth(h.DeleteCountry))
	e.PATCH("/country/:id", middleware.Auth(h.UpdateCountry))

}
