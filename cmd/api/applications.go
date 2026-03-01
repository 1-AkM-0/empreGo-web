package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/1-AkM-0/empreGo-web/internal/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (app *application) createApplicationHandler(c *gin.Context) {
	fmt.Println("chegou no handler")
	session := sessions.Default(c)

	input := struct {
		JobID int `json:"job_id"`
	}{}

	userID, ok := session.Get("userID").(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "erro na requisição: " + err.Error()})
		return
	}

	application := models.Application{
		UserID: userID,
		JobID:  input.JobID,
	}

	err := app.Models.ApplicationModel.Insert(&application)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, application)
}

func (app *application) getApplicationsHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID, ok := session.Get("userID").(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	application, err := app.Models.ApplicationModel.GetAll(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "erro ao retornar candidaturas"})
		return
	}

	c.JSON(http.StatusOK, application)

}

func (app *application) updateApplicationHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID").(string)
	applicationID := c.Param("id")

	input := struct {
		Status string `json:"status"`
	}{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "erro na requisição: " + err.Error()})
		return
	}

	err := app.Models.ApplicationModel.Update(userID, input.Status, applicationID)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecords):
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "candidatura não encontrada"})
			return

		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "erro: " + err.Error()})
			return
		}
	}
	c.Status(http.StatusNoContent)

}
