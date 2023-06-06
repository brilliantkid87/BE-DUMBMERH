package usersdto

import "dumbmerch/models"

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" `
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
	// CountryID int            `json:"country_id" form:"country_id"`
	// Country   models.Country `json:"country"`
	TransactionID int                `json:"transaction_id"`
	Transaction   models.Transaction `json:"transaction"`
}

type DeleteUserResponse struct {
	ID int `json:"id"`
}
