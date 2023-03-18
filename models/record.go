package models

type Record struct {
	Key        string `json:"key" bson:"key"`
	CreatedAt  string `json:"createdAt" bson:"createdAt"`
	TotalCount int    `json:"totalCount" bson:"totalCount"`
}
