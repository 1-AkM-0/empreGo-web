package main

import (
	"errors"
	"net/http"

	"github.com/1-AkM-0/empreGo-web/internal/models"
	"github.com/1-AkM-0/empreGo-web/internal/pagination"
	"github.com/1-AkM-0/empreGo-web/internal/validator"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (app *application) createApplicationHandler(c *gin.Context) {
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "erro na requisição"})
		return
	}

	application := models.Application{
		UserID: userID,
		JobID:  input.JobID,
	}

	err := app.Models.ApplicationModel.Insert(&application)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "erro ao inserir application"})
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

	var input struct {
		pagination.Filter
	}

	v := validator.New()

	qs := c.Request.URL.Query()

	input.Page = app.readInt(qs, "page", 1, v)
	input.PageSize = app.readInt(qs, "page_size", 6, v)

	if pagination.ValidateFilter(v, input.Filter); !v.Valid() {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": v.Errors})
		return
	}

	applications, metadata, err := app.Models.ApplicationModel.GetAll(userID, input.Filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "erro ao retornar candidaturas"})
		return
	}

	c.JSON(http.StatusOK, envelope{"applications": applications, "metadata": metadata})

}

func (app *application) updateApplicationHandler(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID").(string)
	applicationID := c.Param("id")

	input := struct {
		Status string `json:"status"`
	}{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "erro na requisição"})
		return
	}

	err := app.Models.ApplicationModel.Update(userID, input.Status, applicationID)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrNoRecords):
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "candidatura não encontrada"})
			return

		default:
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "erro ao atualizar candidatura"})
			return
		}
	}
	c.Status(http.StatusNoContent)

}
