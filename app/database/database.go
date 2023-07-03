package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DATABASE_NAME = "users"
)

// Подключение к базе данных MongoDB Atlas
func ConnectDB() (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://switchinglaneskg:FBYaeoCTXGitrbnA@cluster.1lb2zks.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)

	// Создаем нового клиента и подключаемся к серверу
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Println("Ошибка при подключении к базе данных:", err)
		return nil, err
	}

	// Проверка подключения к базе данных
	ctxPing, cancelPing := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancelPing()

	err = client.Ping(ctxPing, nil)
	if err != nil {
		log.Println("Ошибка при проверке подключения к базе данных:", err)
		client.Disconnect(ctx)
		return nil, err
	}

	log.Println("Успешное подключение к базе данных MongoDB Atlas")

	return client, nil
}

// Выполняет миграции базы данных
func RunMigrations(client *mongo.Client) error {
	err := createAdministratorCollection(client)
	if err != nil {
		return err
	}

	err = createFreightCompaniesCollection(client)
	if err != nil {
		return err
	}

	err = createUsersCollection(client)
	if err != nil {
		return err
	}

	err = createShipperCollection(client)
	if err != nil {
		return err
	}

	return nil
}

// Создает коллекцию "administrators"
func createAdministratorCollection(client *mongo.Client) error {
	collection := client.Database(DATABASE_NAME).Collection("administrators")

	// Создаем уникальный индекс по полю "admin_id"
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"admin_id": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return err
	}

	return nil
}

// Создает коллекцию "companies"
func createFreightCompaniesCollection(client *mongo.Client) error {
	collection := client.Database(DATABASE_NAME).Collection("freight_companies")

	// Создаем уникальный индекс по полю "employee_id"
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"freight_company_id": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return err
	}

	return nil
}

// Создает коллекцию "employees"
func createUsersCollection(client *mongo.Client) error {
	collection := client.Database(DATABASE_NAME).Collection("users")

	// Создаем уникальный индекс по полю "employee_id"
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"user_id": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return err
	}

	return nil
}

// Создает коллекцию "shippers"
func createShipperCollection(client *mongo.Client) error {
	collection := client.Database(DATABASE_NAME).Collection("shippers")

	// Создаем уникальный индекс по полю "shipper_id"
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"shipper_id": 1},
		Options: options.Index().SetUnique(true),
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return err
	}

	return nil
}