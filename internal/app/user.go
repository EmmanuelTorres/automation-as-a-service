package app

import (
	"log"
	"net/http"
	"strconv"

	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/gin-gonic/gin"
)

// swagger:route GET /users users GetUsers
// Return a list of users from the database
// responses:
//	200: usersResponse

// Gets all users from the database
func (m *MicroserviceServer) GetUsers(c *gin.Context) {
	// Retrieve the users from the database
	users, err := m.userService.GetUsers()
	if err != nil {
		log.Printf("Could not get users due to %v\n", err)
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	// Return the users
	c.IndentedJSON(http.StatusOK, gin.H{"users": users})
}

// swagger:route GET /users/{username} users GetUser
// Returns a single user from the database
// responses:
//	200: userResponse

// Gets a user from the database
func (m *MicroserviceServer) GetUser(c *gin.Context) {
	// Retrieve the user from the middleware
	tokenUser := c.MustGet(USER_KEY).(*datastruct.Person)

	// Parse the user ID from the request
	stringID := c.Param("username")
	requestID, err := strconv.ParseInt(stringID, 10, 64)
	if err != nil {
		log.Printf("Could not convert param %s to int64 due to %v\n", stringID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Retrieve the user from the database
	user, err := m.userService.GetUser(requestID, tokenUser)
	if err != nil {
		log.Printf("Could not get user %d due to %v\n", requestID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Return the requested user
	c.IndentedJSON(http.StatusOK, gin.H{"user": user})
}

// swagger:route POST /users/{username} users CreateUser
// Returns the id of a newly created user
// responses:
//	201: objectCreatedResponse

// Creates a user in the database
func (m *MicroserviceServer) CreateUser(c *gin.Context) {
	// Parse the user from the request body
	var user datastruct.Person
	if err := c.BindJSON(&user); err != nil {
		log.Printf("Could not bind request body to user due to %v\n", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Create the user in the database
	id, err := m.authService.SignUp(user)
	if err != nil {
		log.Printf("Could not create user %v due to %v\n", user, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Return the new user ID
	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})
}

// swagger:route DELETE /users/{username} users DeleteUser
// Deletes the user from the database
// responses:
//	204: noContentResponse

// Deletes a user from the database
func (m *MicroserviceServer) DeleteUser(c *gin.Context) {
	// Retrieve the user from the middleware
	tokenUser := c.MustGet("user").(*datastruct.Person)

	// Parse the user ID from the request
	requestID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Could not convert param id to int64 due to %v\n", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Delete the user from the database
	if err := m.userService.DeleteUser(requestID, tokenUser); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Return an acknowledgement that the deletion has taken place
	c.Status(http.StatusNoContent)
}
