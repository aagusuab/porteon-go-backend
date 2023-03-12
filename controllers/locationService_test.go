package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"porteonBackend/models"
)

// Location Test Functions

func TestCreateLocation(t *testing.T) {
	newLocation := models.Location{
		Name:        "Headquarters",
		Address:     "123 Main St",
		City:        "San Francisco",
		State:       "CA",
		ZipCode:     "94105",
		Country:     "USA",
		Latitude:    37.7749,
		Longitude:   -122.4194,
		CreatedDate: time.Now(),
	}

	router := gin.Default()
	router.POST("/locations", CreateLocationHandler)

	locationJSON, err := json.Marshal(newLocation)
	if err != nil {
		t.Errorf("Error marshalling new location: %s", err)
	}

	req, err := http.NewRequest("POST", "/locations", bytes.NewBuffer(locationJSON))
	if err != nil {
		t.Errorf("Error creating new request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response gin.H
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response: %s", err)
	}

	message, ok := response["message"]
	if !ok || message != "New location created" {
		t.Errorf("Unexpected response message: %v", response)
	}
}

func TestGetLocation(t *testing.T) {
	location, err := CreateLocation(models.Location{
		Name:        "Headquarters",
		Address:     "123 Main St",
		City:        "San Francisco",
		State:       "CA",
		ZipCode:     "94105",
		Country:     "USA",
		Latitude:    37.7749,
		Longitude:   -122.4194,
		CreatedDate: time.Now(),
	})
	if err != nil {
		t.Errorf("Error creating new location: %s", err)
	}

	router := gin.Default()
	router.GET("/locations/:id", GetLocationHandler)

	req, err := http.NewRequest("GET", "/locations/"+location.ID, nil)
	if err != nil {
		t.Errorf("Error creating new request: %s", err)
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response gin.H
	err = json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Error unmarshalling response: %s", err)
	}

	data, ok := response["data"].(map[string]interface{})
	if !ok || data["id"] != location.ID {
		t.Errorf("Unexpected response data: %v", response)
	}
}

func TestUpdateLocation(t *testing.T) {
	location, err := CreateLocation(models.Location{
		Name:        "Headquarters",
		Address:     "123 Main St",
		City:        "San Francisco",
		State:       "CA",
		ZipCode:     "94105",
		Country:     "USA",
		Latitude:    37.7749,
		Longitude:   -122.4194,
		CreatedDate: time.Now(),
	})
	if err != nil {
		t.Errorf("Error creating new location: %s", err)
	}

	location.Name = "New Headquarters"

	router := gin.Default()
	router.PUT("/locations/:id", UpdateLocationHandler)

	locationJSON, err := json.Marshal(location)
	if err != nil {
		t.Errorf("Error marshalling location

