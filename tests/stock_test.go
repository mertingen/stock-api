package main_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/mertingen/stock-api/handlers"
	"github.com/mertingen/stock-api/services"
)

var stockHandler handlers.Stock

func init() {
	// Load the .env file in the root directory
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	// Set the env variables regarding the DB
	redisUser := os.Getenv("REDIS_USER")
	redisPass := os.Getenv("REDIS_PASS")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	// Get the Redis client
	redisClient := services.InitRedis(redisUser, redisPass, redisHost, redisPort)
	redis, err := redisClient.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Set the stock services and the handlers
	stockService := services.InitStock(redis)
	stockHandler = handlers.InitStock(stockService)
}

func TestSuccessStockHandler(t *testing.T) {
	payload := struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{
		Key:   "test-in-go-key",
		Value: "test-in-go-val",
	}

	// Encode the user payload as JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new request with the JSON payload
	req, err := http.NewRequest("POST", "/api/v1/stocks", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the CreateUserHandler function with the request and recorder
	handler := http.HandlerFunc(stockHandler.Fetch)
	handler.ServeHTTP(rr, req)

	// Check the status code returned by the handler
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var resp struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Key != payload.Key {
		t.Errorf("Key is not match: got %v want %v",
			resp.Key, payload.Key)
	}

	if resp.Value != payload.Value {
		t.Errorf("Value is not match: got %v want %v",
			resp.Key, payload.Value)
	}
}
