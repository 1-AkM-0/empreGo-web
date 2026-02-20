package main

import "github.com/gin-gonic/gin"

func (app *application) routes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/healthcheck", app.healthCheckHandler)
		jobs := v1.Group("/jobs")
		{
			jobs.GET("/", app.getAllJobsHandler)
			jobs.GET("/:id", app.getJobByIDHandler)
		}
		auth := v1.Group("/auth")
		{
			auth.POST("/register")
			auth.POST("/login")
		}
		application := v1.Group("/application")
		{
			application.GET("/")
			application.GET("/:id")
			application.POST("/")
			application.PATCH("/:id")
			application.DELETE("/:id")

		}
	}
	return router

}
