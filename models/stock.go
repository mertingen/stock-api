package models

type Stock struct {
	Key   string `json:"key" validate:"required"`
	Value string `json:"value" validate:"required"`
}
