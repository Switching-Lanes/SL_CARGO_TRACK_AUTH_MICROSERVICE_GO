package handlers

import (
	"net/http"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/database"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/models"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EmployeeRegisterHandler(c *gin.Context) {
	var input models.FreightCompanyEmployees

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	ID := primitive.NewObjectID().Hex()
	input.EmployeeID = ID

	input.Permissions = "driver"

	if input.Password == "" || input.Email == "" || input.ContactNumber == "" || input.Experience == "" || input.Name == "" || input.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не указаны все обязательные поля"})
		return
	}

	accessToken, err := utils.GenerateAccessToken(input.EmployeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(input.EmployeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	err = database.RegisterFreightCompanyEmployee(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register employee"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
