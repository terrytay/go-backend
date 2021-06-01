package main

import (
	"github.com/joho/godotenv"
	"github.com/terrytay/go-backend/db"
	"github.com/terrytay/go-backend/helpers"
	"log"
)

func main() {

	// Load config from environment variables
	loadConfig()

	dbClient := &db.Client{}
	// TODO: Put to process env
	err := dbClient.Initialise(db.Config{
		Addr:     helpers.GetEnvVariable("DB_ADDR"),
		User:     helpers.GetEnvVariable("DB_USER"),
		Password: helpers.GetEnvVariable("DB_PASS"),
		Database: helpers.GetEnvVariable("DB_NAME"),
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