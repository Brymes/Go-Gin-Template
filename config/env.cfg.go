package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	SendGridAPIKey string
)

func LoadEnv() {
	_ = godotenv.Load()

	SendGridAPIKey = GetEnv("SENDGRID_API_KEY")
}

func GetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalln("Kindly Pass the environment variable named: ", key)
	}
	return val
}

func InitDSN(name string) string {
	var (
		dbHost = os.Getenv("DB_HOST")
		dbUser = os.Getenv("DB_USERNAME")
		dbPass = os.Getenv("DB_PASSWORD")
		dbPort = os.Getenv("DB_PORT")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s ", dbHost, dbUser, dbPass, dbPort, name)

	return dsn
}
