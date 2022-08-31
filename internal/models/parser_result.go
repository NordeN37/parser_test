package models

import "time"

type ParserResult struct {
	Date        time.Time              `bson:"date"`
	Name        string                 `bson:"name"`
	Description map[string]interface{} `bson:"description"`
}
