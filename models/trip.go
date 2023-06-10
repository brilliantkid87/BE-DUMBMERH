package models

import "time"

type Trip struct {
	ID             int             `json:"id"`
	Title          string          `json:"title" gorm:"type:varchar(255)"`
	Accommodation  string          `json:"accommodation" gorm:"type:varchar(255)"`
	Transportation string          `json:"transportation" gorm:"type:varchar(255)"`
	Eat            string          `json:"eat" gorm:"type:varchar(255)"`
	Day            int             `json:"day"`
	Night          int             `json:"night"`
	DateTrip       string          `json:"date_trip" gorm:"type:varchar(255)"`
	Price          int             `json:"price"`
	Quota          int             `json:"quota"`
	Description    string          `json:"description" gorm:"type:varchar(255)"`
	Image          string          `json:"image" gorm:"type:varchar(255)"`
	CountryID      int             `json:"country_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Country        CountryResponse `json:"country" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	// Country   Country   `json:"country" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
	UserID    int          `json:"user_id"`
	User      UserResponse `json:"user"`
}

type TripsResponse struct {
	ID             int             `json:"id"`
	Title          string          `json:"title"`
	CountryID      int             `json:"country_id"`
	Country        CountryResponse `json:"country"`
	Accommodation  string          `json:"accommodation"`
	Transportation string          `json:"transportation"`
	Eat            string          `json:"eat"`
	Day            int             `json:"day"`
	Night          int             `json:"night"`
	DateTrip       string          `json:"date_trip"`
	Price          int             `json:"price"`
	Quota          int             `json:"quota"`
	Description    string          `json:"description"`
	Image          string          `json:"image"`
}

func (TripsResponse) TableName() string {
	return "trips"
}
