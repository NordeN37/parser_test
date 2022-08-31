package mongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"parser_test/config"
	"time"
)

func NewMongoClient() (*mongo.Client, error) {
	var cred options.Credential

	cred.Username = config.GetSettings().MongoUserName
	cred.Password = config.GetSettings().MongoPassword
	ticker := time.NewTicker(1 * time.Nanosecond)
	timeout := time.After(15 * time.Minute)
	seconds := 1
	client, err := mongo.NewClient(options.Client().ApplyURI(getAddress()).SetAuth(cred))
	if err != nil {
		return nil, err
	}
	try := 0
	for {
		select {
		case <-ticker.C:
			try++
			ticker.Stop()
			if err := client.Connect(context.Background()); err != nil {
				log.Printf("[ERROR] : %s, %s", err.Error(), fmt.Sprintf("не удалось установить соединение с MongoDB, попытка № %d", try))
				ticker = time.NewTicker(time.Duration(seconds) * time.Second)
				seconds *= 2
				if seconds > 30 {
					seconds = 30
				}
				continue
			}
			if err = client.Ping(context.Background(), nil); err != nil {
				return nil, err
			}
			return client, nil
		case <-timeout:
			return nil, errors.New("MongoDB: connection timeout")
		}
	}
}

func getAddress() string {
	settings := config.GetSettings()
	return "mongodb://" + settings.MongoHost + ":" + settings.MongoPort
}
