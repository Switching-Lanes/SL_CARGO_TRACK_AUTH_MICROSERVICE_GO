package handlers

import (
	"net/http"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/database"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/models"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/utils"
	"github.com/gin-gonic/gin"
)

func ShipperLogin(c *gin.Context) {
	var input models.Shipper

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	if input.Password == "" || input.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не указаны все обязательные поля"})
		return
	}

	user, err := database.GetShipperByEmail(input.Email)
	if err != nil {
		// Обработать ошибку, например, отправить сообщение об ошибке клиенту
		c.JSON(http.StatusNotFound, gin.H{"error": "Неверный адрес электронной почты или пароль"})
		return
	}

	if !user.EmailConfirmed {
		c.JSON(http.StatusNotFound, gin.H{"error": "Подтвердите электронную почту"})
		return
	}

	// Проверяем пароль
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный адрес электронной почты или пароль"})
		return
	}

	accessToken, err := utils.GenerateAccessToken(user.ShipperID)
	if err != nil {
		// Обработать ошибку генерации токена
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сгенерировать токен"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ShipperID)
	if err != nil {
		// Обработать ошибку генерации токена
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сгенерировать токен"})
		return
	}

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken

	err = database.UpdateShipperTokens(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить токены пользователя"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shipper": input,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
