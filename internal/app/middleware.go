package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (m *MicroserviceServer) getUserIdFromToken(token string) (int64, error) {
	userID, err := m.tokenManager.Parse(token)
	if err != nil {
		return 0, err
	}

	return *userID, nil
}

func (m *MicroserviceServer) AuthorizeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		userID, err := m.getUserIdFromToken(auth)
		if err != nil {
			log.Printf("Could not read User ID from Authorization header %s due to %v\n", auth, err)
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		user, err := m.userService.GetUserSingle(userID)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
