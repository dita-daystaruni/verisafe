package db

import (
	"github.com/dita-daystaruni/verisafe/models"
	"gorm.io/gorm"
)

func AutoMigrate(DB *gorm.DB) {
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	DB.AutoMigrate(&models.Student{}, &models.RewardTransaction{}, &models.User{}, &models.Token{})
}
