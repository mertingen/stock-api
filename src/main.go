package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/mertingen/stock-api/handlers"
	"github.com/mertingen/stock-api/services"
)

func main() {
	// Load the .env file in the root directory
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Error loading .env file!")
	}

	// Set the env variables regarding the DB
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Set the env variables regarding the DB
	redisUser := os.Getenv("REDIS_USER")
	redisPass := os.Getenv("REDIS_PASS")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	// Get the db client, it doesn't matter which DB it is here.
	dbClient := services.InitDatabase(dbUser, dbPass, dbHost, dbName)
	db, err := dbClient.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Get the Redis client
	redisClient := services.InitRedis(redisUser, redisPass, redisHost, redisPort)
	redis, err := redisClient.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Set the record services and the handlers
	recordService := services.InitRecord(db)
	recordHandler := handlers.InitRecord(recordService)
	// Set the stock services and the handlers
	stockService := services.InitStock(redis)
	stockHandler := handlers.InitStock(stockService)

	// Create the HTTP handlers
	http.HandleFunc("/api/v1/records", recordHandler.FetchAll)
	http.HandleFunc("/api/v1/stocks", stockHandler.Fetch)

	// Start the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Listening on port %s...\n", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
