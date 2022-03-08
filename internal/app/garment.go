package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/automation-as-a-service/internal/dto"
	"github.com/gin-gonic/gin"
)

// Inserts a garment into the database
func (m *MicroserviceServer) CreateGarment(c *gin.Context) {
	// Parse the garment from the request body
	var garment dto.Garment
	if err := c.BindJSON(&garment); err != nil {
		log.Printf("Could not bind request body to garment due to %v", err)
	}

	// Insert the garment into the database
	garmentID, err := m.garmentService.CreateGarment(garment)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Return the new garment ID
	c.IndentedJSON(http.StatusCreated, gin.H{"id": garmentID})
}

// Gets a garment from the database
func (m *MicroserviceServer) GetGarment(c *gin.Context) {
	// Parse the garment ID from the request
	stringID := c.Param("id")
	requestID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		log.Printf("Could not convert param %s to int64 due to %v", stringID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Retrieve the garment from the database
	garment, err := m.garmentService.GetGarment(requestID)
	if err != nil {
		log.Printf("Could not get garment %d due to %v", requestID, err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Return the garment
	c.IndentedJSON(http.StatusOK, garment)
}

// Updates a garment in the database
func (m *MicroserviceServer) UpdateGarment(c *gin.Context) {
	// Parse the garment ID from the request
	stringID := c.Param("id")
	requestID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		log.Printf("Could not convert param %s to int64 due to %v", stringID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Parse the garment from the request body
	var garment dto.Garment
	if err := c.BindJSON(&garment); err != nil {
		log.Printf("Could not bind request body to garment due to %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	garment.ID = requestID

	// Update the garment in the database
	updatedGarment, err := m.garmentService.UpdateGarment(garment)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Return updated garment
	c.IndentedJSON(http.StatusCreated, updatedGarment)
}

// Deletes a garment from the database
func (m *MicroserviceServer) DeleteGarment(c *gin.Context) {
	// Parse the garment ID from the request
	stringID := c.Param("id")
	requestID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		log.Printf("Could not convert param %s to int64 due to %v", stringID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Delete the garment from the database
	err = m.garmentService.DeleteGarment(requestID)
	if err != nil {
		log.Printf("Could not delete garment %d due to %v", requestID, err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Return a successful status code
	c.Status(http.StatusNoContent)
}
