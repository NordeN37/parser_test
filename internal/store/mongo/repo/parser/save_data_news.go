package parser

import (
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"parser_test/internal/models"
)

type IParser interface {
	Store(document *models.ParserResult) error
}

func NewParser(mongo *mongo.Client, dbName string) IParser {
	return &parserResult{
		mongo:          mongo,
		dbName:         dbName,
		collectionName: "parser_result",
	}
}

type parserResult struct {
	mongo          *mongo.Client
	dbName         string
	collectionName string
}

func (pr *parserResult) Store(document *models.ParserResult) error {
	collection := pr.mongo.Database(pr.dbName).Collection(document.Name)
	insertResult, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		return errors.Wrapf(err, "ошибка обращения к collection.InsertOne")
	}
	log.Println("[INFO] : ", "insertResult : ", insertResult)
	return nil
}
