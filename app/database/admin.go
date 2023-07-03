package database

import (
	"context"
	"log"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Получение списка всех администраторов
func GetAllAdmins() ([]models.Administrator, error) {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("administrators")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("Ошибка при выполнении запроса:", err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	admins := make([]models.Administrator, 0)
	for cursor.Next(context.TODO()) {
		admin := models.Administrator{}
		err := cursor.Decode(&admin)
		if err != nil {
			log.Fatal("Ошибка при декодировании результата:", err)
			return nil, err
		}
		admins = append(admins, admin)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal("Ошибка после обхода всех документов:", err)
		return nil, err
	}

	return admins, nil
}
