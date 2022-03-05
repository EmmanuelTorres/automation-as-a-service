package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/automation-as-a-service/internal/dto"
	"github.com/gin-gonic/gin"
)

// Inserts a designer into the database
func (m *MicroserviceServer) CreateDesigner(c *gin.Context) {
	// Parse the designer from the request body
	var designer dto.Designer
	if err := c.BindJSON(&designer); err != nil {
		log.Printf("Could not bind request body to designer due to %v", err)
	}

	// Insert the designer into the database
	designerID, err := m.designerService.CreateDesigner(designer)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Return the new designer ID
	c.IndentedJSON(http.StatusCreated, gin.H{"id": designerID})
}

// Gets a designer from the database
func (m *MicroserviceServer) GetDesigner(c *gin.Context) {
	// Parse the designer ID from the request
	stringID := c.Param("id")
	requestID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		log.Printf("Could not convert param %s to int64 due to %v", stringID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Retrieve the designer from the database
	designer, err := m.designerService.GetDesigner(requestID)
	if err != nil {
		log.Printf("Could not get designer %d due to %v", requestID, err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Return the designer
	c.IndentedJSON(http.StatusOK, designer)
}

// Updates a designer in the database
func (m *MicroserviceServer) UpdateDesigner(c *gin.Context) {
	// Parse the designer ID from the request
	stringID := c.Param("id")
	requestID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		log.Printf("Could not convert param %s to int64 due to %v", stringID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Parse the designer from the request body
	var designer dto.Designer
	if err := c.BindJSON(&designer); err != nil {
		log.Printf("Could not bind request body to designer due to %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	designer.ID = requestID

	// Update the designer in the database
	updatedDesigner, err := m.designerService.UpdateDesigner(designer)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Return updated designer
	c.IndentedJSON(http.StatusCreated, updatedDesigner)
}

// Deletes a designer from the database
func (m *MicroserviceServer) DeleteDesigner(c *gin.Context) {
	// Parse the designer ID from the request
	stringID := c.Param("id")
	requestID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		log.Printf("Could not convert param %s to int64 due to %v", stringID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Delete the designer from the database
	err = m.designerService.DeleteDesigner(requestID)
	if err != nil {
		log.Printf("Could not delete designer %d due to %v", requestID, err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Return a successful status code
	c.Status(http.StatusNoContent)
}
