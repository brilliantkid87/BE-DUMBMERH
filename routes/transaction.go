package routes

import (
	"dumbmerch/handlers"
	"dumbmerch/pkg/middleware"
	"dumbmerch/pkg/mysql"
	"dumbmerch/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	bravo := repositories.NewCountryRepository(mysql.DB)
	h := handlers.HandlerTransaction(bravo)

	e.GET("/transactions", h.FIndTransaction)
	e.GET("/transaction/:id", h.GetTransaction)
	e.GET("/transaction-user/:id", h.GetTransactionByUser)
	// e.PATCH("/transaction/:id", h.UpdateTransaction)
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	e.POST("/notification", h.Notification)
}
