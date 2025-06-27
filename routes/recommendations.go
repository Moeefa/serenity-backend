package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moeefa/serenity/controllers"
)

func addRecommendationsRoute(rg *gin.RouterGroup) {
	recommendationsGroup := rg.Group("/recommendations")

	recommendationsGroup.GET("/", controllers.GetRecommendations)
	recommendationsGroup.POST("/", controllers.CreateRecommendation)
}