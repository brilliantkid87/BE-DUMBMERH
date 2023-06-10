package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

// kontrak
type TripRepository interface {
	CreateTrip(trip models.Trip) (models.Trip, error)
	GetTripByID(id int) (models.Trip, error)
	FindTrip() ([]models.Trip, error)
	UpdateTrip(trip models.Trip) (models.Trip, error)
	DeleteTrip(trip models.Trip, ID int) (models.Trip, error)
	GetCountryByID(id int) (models.CountryResponse, error)
}

// connection
func NewTripRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// method array
func (r *repository) FindTrip() ([]models.Trip, error) {
	var trips []models.Trip
	// if err := r.db.Find(&trips).Error; err != nil {
	// 	return nil, err
	// }
	err := r.db.Preload("Country").Find(&trips).Error

	return trips, err
}

func (r *repository) GetTripByID(id int) (models.Trip, error) {
	var trip models.Trip
	err := r.db.Preload("Country").First(&trip, id).Error
	// err := r.db.First(&trip, id).Error

	return trip, err

}

func (r *repository) CreateTrip(trip models.Trip) (models.Trip, error) {
	err := r.db.Preload("Country").Create(&trip).Error

	return trip, err
}

func (r *repository) UpdateTrip(trip models.Trip) (models.Trip, error) {
	// err := r.db.Save(&trip).Error
	err := r.db.Preload("Country").Save(&trip).Error
	return trip, err
}

func (r *repository) DeleteTrip(trip models.Trip, ID int) (models.Trip, error) {
	err := r.db.Delete(&trip).Error

	return trip, err
}

func (r *repository) GetCountryByID(id int) (models.CountryResponse, error) {
	var country models.CountryResponse
	err := r.db.First(&country, id).Error
	// err := r.db.First(&trip, id).Error

	return country, err

}
