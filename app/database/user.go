package database

import (
	"context"
	"errors"
	"log"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserByRefreshToken(tokenString string) (*models.User, error) {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("users")

	var User models.User
	err = collection.FindOne(context.Background(), bson.M{"refresh_token": tokenString}).Decode(&User)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &User, nil
}

func UpdateUser(userID string, input models.User) error {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("freight_company_employees")
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": input}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
