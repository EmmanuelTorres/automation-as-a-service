package app

import (
	"log"
	"net/http"

	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/gin-gonic/gin"
)

func (m *MicroserviceServer) CreateProject(c *gin.Context) {
	var project datastruct.Project
	if err := c.BindJSON(&project); err != nil {
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

	projectID, err := m.projectService.CreateProject(project, userID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"id": projectID})
}

func (m *MicroserviceServer) GetProject(c *gin.Context) {
	name := c.Param("name")
	project, err := m.projectService.GetProject(name)
	if err != nil {
		log.Printf("Could not get project %s due to %v", name, err)
		c.AbortWithError(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, project)
}
