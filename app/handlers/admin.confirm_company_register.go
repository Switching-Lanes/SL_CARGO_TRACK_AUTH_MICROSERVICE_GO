package handlers

import (
	"fmt"
	"net/http"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/database"
	"github.com/gin-gonic/gin"
)

func ConfirmRegisterFreightCompany(c *gin.Context) {
	confirmationLink := c.Param("confirmationLink")

	LINK := fmt.Sprintf("http://localhost:8080/admin/confirm-register-company/%s", confirmationLink)

	company, err := database.GetFreightCompanyByConfirmationLink(LINK)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при поиске пользователя"})
		return
	}
	if company == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Недействительный код подтверждения"})
		return
	}

	if company.IsRegistered {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Почта уже подтверждена"})
		return
	}

	err = database.UpdateFreightCompanyRegisteredStatus(company.FreightCompanyID, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении статуса подтверждения почты"})
		return
	}

	c.Redirect(http.StatusPermanentRedirect, "https://t.me/luxuryk1d")
}
