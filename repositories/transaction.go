package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

// kontrak
type TransactionRepository interface {
	FIndTransaction() ([]models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransaction(id int) (models.Transaction, error)
	// UpdateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, orderId int) (models.Transaction, error)
	GetUpdateTripByID(id int) (models.TripsResponse, error)
	GetUserByID(id int) (models.UserResponse, error)
	GetCountriesByID(id int) (models.CountryResponse, error)
	GetTransactionByUser(ID int) ([]models.Transaction, error)
}

// connection
func NewTransactionRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// method array
func (r *repository) FIndTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Preload("Trip.Country").Find(&transactions).Error

	return transactions, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("Trip.Country").Create(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransaction(id int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Preload("Trip.Country").First(&transaction, id).Error

	return transaction, err
}

// func (r *repository) UpdateTransaction(transaction models.Transaction) (models.Transaction, error) {
// 	err := r.db.Delete(&transaction).Error

// 	return transaction, err
// }

func (r *repository) GetUpdateTripByID(id int) (models.TripsResponse, error) {
	var trip models.TripsResponse
	err := r.db.Preload("Country").First(&trip, id).Error
	// err := r.db.First(&trip, id).Error

	return trip, err

}

func (r *repository) GetUserByID(id int) (models.UserResponse, error) {
	var trip models.UserResponse
	err := r.db.Preload("User").First(&trip, id).Error
	// err := r.db.First(&trip, id).Error

	return trip, err

}

func (r *repository) GetCountriesByID(id int) (models.CountryResponse, error) {
	var trip models.CountryResponse
	err := r.db.First(&trip, id).Error
	// err := r.db.First(&trip, id).Error

	return trip, err

}

func (r *repository) GetTransactionByUser(ID int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := r.db.Where("user_id =?", ID).Preload("User").Preload("Trip.Country").Find(&transaction).Error

	return transaction, err
}

func (r *repository) UpdateTransaction(status string, orderId int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("Trip.Country").First(&transaction, orderId)

	if status != transaction.Status && status == "success" {
		var trip models.Trip
		r.db.First(&trip, transaction.Trip.ID)
		trip.Quota = trip.Quota - 1
		r.db.Save(&trip)
	}

	transaction.Status = status
	err := r.db.Save(&transaction).Error
	return transaction, err
}
