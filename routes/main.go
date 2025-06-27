package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")

	api.GET("/", func(c *gin.Context) {
		c.Status(200)
	})

	addAuthRoutes(api)
	addUserRoutes(api)
	addRecommendationsRoute(api)
}

