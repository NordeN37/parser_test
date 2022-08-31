package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"parser_test/config"
	"parser_test/internal/store/mongo/repo/parser"
)

type MongoRepo struct {
	Parser parser.IParser
}

func NewMongoRepo(mongo *mongo.Client) (*MongoRepo, error) {
	result := &MongoRepo{
		parser.NewParser(mongo, config.GetSettings().MongoDBName),
	}
	return result, nil
}
