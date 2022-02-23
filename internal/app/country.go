package app

import (
	"log"
	"net/http"

	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/gin-gonic/gin"
)

// Gets a country from the database
func (m *MicroserviceServer) GetCountry(c *gin.Context) {
	countryName := c.Param("name")

	country, err := m.countryService.GetCountryByName(countryName)
	if err != nil {
		log.Printf("Could not get country %s due to %v", countryName, err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, country)
}

func (m *MicroserviceServer) CreateCountry(c *gin.Context) {
	var country datastruct.Country
	if err := c.BindJSON(&country); err != nil {
		log.Printf("Could not bind request body to person due to %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	auth := c.Request.Header.Get("Authorization")
	userID, err := m.getUserIdFromToken(auth)
	if err != nil {
		log.Printf("Could not read User ID from Authorization header %s due to %v\n", auth, err)
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	countryID, err := m.countryService.CreateCountry(country, userID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"id": countryID})
}

func (m *MicroserviceServer) DeleteCountry(c *gin.Context) {
	// Parse the userID from the token
	auth := c.Request.Header.Get("Authorization")
	userID, err := m.getUserIdFromToken(auth)
	if err != nil {
		log.Printf("Could not read User ID from Authorization header %s due to %v\n", auth, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	countryName := c.Param("name")

	err = m.countryService.DeleteCountry(countryName, userID)
	if err != nil {
		log.Printf("Could not delete country %s due to %v", countryName, err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.Status(http.StatusNoContent)
}
