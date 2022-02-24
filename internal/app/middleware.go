package app

import (
	"log"
	"net/http"

	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/gin-gonic/gin"
)

func (m *MicroserviceServer) getUserIdFromToken(token string) (int64, error) {
	userID, err := m.tokenManager.Parse(token)
	if err != nil {
		return 0, err
	}

	return *userID, nil
}

// Verifies that the user from the token is a user
func (m *MicroserviceServer) AuthorizeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the user from the token
		auth := c.Request.Header.Get("Authorization")
		userID, err := m.getUserIdFromToken(auth)
		if err != nil {
			log.Printf("Could not read User ID from Authorization header %s due to %v\n", auth, err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Parse the user from the database
		user, err := m.userService.GetUserSingle(userID)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

// Verifies that the user from the token is an admin
func (m *MicroserviceServer) AuthorizeAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the user from the token
		auth := c.Request.Header.Get("Authorization")
		userID, err := m.getUserIdFromToken(auth)
		if err != nil {
			log.Printf("Could not read User ID from Authorization header %s due to %v\n", auth, err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Parse the user from the database
		user, err := m.userService.GetUserSingle(userID)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		// Verify the user is an admin
		if user.Role != datastruct.ADMIN {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
