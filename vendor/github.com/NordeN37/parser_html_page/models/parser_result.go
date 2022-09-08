package models

type ParseSelectionResult struct {
	Value      *string
	FoundValue *[]ParseSelectionResult
}
