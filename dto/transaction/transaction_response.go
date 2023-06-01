package transactiondto

type TransactionResponse struct {
	ID         int    `json:"id"`
	CounterQty int    `json:"counterqty" form:"counterqty" validate:"required"`
	Total      int    `json:"total" form:"total" validate:"required"`
	Status     string `json:"status" form:"status" validate:"required"`
	Attachment string `json:"attachment" form:"attachment" validate:"required"`
	TripId     int    `json:"tripid" form:"tripid" validate:"required"`
}
