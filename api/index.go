package handler

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/moeefa/serenity/routes"
	"github.com/moeefa/serenity/db"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using default environment variables")
	}
	
	if port, exists := os.LookupEnv("PORT"); !exists || port == "" {
		os.Setenv("PORT", "8080")
	}

	if err := db.Init(); err != nil {
		log.Println("DB Init error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	router := gin.New()
	routes.RegisterRoutes(router)
	router.ServeHTTP(w, r)
}

