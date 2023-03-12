package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"porteonBackend/models"
	"time"
)

// GetJob retrieves a job from the database by ID
func GetJob(id string) (*models.Job, error) {
	var job models.Job
	if err := models.DB.First(&job, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &job, nil
}

// CreateJob creates a new job in the database
func CreateJob(newJob models.Job) (*models.Job, error) {
	newJob.ID = uuid.NewString()
	newJob.CreatedDate = time.Now()
	newJob.Status = models.JobNotYetStarted
	if err := models.DB.Create(&newJob).Error; err != nil {
		return nil, err
	}
	return &newJob, nil
}

// UpdateJob updates an existing job in the database
func UpdateJob(job *models.Job) error {
	if err := models.DB.Save(&job).Error; err != nil {
		return err
	}
	return nil
}

// DeleteJob deletes a job from the database by ID
func DeleteJob(id string) error {
	if err := models.DB.Delete(&models.Job{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

// GetJobs Handler function to get all jobs
func GetJobs(c *gin.Context) {
	var jobs []models.Job
	models.DB.Find(&jobs)
	c.JSON(http.StatusOK, gin.H{"data": jobs})
}

// GetJobHandler Handler function to get a specific job
func GetJobHandler(c *gin.Context) {
	id := c.Param("id")
	job, err := GetJob(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving job"})
		return
	}
	if job == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": job})
}

// Handler function to update an existing job
func UpdateJobHandler(c *gin.Context) {
	id := c.Param("id")
	job, err := GetJob(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving job"})
		return
	}
	if job == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
		return
	}

	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := UpdateJob(job); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating job"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": job})
}

// DeleteJobHandler Handler function to delete a job
func DeleteJobHandler(c *gin.Context) {
	id := c.Param("id")
	if err := DeleteJob(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting job"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
}
