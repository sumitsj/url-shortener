package services

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InitMongoDB() {
	err := mgm.SetDefaultConfig(nil, Config.MongodbDatabase, options.Client().ApplyURI(Config.MongodbUri))
	if err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB!")
}
