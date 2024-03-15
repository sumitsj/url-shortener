package services

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InitMongoDB(config AppConfig) {
	err := mgm.SetDefaultConfig(nil, config.GetMongoDbDatabase(), options.Client().ApplyURI(config.GetMongoDbUri()))
	if err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB!")
}
