package services

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/spf13/viper"
)

type AppConfig interface {
	GetServerPort() string
	GetServerAddr() string
	GetMongoDbUri() string
	GetMongoDbDatabase() string
}
type config struct {
	ServerPort      string `mapstructure:"SERVER_PORT"`
	ServerAddr      string `mapstructure:"SERVER_ADDR"`
	MongodbUri      string `mapstructure:"MONGO_URI"`
	MongodbDatabase string `mapstructure:"MONGO_DATABASE"`
	Mode            string `mapstructure:"MODE"`
}

func (config *config) validate() error {
	return validation.ValidateStruct(config,
		validation.Field(&config.ServerPort, is.Port),
		validation.Field(&config.ServerAddr, validation.Required),

		validation.Field(&config.MongodbUri, validation.Required),
		validation.Field(&config.MongodbDatabase, validation.Required),

		validation.Field(&config.Mode, validation.In("debug", "release")),
	)
}

func (config *config) GetServerPort() string {
	return config.ServerPort
}

func (config *config) GetServerAddr() string {
	return config.ServerAddr
}

func (config *config) GetMongoDbUri() string {
	return config.MongodbUri
}

func (config *config) GetMongoDbDatabase() string {
	return config.MongodbDatabase
}

func LoadConfig() AppConfig {
	var configuration config
	v := viper.New()
	v.AutomaticEnv()
	v.SetDefault("SERVER_PORT", "8080")
	v.SetDefault("MODE", "debug")
	v.SetConfigType("dotenv")
	v.SetConfigName(".env")
	v.AddConfigPath("./")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&configuration); err != nil {
		panic(err)
	}

	if err := configuration.validate(); err != nil {
		panic(err)
	}

	return &configuration
}
