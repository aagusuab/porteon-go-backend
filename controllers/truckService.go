package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"time"

	"porteonBackend/models"
)

// GetTruck retrieves a truck from the database by ID
func GetTruck(id string) (*models.Truck, error) {
	var truck models.Truck
	if err := models.DB.First(&truck, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &truck, nil
}

// CreateTruck creates a new truck in the database
func CreateTruck(newTruck models.Truck) (*models.Truck, error) {
	newTruck.ID = uuid.NewString()
	newTruck.VehicleCreatedYear = uint(time.Now().Year())
	newTruck.Status = models.TruckReady
	if err := models.DB.Create(&newTruck).Error; err != nil {
		return nil, err
	}
	return &newTruck, nil
}

// UpdateTruck updates an existing truck in the database
func UpdateTruck(truck *models.Truck) error {
	if err := models.DB.Save(&truck).Error; err != nil {
		return err
	}
	return nil
}

// DeleteTruck deletes a truck from the database by ID
func DeleteTruck(id string) error {
	if err := models.DB.Delete(&models.Truck{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

// GetTrucks Handler function to get all trucks
func GetTrucks(c *gin.Context) {
	var trucks []models.Truck
	models.DB.Find(&trucks)
	c.JSON(http.StatusOK, gin.H{"data": trucks})
}

// GetTruckHandler Handler function to get a specific truck
func GetTruckHandler(c *gin.Context) {
	id := c.Param("id")
	truck, err := GetTruck(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving truck"})
		return
	}
	if truck == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Truck not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": truck})
}

// Handler function to create a new truck
func CreateTruckHandler(c *gin.Context) {
	var truck models.Truck
	if err := c.ShouldBindJSON(&truck); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTruck, err := CreateTruck(truck)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating truck"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": newTruck})
}

// Handler function to update an existing truck
func UpdateTruckHandler(c *gin.Context) {
	id := c.Param("id")
	truck, err := GetTruck(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving truck"})
		return
	}
	if truck == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Truck not found"})
		return
	}

	if err := c.ShouldBindJSON(&truck); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := UpdateTruck(truck); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating truck"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": truck})
}

// DeleteTruckHandler DeleteJobHandler Handler function to delete a job
func DeleteTruckHandler(c *gin.Context) {
	id := c.Param("id")
	if err := DeleteTruck(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting Truck"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Truck deleted successfully"})
}
