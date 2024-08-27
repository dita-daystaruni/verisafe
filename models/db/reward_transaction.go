package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/dita-daystaruni/verisafe/api/v1/handlers/events"
	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/dita-daystaruni/verisafe/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RewardTransactionStore struct {
	DB *gorm.DB
}

// Saves a reward transaction
func (rts *RewardTransactionStore) NewRewardTransaction(rewardTransaction models.RewardTransaction, cfg configs.Config) (*models.RewardTransaction, error) {
	rewardTransaction.AwardedAt = time.Now()

	// Retrieve the user in question
	var student models.Student
	err := rts.DB.First(&student, rewardTransaction.StudentID).Error
	if err != nil {
		return nil, errors.New("Please check user you are trying to allocate points")
	}

	// Write the transaction
	if err := rts.DB.Save(&rewardTransaction).Error; err != nil {
		return nil, err
	}

	// Add the points
	student.VibePoints += rewardTransaction.Points
	if err := rts.DB.Debug().Model(&student).Update("vibe_points", student.VibePoints).Error; err != nil {
		return nil, err
	}

	go events.EmitUserUpdated(&student, &cfg)

	return &rewardTransaction, nil
}

func (rts *RewardTransactionStore) GetUserTransactions(userID uuid.UUID) (*[]models.RewardTransaction, error) {
	var rewardTransactions []models.RewardTransaction

	err := rts.DB.Where("student_id = ? and points > 0", userID).Find(&rewardTransactions).Error
	if err != nil {
		return nil, errors.New("Please check user you are trying to allocate points")
	}

	return &rewardTransactions, nil
}

func (rts *RewardTransactionStore) GetAllTransactions() (*[]models.RewardTransaction, error) {
	var rewardTransactions []models.RewardTransaction

	err := rts.DB.Find(&rewardTransactions).Error
	if err != nil {
		return nil, err
	}

	return &rewardTransactions, nil
}

func (rts *RewardTransactionStore) DeleteTransaction(id uuid.UUID) (*models.RewardTransaction, error) {
	var rewardTransaction models.RewardTransaction

	err := rts.DB.Find(&rewardTransaction, id).Error
	if err != nil {
		return nil, err
	}

	// Retrieve the user in question
	var student models.Student
	err = rts.DB.First(&student, rewardTransaction.StudentID).Error
	if err != nil {
		return nil, errors.New("Failed to retrieve student associated with that transaction")
	}

	// Subtract the points
	student.VibePoints -= rewardTransaction.Points

	if err := rts.DB.Debug().Model(&student).Update("vibe_points", student.VibePoints).Error; err != nil {
		return nil, err
	}

	// Write the transaction
	rewardTransaction.Reason += fmt.Sprintf("(revoked) %d points", rewardTransaction.Points)
	rewardTransaction.Points = 0
	if err := rts.DB.Save(&rewardTransaction).Error; err != nil {
		return nil, err
	}

	return &rewardTransaction, nil
}
