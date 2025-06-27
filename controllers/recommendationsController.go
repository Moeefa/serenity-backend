package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moeefa/serenity/db"
	"github.com/moeefa/serenity/models"
)

func GetRecommendations(c *gin.Context) {
	var recommendationsList []models.Recommendation
	result := db.GetDB().Find(&recommendationsList)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, recommendationsList)
}

func CreateRecommendation(c *gin.Context) {
	var recommendation models.Recommendation
	if err := c.ShouldBindJSON(&recommendation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.GetDB().Create(&recommendation)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, recommendation)
}
// curl http://localhost:8080/v1/recommendations/ -X POST -H "Content-Type: application/json" -d '{"title":"New Recommendation","description":"This is a new recommendation","tags":["tag1","tag2"],"time_required":"30 minutes"}'