package main

import (
	"github.com/ara-o/gren/server"
)

func main() {
	server := server.New(":8080")
	server.Start()
}
