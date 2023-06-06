package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Role     string `json:"role" gorm:"type: varchar(255)"`
	// Transaction TransactionResponse `json:"transaction" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TransactionID int                   `json:"transaction_id"`
	Transaction   []TransactionResponse `json:"transaction" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func (UserResponse) TableName() string {
	return "users"
}
