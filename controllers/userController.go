package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/moeefa/serenity/db"
	"github.com/moeefa/serenity/models"
	"golang.org/x/crypto/bcrypt"
)

func GetCurrentUser(c *gin.Context) {
	// The user is already set in the context by the RequireAuth middleware
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get user from context"})
		return
	}

	// Return the user object (sensitive fields will be omitted by JSON serialization)
	c.JSON(http.StatusOK, gin.H{"data": user})
}
// curl -X GET http://localhost:8080/v1/auth/me -H "Authorization: Bearer YOUR_TOKEN"

func UpdateUser(c *gin.Context) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	db := db.GetDB()
	// The user is already set in the context by the RequireAuth middleware
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get user from context"})
		return
	}

	// Bind the input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the user object from context
	currentUser, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse user"})
		return
	}

	// Update user fields
	if input.Name != "" {
		currentUser.Name = input.Name
	}
	if input.Email != "" {
		// Check if email is already in use by another user
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
		// Hash the password using bcrypt
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
			return
		}
		currentUser.Password = string(hashedPassword)
	}

	// Save the updated user to database
	if err := db.Save(&currentUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}

	// Return the updated user without sensitive fields
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id":    currentUser.ID,
			"name":  currentUser.Name,
			"email": currentUser.Email,
			"phone": currentUser.Phone,
		},
	})
}
// curl -X PUT http://localhost:8080/v1/auth/me -H "Authorization: Bearer YOUR_TOKEN" -H "Content-Type: application/json" -d '{"name":"Updated Name","email":"new@example.com","phone":"0987654321","password":"newpassword"}'

func DeleteUser(c *gin.Context) {
	db := db.GetDB()
	
	// The user is already set in the context by the RequireAuth middleware
	user, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get user from context"})
		return
	}

	// Retrieve the user object from context
	currentUser, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse user"})
		return
	}

	// Delete the user from the database
	if err := db.Delete(&currentUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
// curl -X DELETE http://localhost:8080/v1/auth/me -H "Authorization: Bearer YOUR_TOKEN"