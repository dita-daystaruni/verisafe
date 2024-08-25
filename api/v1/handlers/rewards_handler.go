package handlers

import (
	"net/http"

	"github.com/dita-daystaruni/verisafe/configs"
	"github.com/dita-daystaruni/verisafe/models"
	"github.com/dita-daystaruni/verisafe/models/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RewardsHandler struct {
	Store *db.RewardTransactionStore
	Cfg   *configs.Config
}

func (rh *RewardsHandler) NewTransaction(c *gin.Context) {
	var newTransaction models.RewardTransaction

	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Please check your request body"})
		return
	}

	// write to the db
	transaction, err := rh.Store.NewRewardTransaction(newTransaction, *rh.Cfg)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, transaction)
}

func (rh *RewardsHandler) GetUserTransactions(c *gin.Context) {
	rawID := c.Param("userid")
	id, err := uuid.Parse(rawID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to parse student id"})
		return

	}

	// write to the db
	transactions, err := rh.Store.GetUserTransactions(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, transactions)
}

func (rh *RewardsHandler) GetAllTransactions(c *gin.Context) {
	transactions, err := rh.Store.GetAllTransactions()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, transactions)
}

func (rh *RewardsHandler) DeleteRewardTransaction(c *gin.Context) {
	rawID := c.Param("transaction")
	id, err := uuid.Parse(rawID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Failed to parse transaction id"})
		return

	}

	// write to the db
	transaction, err := rh.Store.DeleteTransaction(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, transaction)
}
