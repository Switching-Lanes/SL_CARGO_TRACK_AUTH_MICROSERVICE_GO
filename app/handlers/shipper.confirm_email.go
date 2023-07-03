package handlers

import (
	"fmt"
	"net/http"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/database"
	"github.com/gin-gonic/gin"
)

// ConfirmEmailHandler обработчик для подтверждения электронной почты пользователя
func ConfirmEmailGetHandler(c *gin.Context) {
	// Получение кода подтверждения из параметров запроса
	confirmationLink := c.Param("confirmationLink")

	LINK := fmt.Sprintf("http://localhost:8080/shipper/confirm-email/%s", confirmationLink)

	// Проверка наличия пользователя с указанным кодом подтверждения
	user, err := database.GetShipperByConfirmationLink(LINK)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при поиске пользователя"})
		return
	}
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Недействительный код подтверждения"})
		return
	}

	// Проверка статуса подтверждения почты пользователя
	if user.EmailConfirmed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Почта уже подтверждена"})
		return
	}

	// Обновление статуса подтверждения почты пользователя
	err = database.UpdateEmailConfirmationStatus(user.ShipperID, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении статуса подтверждения почты"})
		return
	}

	c.Redirect(http.StatusPermanentRedirect, "https://t.me/MARTE11XO")
}

func ConfirmEmailHandler(c *gin.Context) {
	confirmationCode := c.PostForm("confirmationCode")

	if confirmationCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Отсутствует код подтверждения"})
		return
	}

	user, err := database.GetShipperByConfirmationCode(confirmationCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при поиске пользователя"})
		return
	}
	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Недействительный код подтверждения"})
		return
	}

	if user.EmailConfirmed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Почта уже подтверждена"})
		return
	}

	err = database.UpdateEmailConfirmationStatus(user.ShipperID, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении статуса подтверждения почты"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Почта успешно подтверждена"})
}
