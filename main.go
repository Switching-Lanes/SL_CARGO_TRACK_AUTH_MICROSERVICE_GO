package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/database"
	"github.com/BisquitDubouche/CargoTrack_auth_microservice/app/routes"
)

func main() {
	// CONNECT TO THE DATABASE ...
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("FAILED TO CONNECT TO THE DATABASE:", err)
	}

	// CLOSE THE DATABASE CONNECTION WHEN THE APPLICATION EXITS ...
	defer func() {
		err = db.Disconnect(context.TODO())
		if err != nil {
			log.Fatal("FAILED TO CLOSE THE DATABASE CONNECTION:", err)
		}
	}()

	// RUN DATABASE MIGRATIONS ...
	err = database.RunMigrations(db)
	if err != nil {
		log.Fatal("FAILED TO RUN DATABASE MIGRATIONS:", err)
	}

	// INITIALIZE THE HTTP ROUTES ..
	router := routes.SetupRouter()

	// ATTEMPT TO START THE HTTP SERVER ON PORT 8080 ...
	err = startServer(":8080", router)
	if err != nil {
		log.Println("FAILED TO START THE HTTP SERVER ON PORT 8080:", err)
		log.Println("STARTING THE HTTP SERVER ON A FREE PORT...")
		// IF PORT 8080 IS UNAVAILABLE, TRY TO FIND A FREE PORT AND START THE SERVER ON IT ...
		freePort, err := getFreePort()
		if err != nil {
			log.Fatal("FAILED TO FIND A FREE PORT:", err)
		}
		log.Println("FREE PORT FOUND:", freePort)
		err = startServer(freePort, router)
		if err != nil {
			log.Fatal("FAILED TO START THE HTTP SERVER ON THE FREE PORT:", err)
		}
	}
}

// FUNCTION TO START THE HTTP SERVER ON THE SPECIFIED PORT ...
func startServer(port string, router http.Handler) error {
	log.Println("MICROSERVICE STARTED ON " + port + " PORT")
	// CORS HANDLER
	corsHandler := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			h.ServeHTTP(w, r)
		})
	}
	err := http.ListenAndServe(port, corsHandler(router))
	return err
}

// FUNCTION TO FIND A FREE PORT ...
func getFreePort() (string, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}
	defer l.Close()
	addr := l.Addr().(*net.TCPAddr)
	return fmt.Sprintf(":%d", addr.Port), nil
}
