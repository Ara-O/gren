package main

import (
	"fmt"
	"log"
	"os"
	"github.com/ara-o/gren/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		panic("Failed to read PORT from environment variables")
	}

	server := server.New(fmt.Sprintf(":%s", port))
	err = server.Start()
	if err != nil {
		panic(err)
	}
}
