package usersdto

import "dumbmerch/models"

type CreateUserRequest struct {
	Name        string                     `json:"name" form:"name"`
	Email       string                     `json:"email" form:"email"`
	Password    string                     `json:"password" form:"password"`
	Phone       string                     `json:"phone" form:"phone"`
	Address     string                     `json:"address" form:"address"`
	Transaction models.TransactionResponse `json:"transaction" form:"country" validate:"required"`
	// Transaction int `json:"transaction" form:"transaction" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Address  string `json:"address" form:"address"`
	// Transaction int    `json:"transaction" form:"transaction" validate:"required"`
}
