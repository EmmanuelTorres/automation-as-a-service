package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
	"github.com/gin-gonic/gin"
)

// Inserts a country into the database
func (m *MicroserviceServer) CreateCountry(c *gin.Context) {
	// Parse the country from the request body
	var country datastruct.Country
	if err := c.BindJSON(&country); err != nil {
		log.Printf("Could not bind request body to person due to %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Insert the country into the database
	countryID, err := m.countryService.CreateCountry(country)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Return the new country ID
	c.IndentedJSON(http.StatusCreated, gin.H{"id": countryID})
}

// Gets a country from the database
func (m *MicroserviceServer) GetCountry(c *gin.Context) {
	// Parse the country ID from the request
	stringID := c.Param("id")
	requestID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		log.Printf("Could not convert param %s to int64 due to %v\n", stringID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Retrieve the country from the database
	country, err := m.countryService.GetCountry(requestID)
	if err != nil {
		log.Printf("Could not get country %d due to %v", requestID, err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Return the country
	c.IndentedJSON(http.StatusOK, country)
}

// Updates a country in the database
func (m *MicroserviceServer) UpdateCountry(c *gin.Context) {
	// Parse the country ID from the request
	stringID := c.Param("id")
	requestID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		log.Printf("Could not convert param %s to int64 due to %v\n", stringID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Parse the country from the request body
	var country dto.Country
	if err := c.BindJSON(&country); err != nil {
		log.Printf("Could not bind request body to person due to %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	country.ID = requestID

	// Insert the country into the database
	updatedCountry, err := m.countryService.UpdateCountry(country)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Return the new country ID
	c.IndentedJSON(http.StatusCreated, updatedCountry)
}

// Deletes a country from the database
func (m *MicroserviceServer) DeleteCountry(c *gin.Context) {
	// Parse the country ID from the request
	stringID := c.Param("id")
	requestID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		log.Printf("Could not convert param %s to int64 due to %v\n", stringID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Delete the country from the database
	err = m.countryService.DeleteCountry(requestID)
	if err != nil {
		log.Printf("Could not delete country %d due to %v", requestID, err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	// Return a successful status code
	c.Status(http.StatusNoContent)
}
