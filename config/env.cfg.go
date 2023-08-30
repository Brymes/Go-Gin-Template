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

	//InitSendGrid()
}

func InitSendGrid() {
	SendGridAPIKey = os.Getenv("SENDGRID_API_KEY")
	if SendGridAPIKey == "" {
		log.Fatalln("Kindly Pass SendGrid API key as an environment variable named SENDGRID_API_KEY")
	}
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
