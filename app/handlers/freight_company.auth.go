package handlers

import (
	"net/http"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/database"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/models"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/utils"
	"github.com/gin-gonic/gin"
)

func FreightCompanyLogin(c *gin.Context) {
	var input models.FreightCompany

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return 
	}

	if input.Password == "" || input.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не указаны все обязательные поля"})
		return
	}

	company, err := database.GetFreightCompanyByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Неверный адрес электронной почты или пароль"})
		return 
	}

	if !company.IsRegistered {
		c.JSON(http.StatusNotFound, gin.H{"error": "Компания еще не зарегистрирована, ожидайте подтверждения"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, company.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный адрес электронной почты или пароль"})
		return 
	}

	accessToken, err := utils.GenerateAccessToken(company.FreightCompanyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сгенерировать токен"})
		return 
	}

	refreshToken, err := utils.GenerateRefreshToken(company.FreightCompanyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сгенерировать токен"})
		return
	}

	company.AccessToken = accessToken
	company.RefreshToken = refreshToken

	err = database.UpdateFreightCompanyTokens(company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить токены пользователя"})
		return
	}
	c.JSON(http.StatusOK, company)
}
