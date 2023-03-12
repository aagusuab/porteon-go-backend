package controllers

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"porteonBackend/models"
	"time"

	"github.com/gin-gonic/gin"
)

// GetUser retrieves a user from the database by ID
func GetUser(id string) (*models.User, error) {
	var user models.User
	if err := models.DB.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user in the database
func CreateUser(newUser models.User) (*models.User, error) {
	newUser.ID = uuid.NewString()
	newUser.JoinDate = time.Now()
	newUser.Status = models.UserRegistrationInProg
	if err := models.DB.Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

// UpdateUser updates an existing user in the database
func UpdateUser(user *models.User) error {
	if err := models.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user from the database by ID
func DeleteUser(id string) error {
	if err := models.DB.Delete(&models.User{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

// GetUsers Handler function to get all users
func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUserHandler Handler function to get a specific user
func GetUserHandler(c *gin.Context) {
	id := c.Param("id")
	user, err := GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Handler function to create a new user
func CreateUserHandler(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// Handler function to update an existing user
func UpdateUserHandler(c *gin.Context) {
	id := c.Param("id")
	user, err := GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUserHandler Handler function to delete a user
func DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	if err := DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
