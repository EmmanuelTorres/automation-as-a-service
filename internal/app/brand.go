package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/automation-as-a-service/internal/dto"
	"github.com/gin-gonic/gin"
)

// Inserts a brand into the database
func (m *MicroserviceServer) CreateBrand(c *gin.Context) {
	// Parse the brand from the request body
	var brand dto.Brand
	if err := c.BindJSON(&brand); err != nil {
		log.Printf("Could not bind request body to brand due to %v", err)
	}

	// Insert the brand into the database
	brandID, err := m.brandService.CreateBrand(brand)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Return the new brand ID
	c.IndentedJSON(http.StatusCreated, gin.H{"id": brandID})
}

// Gets a brand from the database
func (m *MicroserviceServer) GetBrand(c *gin.Context) {
	// Parse the brand ID from the request
	stringID := c.Param("id")
	requestID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		log.Printf("Could not convert param %s to int64 due to %v", stringID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Retrieve the brand from the database
	brand, err := m.brandService.GetBrand(requestID)
	if err != nil {
		log.Printf("Could not get brand %d due to %v", requestID, err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Return the brand
	c.IndentedJSON(http.StatusOK, brand)
}

// Updates a brand in the database
func (m *MicroserviceServer) UpdateBrand(c *gin.Context) {
	// Parse the brand ID from the request
	stringID := c.Param("id")
	requestID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		log.Printf("Could not convert param %s to int64 due to %v", stringID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Parse the brand from the request body
	var brand dto.Brand
	if err := c.BindJSON(&brand); err != nil {
		log.Printf("Could not bind request body to brand due to %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	brand.ID = requestID

	// Update the brand in the database
	updatedBrand, err := m.brandService.UpdateBrand(brand)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Return updated brand
	c.IndentedJSON(http.StatusCreated, updatedBrand)
}

// Deletes a brand from the database
func (m *MicroserviceServer) DeleteBrand(c *gin.Context) {
	// Parse the brand ID from the request
	stringID := c.Param("id")
	requestID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		log.Printf("Could not convert param %s to int64 due to %v", stringID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Delete the brand from the database
	err = m.brandService.DeleteBrand(requestID)
	if err != nil {
		log.Printf("Could not delete brand %d due to %v", requestID, err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Return a successful status code
	c.Status(http.StatusNoContent)
}
