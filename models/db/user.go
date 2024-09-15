package db

import (
	"time"

	"github.com/dita-daystaruni/verisafe/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStore struct {
	DB *gorm.DB
}

func (us *UserStore) RegisterUser(user models.User) (*models.User, error) {
	// validate the user
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	user.DateCreated = time.Now()
	user.DateUpdated = time.Now()

	// Create the user
	if err := us.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, err
}

// Retrieve a user specified by [id]
func (us *UserStore) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	if err := us.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Retrieve a user specified by [username]
func (us *UserStore) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := us.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Retrieves all users stored
func (us *UserStore) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	if err := us.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (us *UserStore) UpdateUser(user models.User) (*models.User, error) {
	stored := models.User{}
	stored.FirstName = user.FirstName
	stored.LastName = user.LastName
	stored.Active = user.Active
	stored.Username = user.Username
	stored.Password = user.Password

	err := us.DB.Debug().Model(&models.User{}).Where("id = ?", user.ID).Updates(&stored).Error
	if err != nil {
		return nil, err
	}

	return &stored, nil
}

// Deletes a user specified by [id] from the database
func (us *UserStore) DeleteStudent(id uuid.UUID) error {
	if err := us.DB.Where("student_id = ?", id).Delete(&models.RewardTransaction{}).Error; err != nil {
		return err
	}

	if err := us.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
