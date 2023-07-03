package middleware

import (
	"net/http"

	// "github.com/BisquitDubouche/CargoTrack_auth_microservice/app/models"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware выполняет проверку аутентификации и авторизации компании
func AuthFreightCompanyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получение JWT-токена из заголовка Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ошибочка": "Чел, где токен?"})
			return
		}

		// Проверка валидности токена
		_, err := utils.ValidateFreightCompanyToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"ошибочка": "Чел, обнови токен доступа"})
			return
		}

		// Проход к следующему обработчику
		c.Next()
	}
}

// Authorize проверяет права доступа компании
// func AuthorizeFreightCompany(permission string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Проверяем права доступа отправителя
// 		company, exists := c.Get("company")
// 		if !exists {
// 			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Пользователь не найден"})
// 			return
// 		}

// 		// Проверяем права доступа отправителя
// 		companyObj, ok := company.(models.FreightCompany)
// 		if !ok {
// 			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Неверный тип пользователя"})
// 			return
// 		}

// 		if !hasPermission(companyObj.Permissions, permission) {
// 			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Недостаточно прав доступа"})
// 			return
// 		}

// 		// Продолжаем выполнение следующих обработчиков
// 		c.Next()
// 	}
// }
