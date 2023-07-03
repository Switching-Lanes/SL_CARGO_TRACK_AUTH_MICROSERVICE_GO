package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/database"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/models"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FreightCompanyRegisterHandler(c *gin.Context) {
	var input models.FreightCompany

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	companyID := primitive.NewObjectID().Hex()
	input.FreightCompanyID = companyID
	ConfirmationLink := utils.GenerateConfirmationLinkAdmin(input.FreightCompanyID)
	input.ConfirmationLink = ConfirmationLink
	input.Permissions = "company"

	if input.Name == "" || input.Password == "" || input.Email == "" || input.Description == "" || input.Address == "" || len(input.Contacts) == 0 || len(input.Fleet) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не указаны все обязательные поля"})
		return
	}

	now := time.Now()
	input.CreatedAt = now
	input.UpdatedAt = now

	existingCompany, _ := database.GetFreightCompanyByEmail(input.Email)
	if existingCompany != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Компания с такой электронной почтой уже существует"})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хэшировании пароля"})
		return
	}
	input.Password = hashedPassword

	err = database.CreateFreightCompany(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании компании"})
		fmt.Println(err)
		return
	}

	admins, err := database.GetAllAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении коллекции администраторов"})
		fmt.Println(err)
		return
	}

	err = utils.SendConfirmationEmailToAdmins(admins, input.ConfirmationLink, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при отправке письма с кодом подтверждения"})
		return
	}

	accessToken, err := utils.GenerateAccessToken(input.FreightCompanyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при генерации токена доступа"})
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(input.FreightCompanyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при генерации токена обновления"})
		return
	}

	fmt.Println(accessToken, refreshToken)

	updateData := bson.M{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	err = database.UpdateFreightCompany(input.FreightCompanyID, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении информации о пользователе"})
		return
	}

	// Возвращение успешного ответа с токенами
	c.JSON(http.StatusCreated, gin.H{
		"message": "Регистрация прошла успешно",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
