package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	MongoDBHost         string `env:"MONGODB_HOST" env-default:"localhost:27017"`
	MongoDBCreds        string `env:"MONGODB_CREDS" env-default:""`
	ServiceHost         string `env:"SERVICE_HOST" env-default:"localhost:8080"`
	MongoDBName         string `env:"MONGO_DB_NAME" env-default:"wave"`
	MongoCollectionName string `env:"MONGO_COLLECTION_NAME" env-default:"numbers"`
	User                string `env:"SERVICE_USER" env-default:"user"`
	Password            string `env:"SERVICE_PASSWORD" env-default:"password"`
}

func GetConfig() Config {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatal("Can not read config", err)
	}
	return cfg
}
