package transactiondto

import "dumbmerch/models"

type CreateTransactionRequest struct {
	CounterQty int                  `json:"counterqty" form:"counterqty" `
	Total      int                  `json:"total" form:"total" `
	Status     string               `json:"status" form:"status" `
	Attachment string               `json:"attachment" form:"attachment" `
	TripId     int                  `json:"trip_id" form:"trip_id"`
	Trip       models.TripsResponse `json:"trip"`
	UserID     int                  `json:"user_id" form:"user_id"`
	User       models.User          `json:"user"`
}

type UpdateTransactionRequest struct {
	CounterQty int    `json:"counterqty" form:"counterqty" validate:"required"`
	Total      int    `json:"total" form:"total" validate:"required"`
	Status     string `Json:"status" form:"status" validate:"required"`
	Attachment string `json:"attachment" form:"attachment" validate:"required"`
	TripId     int    `json:"trip_id" form:"trip_id" validate:"required`
}
