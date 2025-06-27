package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/moeefa/serenity/db"
	routes "github.com/moeefa/serenity/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using default environment variables")
	}
	
	if port, exists := os.LookupEnv("PORT"); !exists || port == "" {
		os.Setenv("PORT", "8080")
	}

	db.Init()
	routes.Run()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.serveHTTP(w, r)
}
