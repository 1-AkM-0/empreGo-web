package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/1-AkM-0/empreGo-web/internal/pagination"
	"github.com/1-AkM-0/empreGo-web/internal/validator"
	"github.com/gin-gonic/gin"
)

func (app *application) getJobByIDHandler(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Job Id inválido",
		})
		return
	}
	job, err := app.Models.JobModel.GetJobByID(id)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Job não encontrado",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, job)

}

func (app *application) getAllJobsHandler(c *gin.Context) {

	userID := app.tryExtractUserID(c)

	var input struct {
		pagination.Filter
	}

	v := validator.New()

	qs := c.Request.URL.Query()

	input.Page = app.readInt(qs, "page", 1, v)
	input.PageSize = app.readInt(qs, "page_size", 6, v)

	jobs, metadata, err := app.Models.JobModel.GetJobs(input.Filter, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if pagination.ValidateFilter(v, input.Filter); !v.Valid() {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": v.Errors})
		return
	}

	c.JSON(http.StatusOK, envelope{"jobs": jobs, "metadata": metadata})
}
