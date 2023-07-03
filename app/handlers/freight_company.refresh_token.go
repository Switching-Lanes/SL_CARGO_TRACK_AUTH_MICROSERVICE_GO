package handlers

import (
	"net/http"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/database"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// Обработчик для обновления токена доступа
func FreightCompanyRefreshTokenHandler(c *gin.Context) {
	refreshToken := c.GetHeader("Authorization")
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Отсутствует токен обновления"})
		return
	}

	token, err := utils.ValidateFreightCompanyRefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный токен обновления"})
		return
	}
	tokenString := token.Raw

	freightCompany, err := database.GetFreightCompanyByRefreshToken(tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Компания не найдена"})
		return
	}

	accessToken, err := utils.GenerateAccessToken(freightCompany.FreightCompanyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать новый токен доступа"})
		return
	}

	freightCompany.RefreshToken = accessToken
	err = database.UpdateFreightCompany(freightCompany.FreightCompanyID, bson.M{"refresh_token": freightCompany.RefreshToken})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении информации о пользователе"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
