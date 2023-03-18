package services

import (
	"context"
	"time"

	"github.com/mertingen/stock-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Record struct {
	db *mongo.Database
}

func InitRecord(db *mongo.Database) Record {
	return Record{db: db}
}

func (r *Record) FetchAll(filter bson.M) ([]models.Record, error) {
	// Init models.Record format for rows
	records := make([]models.Record, 0)
	coll := r.db.Collection("records")

	// Create a cancellation for DB queries according to the timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var record models.Record
		cursor.Decode(&record)
		records = append(records, record)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return records, nil
}

func (r *Record) ValidateDate(d string) bool {
	// "2006-01-02" is a predefined constant in the time package that represents the format "YYYY-MM-DD"
	// It does not check if the date is a valid date according to the Gregorian calendar.
	_, err := time.Parse("2006-01-02", d)
	return err == nil
}
