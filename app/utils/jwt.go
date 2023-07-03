package utils

import (
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	AccessTokenSecret = []byte("switchinglanes-cargo-track")
)

type FreightCompanyClaims struct {
	FreightCompanyID string `json:"id"`
	jwt.StandardClaims
}

type ShipperClaims struct {
	ShipperID string `json:"id"`
	jwt.StandardClaims
}

type RefreshClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func GetCurrentTimestamp() int64 {
	return time.Now().Unix()
}

// Генерация токена доступа
func GenerateAccessToken(ID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  ID,
		"exp": time.Now().Add(time.Hour * 1).Unix(), // Время истечения срока действия токена (например, 1 час)
	})

	// Подпись токена
	accessToken, err := token.SignedString(AccessTokenSecret)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// Генерация токена обновления

func GenerateRefreshToken(ID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  ID,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	// Подпись токена
	refreshToken, err := token.SignedString(AccessTokenSecret)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func ValidateShipperAccessToken(tokenString string) (*jwt.Token, error) {
	// Извлечение токена из строкового представления
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Проверка валидности токена
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Здесь вы должны вернуть ваш секретный ключ для проверки подписи токена
		return []byte(AccessTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return token, nil
}

func ValidateFreightCompanyToken(tokenString string) (*jwt.Token, error) {
	// Извлечение токена из строкового представления
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Проверка валидности токена
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Здесь вы должны вернуть ваш секретный ключ для проверки подписи токена
		return []byte(AccessTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return token, nil
}

func ValidateRefreshToken(tokenString string) (*jwt.Token, error) {
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Проверка валидности токена
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Здесь вы должны вернуть ваш секретный ключ для проверки подписи токена
		return []byte(AccessTokenSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return token, nil
}
