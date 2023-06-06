package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func TripRoutes(e *echo.Group) {
	delta := repositories.NewTripRepository(mysql.DB)
	h := handlers.NewTripHandler(delta)

	e.GET("/trips", h.FindTrip)
	e.GET("/trip/:id", h.GetTripByID)
	e.POST("/trip", middleware.Auth(middleware.UploadFile(h.CreateTrip)))
	// e.POST("/trip", h.CreateTrip)
	e.DELETE("/trip/:id", middleware.Auth(h.DeleteTrip))
	e.PATCH("/trip/:id", middleware.Auth(middleware.UploadFile(h.UpdateTrip)))
}
