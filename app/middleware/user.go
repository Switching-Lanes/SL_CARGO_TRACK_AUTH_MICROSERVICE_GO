package middleware

import (
	"net/http"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/utils"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware выполняет проверку аутентификации и авторизации пользователя
func AuthUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получение JWT-токена из заголовка Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"ошибочка": "Чел, где токен?"})
			return
		}

		// Проверка валидности токена
		_, err := utils.ValidateShipperAccessToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"ошибочка": "Чел, обнови токен доступа"})
			return
		}

		// Проход к следующему обработчику
		c.Next()
	}
}