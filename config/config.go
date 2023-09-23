package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Secrets struct {
	DatabaseName string `json:"DB_NAME"`
	DatabaseURL string `json:"DB_URL"`
	RedisUrl string `json:"REDIS_URL"`
	Port string
}

var ss Secrets

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading env file")
	}

	ss = Secrets{}

	ss.DatabaseName = os.Getenv("DB_NAME")
	ss.DatabaseURL = os.Getenv("DB_URL")
	ss.RedisUrl = os.Getenv("REDIS_URL")

	if ss.Port = os.Getenv("PORT"); ss.Port == "" {
		ss.Port = "80"
	}
}

func GetSecrets() Secrets {
	return ss
}
