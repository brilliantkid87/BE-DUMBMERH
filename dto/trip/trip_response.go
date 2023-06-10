package tripdto

import "dumbmerch/models"

type TripResponse struct {
	ID             int            `json:"id"`
	Title          string         `json:"title" form:"title"`
	CountryID      int            `json:"country_id" form:"country_id"`
	Country        models.Country `json:"country"`
	Accommodation  string         `json:"accommodation" form:"accommodation"`
	Transportation string         `json:"transportation" form:"transportatiob"`
	Eat            string         `json:"eat" form:"eat"`
	Day            int            `json:"day" form:"day"`
	Night          int            `json:"night" form:"night"`
	DateTrip       string         `json:"date_trip" form:"date_trip"`
	Price          int            `json:"price" form:"price"`
	Quota          int            `json:"quota" form:"quota"`
	Description    string         `json:"description" form:"description"`
	Image          string         `json:"image" form:"image"`
}
