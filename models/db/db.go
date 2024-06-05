package db

import (
	"github.com/dita-daystaruni/verisafe/models"
	"gorm.io/gorm"
)

func AutoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.Student{})
}
