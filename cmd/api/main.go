package main

import (
	"flag"

	"github.com/1-AkM-0/empreGo-web/internal/models"
	"github.com/1-AkM-0/empreGo-web/internal/storage"
)

type application struct {
	Models models.Models
}

func main() {
	addr := flag.String("addr", ":8080", "server port")
	flag.Parse()

	db, err := storage.Open()
	if err != nil {
		panic(err)
	}

	app := application{
		Models: models.NewModels(db),
	}

	router := app.routes()

	router.Run(*addr)

}
