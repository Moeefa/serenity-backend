package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moeefa/serenity/controllers"
	"github.com/moeefa/serenity/middlewares"
)

func addAuthRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.LoginUser)
	router.GET("/verify", middlewares.CheckAuth, controllers.VerifyUser)
}