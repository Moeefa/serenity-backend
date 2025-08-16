package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moeefa/serenity/db"
	"github.com/moeefa/serenity/models"
	"golang.org/x/crypto/bcrypt"
)

func GetCurrentUser(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get user from context"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	db := db.GetDB()
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get user from context"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse user"})
		return
	}

	if input.Name != "" {
		currentUser.Name = input.Name
	}
	if input.Email != "" {
		var existingUser models.User
		if result := db.Where("email = ? AND id != ?", input.Email, currentUser.ID).First(&existingUser); result.RowsAffected > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email already in use"})
			return
		}
		currentUser.Email = input.Email
	}
	if input.Phone != "" {
		currentUser.Phone = input.Phone
	}
	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
			return
		}
		currentUser.Password = string(hashedPassword)
	}

	if err := db.Save(&currentUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":    currentUser.ID,
			"name":  currentUser.Name,
			"email": currentUser.Email,
			"phone": currentUser.Phone,
		},
	})
}

func DeleteUser(c *gin.Context) {
	db := db.GetDB()
	
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get user from context"})
		return
	}

	currentUser, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse user"})
		return
	}

	if err := db.Delete(&currentUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
