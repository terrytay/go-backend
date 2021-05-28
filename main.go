package main

import (
	"github.com/joho/godotenv"
	"github.com/terrytay/backend-architecture/db"
	"log"
	"os"
)

func main() {
	// Load config from environment variables
	loadConfig()

	dbClient := &db.Client{}
	// TODO: Put to process env
	err := dbClient.Initialise(db.Config{
		Addr:     getEnvVariable("DB_ADDR"),
		User:     getEnvVariable("DB_USER"),
		Password: getEnvVariable("DB_PASS"),
		Database: getEnvVariable("DB_NAME"),
	})

	if err != nil {
		log.Fatal(err)
	}

	initApp(dbClient)
}

func loadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func getEnvVariable(key string) string {
	return os.Getenv(key)
}
