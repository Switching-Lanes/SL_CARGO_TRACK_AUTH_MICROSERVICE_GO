package handlers

import (
	"net/http"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/database"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/models"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/utils"
	"github.com/gin-gonic/gin"
)

// Обработчик для обновления токена доступа
func UserRefreshTokenHandler(c *gin.Context) {
	refreshToken := c.GetHeader("Authorization")
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Отсутствует токен обновления"})
		return
	}

	token, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный токен обновления"})
		return
	}
	tokenString := token.Raw

	user, err := database.GetUserByRefreshToken(tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Компания не найдена"})
		return
	}

	accessToken, err := utils.GenerateAccessToken(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать новый токен доступа"})
		return
	}

	update := models.User{
		RefreshToken: accessToken,
	}

	user.RefreshToken = accessToken
	err = database.UpdateUser(user.UserID, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении информации о пользователе"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
