package handlers

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"time"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/database"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/models"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RegisterHandler обработчик для регистрации заказчика
func ShipperRegisterHandler(c *gin.Context) {
	// Парсинг данных запроса
	var input models.Shipper

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
		return
	}

	// Генерация уникального идентификатора пользователя
	userID := primitive.NewObjectID().Hex()
	input.ShipperID = userID

	input.Permissions = "shipper"

	// Генерация кода подтверждения
	confirmationCode := generateConfirmationCode()
	// Сохранение кода подтверждения и ссылки подтверждения в сущности пользователя
	input.ConfirmationCode = confirmationCode
	input.ConfirmationLink = utils.GenerateConfirmationLink(input.ShipperID)

	// Проверка наличия обязательных полей
	if input.Name == "" || input.Password == "" || input.Email == "" || input.ContactNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не указаны все обязательные поля"})
		return
	}

	// Установка времени создания и обновления профиля
	now := time.Now()
	input.CreatedAt = now
	input.UpdatedAt = now

	// Проверка наличия пользователя с такой же электронной почтой
	existingUser, _ := database.GetShipperByEmail(input.Email)
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Пользователь с такой электронной почтой уже существует"})
		return
	}

	// Хэширование пароля
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хэшировании пароля"})
		return
	}
	input.Password = hashedPassword

	// Создание пользователя
	err = database.CreateUser(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании пользователя"})
		fmt.Println(err)
		return
	}

	// Отправка письма с кодом подтверждения
	err = utils.SendConfirmationEmail(input.Email, confirmationCode, input.ConfirmationLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при отправке письма с кодом подтверждения"})
		return
	}

	// Генерация и подпись токена доступа
	accessToken, err := utils.GenerateAccessToken(input.ShipperID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при генерации токена доступа"})
		return
	}

	// Генерация и подпись токена обновления
	refreshToken, err := utils.GenerateRefreshToken(input.ShipperID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при генерации токена обновления"})
		return
	}

	// Сохранение токенов в сущности пользователя
	input.AccessToken = accessToken
	input.RefreshToken = refreshToken

	err = database.UpdateShipper(input.ShipperID, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении информации о пользователе"})
		return
	}

	// Возвращение успешного ответа с токенами
	c.JSON(http.StatusCreated, gin.H{
		"message": "Регистраци прошла успешно",
	})

}

// Функция генерации кода подтверждения
func generateConfirmationCode() string {
	code := make([]byte, 3)
	rand.Read(code)

	// Преобразование кода в строку шестнадцатеричного формата
	return fmt.Sprintf("%x", code)
}
