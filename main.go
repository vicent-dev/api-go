package main

import (
	"api/src/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	server.Start(os.Getenv("API_PORT"))
}
