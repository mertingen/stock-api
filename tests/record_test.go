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
	"github.com/mertingen/stock-api/models"
	"github.com/mertingen/stock-api/services"
)

var recordHandler handlers.Record

func init() {
	// Load the .env file in the root directory
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	// Set the env variables regarding the DB
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Get the db client, it doesn't matter which DB it is here.
	dbClient := services.InitDatabase(dbUser, dbPass, dbHost, dbName)
	db, err := dbClient.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Set the record services and the handlers
	recordService := services.InitRecord(db)
	recordHandler = handlers.InitRecord(recordService)
}

func TestSuccessRecordHandler(t *testing.T) {
	payload := struct {
		Startdate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		MinCount  int16  `json:"minCount"`
		MaxCount  int16  `json:"maxCount"`
	}{
		Startdate: "2016-01-26",
		EndDate:   "2018-02-02",
		MinCount:  2700,
		MaxCount:  300,
	}

	// Encode the user payload as JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new request with the JSON payload
	req, err := http.NewRequest("POST", "/api/v1/records", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the CreateUserHandler function with the request and recorder
	handler := http.HandlerFunc(recordHandler.FetchAll)
	handler.ServeHTTP(rr, req)

	// Check the status code returned by the handler
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var resp struct {
		Code    int8            `json:"code"`
		Msg     string          `json:"msg"`
		Records []models.Record `json:"records"`
	}

	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Code != 0 {
		t.Errorf("Code is not match: got %v want %v",
			resp.Code, 0)
	}

	if resp.Msg != "Success" {
		t.Errorf("handler returned unexpected email: got %v want %v",
			resp.Msg, "Success")
	}
}

func TestFailRecordHandler(t *testing.T) {
	payload := struct {
		Startdate string `json:"startDate"`
		EndDate   string `json:"endDate"`
		MinCount  int16  `json:"minCount"`
	}{
		Startdate: "2016-01-26",
		EndDate:   "2018-02-02",
		MinCount:  2700,
	}

	// Encode the user payload as JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new request with the JSON payload
	req, err := http.NewRequest("POST", "/api/v1/records", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a recorder to capture the response
	rr := httptest.NewRecorder()

	// Call the CreateUserHandler function with the request and recorder
	handler := http.HandlerFunc(recordHandler.FetchAll)
	handler.ServeHTTP(rr, req)

	// Check the status code returned by the handler
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var resp struct {
		Code    int8            `json:"code"`
		Msg     string          `json:"msg"`
		Records []models.Record `json:"records"`
	}

	err = json.Unmarshal(rr.Body.Bytes(), &resp)
	if err != nil {
		t.Fatal(err)
	}

	if resp.Code != -3 {
		t.Errorf("Code is not match: got %v want %v",
			resp.Code, 0)
	}
}
