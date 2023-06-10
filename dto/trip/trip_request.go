package tripdto

import "dumbmerch/models"

type CreateTripRequest struct {
	Title          string         `json:"title" form:"title"`
	CountryID      int            `json:"country_id" form:"country_id"`
	Country        models.Country `json:"country" form:"country"`
	Accommodation  string         `json:"accommodation" form:"accommodation"`
	Transportation string         `json:"transportation" form:"transportation"`
	Eat            string         `json:"eat" form:"eat"`
	Day            int            `json:"day" form:"day"`
	Night          int            `json:"night" form:"night"`
	DateTrip       string         `json:"date_trip" form:"date_trip"`
	Price          int            `json:"price" form:"price"`
	Quota          int            `json:"quota" form:"quota"`
	Description    string         `json:"description" form:"description"`
	Image          string         `json:"image" form:"image"`
}

type UpdateTripRequest struct {
	Title          string                 `json:"title" form:"title" validate:"required"`
	CountryID      int                    `json:"country_id" form:"country_id" validate:"required"`
	Country        models.CountryResponse `json:"country" form:"country" validate:"required"`
	Accommodation  string                 `json:"accommodation" form:"accommodation" validate:"required"`
	Transportation string                 `json:"transportation" form:"transportation" validate:"required"`
	Eat            string                 `json:"eat" form:"eat" validate:"required"`
	Day            int                    `json:"day" form:"day" validate:"required"`
	Night          int                    `json:"night" form:"night" validate:"required"`
	DateTrip       string                 `json:"date_trip" form:"date_trip" validate:"required"`
	Price          int                    `json:"price" form:"price" validate:"required"`
	Quota          int                    `json:"quota" form:"quota" validate:"required"`
	Description    string                 `json:"description" form:"description" validate:"required"`
	Image          string                 `json:"image" form:"image" validate:"required"`
}
