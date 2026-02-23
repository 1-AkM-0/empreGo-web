package main

import (
	"github.com/1-AkM-0/empreGo-web/internal/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	store := auth.Setup(app.Config.githubClientID, app.Config.githubClientSecret, app.Config.sessionSecret)
	router := gin.Default()
	router.Use(sessions.Sessions("session", store))
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
			auth.GET("/:provider/callback", app.authCallbackHandler)
			auth.GET("/:provider", app.authBeginHandler)
			auth.POST("/logout", app.logoutHandler)
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
