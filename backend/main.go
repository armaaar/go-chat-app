package main

import (
	"log"
	"net/http"
	"os"

	"github.com/armaaar/go-chat-app/database"
	"github.com/armaaar/go-chat-app/health"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	database.SetupDatabase()

	baseRoute := "/"
	health.RegisterRoutes(baseRoute)

	log.Printf("Running: http://localhost:%s", os.Getenv("PORT"))
	log.Panic(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
