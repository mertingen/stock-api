package services

import (
	"github.com/go-redis/redis"
	"github.com/mertingen/stock-api/models"
)

type Stock struct {
	redis *redis.Client
}

func InitStock(redis *redis.Client) Stock {
	return Stock{redis: redis}
}

func (s *Stock) Fetch(key string) (models.Stock, error) {
	// Init models.Record format for rows
	stock := models.Stock{}

	// Get the value of a key
	val, err := s.redis.Get(key).Result()
	if err != nil {
		return stock, err
	}

	stock.Key = key
	stock.Value = val

	return stock, nil
}

func (s *Stock) Save(stock models.Stock) (models.Stock, error) {
	// Call set with a `Key` and a `Value`.
	err := s.redis.Set(stock.Key, stock.Value, 0).Err()
	// If there has been an error setting the value
	// Handle the error
	if err != nil {
		return stock, err
	}

	return stock, nil
}
