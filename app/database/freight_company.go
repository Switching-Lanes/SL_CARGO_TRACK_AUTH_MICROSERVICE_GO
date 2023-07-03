package database

import (
	"context"
	"errors"
	"log"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateFreightCompany создает новую грузовую компанию в базе данных

func CreateFreightCompany(freightCompany *models.FreightCompany) error {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("freight_companies")
	_, err = collection.InsertOne(context.Background(), freightCompany)
	if err != nil {
		return err
	}

	return nil
}

func RegisterFreightCompanyEmployee(employee models.User) error {
	// Подключение к базе данных
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("users")

	_, err = collection.InsertOne(context.Background(), employee)
	if err != nil {
		return err
	}

	return nil
}

func UpdateFreightCompany(freightCompanyID string, updatedFields bson.M) error {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("freight_companies")
	filter := bson.M{"_id": freightCompanyID}
	update := bson.M{"$set": updatedFields}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func UpdateFreightCompanyRegisteredStatus(freightCompanyID string, isRegistered bool) error {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("freight_companies")
	filter := bson.M{"freight_company_id": freightCompanyID}
	update := bson.M{"$set": bson.M{"is_registered": isRegistered}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func GetFreightCompanyByID(freightCompanyID string) (*models.FreightCompany, error) {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("freight_companies")

	var company models.FreightCompany
	err = collection.FindOne(context.Background(), bson.M{"_id": freightCompanyID}).Decode(&company)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("f company not found")
		}
		return nil, err
	}

	return &company, nil
}

func GetFreightCompanyByEmail(email string) (*models.FreightCompany, error) {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("freight_companies")

	var FreightCompany models.FreightCompany
	err = collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&FreightCompany)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("freight company not found")
		}
		return nil, err
	}

	return &FreightCompany, nil
}

func GetFreightCompanyByConfirmationLink(confirmationLink string) (*models.FreightCompany, error) {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("freight_companies")

	var company models.FreightCompany
	err = collection.FindOne(context.Background(), bson.M{"confirmation_link": confirmationLink}).Decode(&company)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("freight company not found")
		}
		return nil, err
	}

	return &company, nil
}

func GetFreightCompanyByRefreshToken(tokenString string) (*models.FreightCompany, error) {
	client, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("freight_companies")

	var FreightCompany models.FreightCompany
	err = collection.FindOne(context.Background(), bson.M{"refresh_token": tokenString}).Decode(&FreightCompany)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("freight company not found")
		}
		return nil, err
	}

	return &FreightCompany, nil
}

func UpdateFreightCompanyTokens(company *models.FreightCompany) error {
	client, err := ConnectDB()
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	collection := client.Database(DATABASE_NAME).Collection("freight_companies")

	// Создаем фильтр для поиска записи FreightCompany по идентификатору
	filter := bson.M{"freight_company_id": company.FreightCompanyID}

	// Создаем обновляемый документ с новыми значениями токенов и их сроками действия
	update := bson.M{
		"$set": bson.M{
			"access_token":      company.AccessToken,
			"refresh_token":     company.RefreshToken,
		},
	}

	// Обновляем запись FreightCompany в базе данных
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
