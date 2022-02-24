package app

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/automation-as-a-service/internal/datastruct"
	"github.com/automation-as-a-service/internal/dto"
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

func (m *MicroserviceServer) UpdateProject(c *gin.Context) {
	paramID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Could not parse id due to %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var requestProject datastruct.Project
	if err := c.BindJSON(&requestProject); err != nil {
		log.Printf("Could not bind request body to person due to %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if paramID != requestProject.ID {
		log.Printf("Could not match paramID %d and bodyID %d", paramID, requestProject.ID)
		c.AbortWithError(http.StatusBadRequest, errors.New("the project uri id does not match the request id"))
		return
	}

	project := dto.Project{
		ID:   requestProject.ID,
		Name: requestProject.Name,
	}

	updatedProject, err := m.projectService.UpdateProject(project)
	if err != nil {
		log.Printf("Could not update project %v due to %v", project, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, updatedProject)
}

func (m *MicroserviceServer) DeleteProject(c *gin.Context) {
	paramID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		log.Printf("Could not parse param id due to %v", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	authorizedUser := c.MustGet("user").(*datastruct.Person)

	err = m.projectService.DeleteProject(paramID, authorizedUser)
	if err != nil {
		log.Printf("Could not delete project %d due to %v", paramID, err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
