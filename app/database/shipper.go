package database

import (
	"context"
	"errors"
	"log"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateUser создает нового пользователя в базе данных
func CreateUser(user *models.Shipper) error {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("shippers")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}

// GetUserByEmail извлекает пользователя из базы данных по его адресу электронной почты
func GetShipperByEmail(email string) (*models.Shipper, error) {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("shippers")

	var user models.Shipper
	err = collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByID извлекает пользователя из базы данных по его идентификатору
func GetShipperByID(userID string) (*models.Shipper, error) {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("shippers")

	var user models.Shipper
	err = collection.FindOne(context.Background(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByConfirmationCode извлекает пользователя из базы данных по коду подтверждения электронной почты
func GetShipperByConfirmationLink(confirmationLink string) (*models.Shipper, error) {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("shippers")

	var user models.Shipper
	err = collection.FindOne(context.Background(), bson.M{"confirmation_link": confirmationLink}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("shipper not found")
		}
		return nil, err
	}

	return &user, nil
}

func GetShipperByConfirmationCode(confirmationCode string) (*models.Shipper, error) {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("shippers")

	var user models.Shipper
	err = collection.FindOne(context.Background(), bson.M{"confirmation_code": confirmationCode}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("shipper not found")
		}
		return nil, err
	}

	return &user, nil
}

func GetShipperByRefreshToken(tokenString string) (*models.Shipper, error) {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("shippers")

	var Shipper models.Shipper
	err = collection.FindOne(context.Background(), bson.M{"refresh_token": tokenString}).Decode(&Shipper)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("shipper not found")
		}
		return nil, err
	}

	return &Shipper, nil
}

// UpdateUser обновляет поля пользователя в базе данных
func UpdateShipper(userID string, input models.Shipper) error {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("shippers")
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": input}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

// UpdateEmailConfirmationStatus обновляет статус подтверждения электронной почты пользователя
func UpdateEmailConfirmationStatus(userID string, isConfirmed bool) error {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("shippers")
	filter := bson.M{"shipper_id": userID}
	update := bson.M{"$set": bson.M{"email_confirmed": isConfirmed}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func UpdateShipperTokens(shipper *models.Shipper) (*models.Shipper, error) {
	client, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("shippers")

	// Создаем фильтр для поиска записи Shipper по идентификатору
	filter := bson.M{"shipper_id": shipper.ShipperID}

	var user models.Shipper
	err = collection.FindOne(context.Background(), bson.M{"_id": shipper.ShipperID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Создаем обновляемый документ с новыми значениями токенов и их сроками действия
	update := bson.M{
		"$set": bson.M{
			"access_token":  shipper.AccessToken,
			"refresh_token": shipper.RefreshToken,
		},
	}

	// Обновляем запись Shipper в базе данных
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser удаляет пользователя из базы данных
func DeleteShipper(userID string) error {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("shippers")
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": userID})
	if err != nil {
		return err
	}

	return nil
}
