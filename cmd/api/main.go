package main

import (
	"os"

	"github.com/1-AkM-0/empreGo-web/internal/auth"
	"github.com/1-AkM-0/empreGo-web/internal/models"
	"github.com/1-AkM-0/empreGo-web/internal/storage"
)

type config struct {
	githubClientID     string
	githubClientSecret string
	sessionSecret      string
}

type application struct {
	Config       config
	Models       models.Models
	TokenManager auth.TokenManager
}

func main() {
	cfg := config{
		githubClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		githubClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		sessionSecret:      os.Getenv("SESSION_SECRET"),
	}

	db, err := storage.Open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := application{
		Config: cfg,
		Models: models.NewModels(db),
	}

	router := app.routes()

	router.Run(":8080")

}
