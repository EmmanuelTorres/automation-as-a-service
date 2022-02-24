package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
	"github.com/gin-gonic/gin"
)

// Gets a user from the database
func (m *MicroserviceServer) GetUser(c *gin.Context) {
	// Parse the user from the middleware
	tokenUser := c.MustGet("user").(*datastruct.Person)

	// Parse the userID from the token
	auth := c.Request.Header.Get("Authorization")
	tokenID, err := m.getUserIdFromToken(auth)
	if err != nil {
		log.Printf("Could not read User ID from Authorization header %s due to %v\n", auth, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Parse the userID from the request
	requestID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Could not convert param id to int64 due to %v\n", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Retrieve the datastruct user from the database
	user, err := m.userService.GetUser(requestID, tokenID)
	if err != nil {
		log.Printf("Could not get user %d due to %v\n", requestID, err)
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// Return the datastruct user
	c.IndentedJSON(http.StatusOK, user)
}

// Creates a user in the database
func (m *MicroserviceServer) CreateUser(c *gin.Context) {
	// Parse the user from the request body
	var person dto.Person
	if err := c.BindJSON(&person); err != nil {
		log.Printf("Could not bind request body to person due to %v\n", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user := datastruct.Person{
		FirstName:   person.FirstName,
		LastName:    person.LastName,
		Email:       person.Email,
		Password:    person.Password,
		PhoneNumber: person.PhoneNumber,
		Role:        "user",
	}

	// Create the user in the database
	id, err := m.authService.SignUp(user)
	if err != nil {
		log.Printf("Could not create user %v due to %v\n", person, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Return the created status code
	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

// Deletes a user in the database
func (m *MicroserviceServer) DeleteUser(c *gin.Context) {
	// Parse the userID from the token
	auth := c.Request.Header.Get("Authorization")
	tokenID, err := m.getUserIdFromToken(auth)
	if err != nil {
		log.Printf("Could not read User ID from Authorization header %s due to %v\n", auth, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Parse the userID from the request
	requestID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Could not convert param id to int64 due to %v\n", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := m.userService.DeleteUser(requestID, tokenID); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
