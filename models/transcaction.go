package models

type Transaction struct {
	ID         int           `json:"id"`
	UserID     int           `json:"user_id" gorm:"type:int"`
	User       UserResponse  `json:"User" gorm:"foreignkey:UserID"`
	CounterQty int           `json:"counterqty" gorm:"type:int"`
	Total      int           `json:"total" gorm:"type:int"`
	Status     string        `json:"status" gorm:"type: varchar(255)"`
	Attachment string        `json:"attachment" gorm:"type: varchar(255)"`
	TripId     int           `json:"trip_id" gorm:"type:int"`
	Trip       TripsResponse `json:"trips" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// CountryID CountryResponse
}

type TransactionResponse struct {
	ID         int    `json:"id"`
	CounterQty int    `json:"counterqty"`
	Total      int    `json:"total" `
	Status     string `json:"status"`
	Attachment string `json:"attachment"`
	TripId     int    `json:"trip_id"`
	UserID     int    `json:"user_id"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
