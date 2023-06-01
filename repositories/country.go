package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

// kontrak
type CountryRepository interface {
	FindCountry() ([]models.Country, error)
	GetCountry(id int) (models.Country, error)
	CreateCountry(country models.Country) (models.Country, error)
	UpdateCountry(country models.Country) (models.Country, error)
	DeleteCountry(country models.Country, ID int) (models.Country, error)
}

// connection
func NewCountryRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// array
func (r *repository) FindCountry() ([]models.Country, error) {
	var countries []models.Country
	if err := r.db.Find(&countries).Error; err != nil {
		return nil, err
	}
	return countries, nil
}

func (r *repository) GetCountry(id int) (models.Country, error) {
	var country models.Country
	// err := r.db.Preload("Trip").First(&country, id).Error
	err := r.db.First(&country, id).Error

	return country, err
}

func (r *repository) CreateCountry(country models.Country) (models.Country, error) {
	err := r.db.Create(&country).Error

	return country, err
}

func (r *repository) UpdateCountry(country models.Country) (models.Country, error) {
	err := r.db.Save(&country).Error

	return country, err
}

func (r *repository) DeleteCountry(country models.Country, ID int) (models.Country, error) {
	err := r.db.Delete(&country, ID).Scan(&country).Error

	return country, err
}
