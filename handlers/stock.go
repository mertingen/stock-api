package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mertingen/stock-api/models"
	"github.com/mertingen/stock-api/services"
)

type Stock struct {
	stockService services.Stock
}

func InitStock(stockService services.Stock) Stock {
	return Stock{stockService: stockService}
}

func (s *Stock) Fetch(w http.ResponseWriter, r *http.Request) {
	// Set response content type and response struct
	w.Header().Add("Content-Type", "application/json")
	stock := models.Stock{}

	// Accept only HTTP POST verb
	switch r.Method {
	case "POST":
		// Accept only 'application/json' conent-type
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			json.NewEncoder(w).Encode(map[string]string{"error": "Content type is not valid"})
			return
		}

		body, _ := io.ReadAll(r.Body)

		err := json.Unmarshal(body, &stock)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"error": "Request body is not valid"})
			return
		}

		// Validate the HTTP Body fields according to the recordsReq struct
		validate := validator.New()
		err = validate.Struct(stock)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Request body is not valid"})
			return
		}

		stock, err := s.stockService.Save(stock)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "An error occurs while saving in the Redis"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stock)
		return

	case "GET":
		key := r.URL.Query().Get("key")
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Key value is required"})
			return
		}

		stock, err := s.stockService.Fetch(key)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Data is not found"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(stock)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "HTTP verb is not valid"})
		return
	}
}
