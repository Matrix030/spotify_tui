package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	client_id := os.Getenv("CLIENT_ID")
	client_secret := os.Getenv("CLIENT_SECRET")

	fmt.Println("This is the start")
}
