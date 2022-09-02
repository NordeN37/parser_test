package models

import "time"

type ParserResult struct {
	Date        time.Time               `bson:"date"`
	Name        string                  `bson:"name"`
	Description *[]ParseSelectionResult `bson:"description"`
}

type ParseSelectionResult struct {
	Value      *string
	FoundValue *[]ParseSelectionResult
}
