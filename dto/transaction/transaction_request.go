package transactiondto

type CreateTransactionRequest struct {
	CounterQty int    `json:"counterqty" form:"counterqty" validate:"required"`
	Total      int    `json:"total" form:"total" validate:"required"`
	Status     string `json:"status" form:"status" validate:"required"`
	Attachment string `json:"attachment" form:"attachment" validate:"required"`
	TripId     int    `json:"tripid" form:"tripid" validate:"required"`
	UserID     int    `json:"userid" form:"userid" validate:"required"`
}

type UpdateTransactionRequest struct {
	CounterQty int    `json:"counterqty" form:"counterqty" validate:"required"`
	Total      int    `json:"total" form:"total" validate:"required"`
	Status     string `Json:"status" form:"status" validate:"required"`
	Attachment string `json:"attachment" form:"attachment" validate:"required"`
	TripId     int    `json:"tripid" form:"tripid" validate:"required"`
}
