package app

import (
	"log"
	"net/http"

	"github.com/automation-as-a-service/internal/dto"
	"github.com/gin-gonic/gin"
)

func (m *MicroserviceServer) Login(c *gin.Context) {
	// Parse the user from the request body
	var person dto.Person
	if err := c.BindJSON(&person); err != nil {
		log.Printf("Could not bind request body to person due to %v\n", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := m.authService.SignIn(person.Email, person.Password)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.IndentedJSON(http.StatusOK, token)
}
