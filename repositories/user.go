package repositories

import (
	"dumbmerch/models"

	"gorm.io/gorm"
)

// kontrak
type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User, ID int) (models.User, error)
}

// connection
func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

// method array
func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Transaction").Find(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	// err := r.db.Preload("Transaction").Preload("Trip.Country").First(&user).Error
	err := r.db.Preload("Transaction").First(&user, ID).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {
	err := r.db.Delete(&user, ID).Scan(&user).Error

	return user, err
}
