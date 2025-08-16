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
