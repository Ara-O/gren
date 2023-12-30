package main

import (
	"log"

	"github.com/ara-o/gren/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := server.New(":8080")
	server.Start()
}
