package main

import (
	"github.com/joho/godotenv"
	"go-tenancy/api/router"
	"go-tenancy/api/router/routes"
	"go-tenancy/storage/database"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("[API] error while loading .env file")
	}

	address := os.Getenv("API_URL")

	database.Initialize()

	router.Route("/tenancies", routes.GetTenancies)

	log.Printf("[API] listening on http://%s/\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
