package main

import (
	"github.com/joho/godotenv"
	"go-tenancy/database"
	"go-tenancy/router"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error while loading .env file")
	}

	address := os.Getenv("API_URL")

	database.Initialize()

	router.Route("/", router.GetTenancies)

	log.Printf("[API] listening on http://%s/\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
