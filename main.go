package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/majidehamide/jwt-redis/config"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("No config detected")
	}
	_, err := config.ConnectDB()
	if err != nil {
		log.Panic("No connection to database")
	}
}
