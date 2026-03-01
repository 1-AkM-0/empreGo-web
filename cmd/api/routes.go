package main

import (
	"github.com/1-AkM-0/empreGo-web/internal/auth"
	"github.com/1-AkM-0/empreGo-web/internal/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {
	store := auth.Setup(app.Config.githubClientID, app.Config.githubClientSecret, app.Config.sessionSecret)
	router := gin.Default()

	router.Use(sessions.Sessions("session", store))

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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
			auth.POST("/logout", middleware.RequireAuth(app.Models.UserModel), app.logoutHandler)
			auth.GET("/me", app.getMeHandler)
		}
		application := v1.Group("/applications")
		{
			application.GET("/", middleware.RequireAuth(app.Models.UserModel), app.getApplicationsHandler)
			application.GET("/:id")
			application.POST("/", middleware.RequireAuth(app.Models.UserModel), app.createApplicationHandler)
			application.PATCH("/:id", middleware.RequireAuth(app.Models.UserModel), app.updateApplicationHandler)
			application.DELETE("/:id")

		}
	}
	return router

}
