package authdto

type LoginResponse struct {
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	// Password string `gorm:"type: varchar(255)" json:"password"`
	Role    string `gorm:"type: varchar(255)" json:"role"`
	Token   string `gorm:"type: varchar(255)" json:"token"`
	Phone   string `gorm:"type: varchar(255)" json:"phone"`
	Address string `gorm:"type: varchar(255)" json:"address"`
}

type RegisterResponse struct {
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	// Password string `gorm:"type: varchar(255)" json:"password"`
	Role    string `gorm:"type: varchar(255)" json:"role"`
	Token   string `gorm:"type: varchar(255)" json:"token"`
	Phone   string `gorm:"type: varchar(255)" json:"phone"`
	Address string `gorm:"type: varchar(255)" json:"address"`
}
