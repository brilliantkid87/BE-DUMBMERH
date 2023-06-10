package transactiondto

import "dumbmerch/models"

type TransactionResponse struct {
	ID         int                  `json:"id"`
	CounterQty int                  `json:"counterqty" form:"counterqty" validate:"required"`
	Total      int                  `json:"total" form:"total" validate:"required"`
	Status     string               `json:"status" form:"status" validate:"required"`
	Attachment string               `json:"attachment" form:"attachment" validate:"required"`
	TripId     int                  `json:"trip_id" form:"trip_id" validate:"required"`
	Trip       models.TripsResponse `json:"trip"`
	UserID     int                  `json:"user_id"`
}