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
	e.GET("/countries", h.FindCountry)
	e.GET("/country/:id", h.GetCountry)
	e.POST("/country", h.CreateCountry)
	e.DELETE("/country/:id", middleware.Auth(h.DeleteCountry))
	e.PATCH("/country/:id", h.UpdateCountry)

}
