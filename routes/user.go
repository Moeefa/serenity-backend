package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/moeefa/serenity/controllers"
	"github.com/moeefa/serenity/middlewares"
)

func addUserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/user")

	router.GET("/me", middlewares.CheckAuth, controllers.GetCurrentUser)
	router.PUT("/me", middlewares.CheckAuth, controllers.UpdateUser)
	router.DELETE("/me", middlewares.CheckAuth, controllers.DeleteUser)
}