package main

import (
	"fmt"
	"log/slog"
	"os"
	"sync"

	"github.com/1-AkM-0/empreGo-web/internal/discord"
	"github.com/1-AkM-0/empreGo-web/internal/models"
	"github.com/1-AkM-0/empreGo-web/internal/scraper"
	"github.com/1-AkM-0/empreGo-web/internal/storage"
	"github.com/nats-io/nats.go"
)

type application struct {
	Bot      discord.Bot
	Logger   *slog.Logger
	JobModel models.JobModel
	Nc       *nats.Conn
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	bot, err := discord.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	db, err := storage.Open()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	natsUrl := nats.DefaultURL
	nc, err := nats.Connect(natsUrl)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		Bot:      *bot,
		Logger:   logger,
		JobModel: models.JobModel{DB: db},
		Nc:       nc,
	}

	app.run()
}

func (app *application) run() {
	app.Logger.Info("Iniciando busca")

	var wg sync.WaitGroup
	jobChannel := make(chan models.Job, 10)

	sources := []func(jobChannel chan models.Job) error{
		scraper.SearchLinkedin,
		scraper.SearchGupy,
	}

	for _, search := range sources {
		wg.Go(func() {
			err := search(jobChannel)
			if err != nil {
				app.Logger.Error(err.Error())
			}
		})
	}

	go func() {
		wg.Wait()
		close(jobChannel)
	}()
	count := 0

	for job := range jobChannel {
		if app.JobModel.Exists(job.Link) {
			continue
		}

		msg := []byte("Nova vaga: " + job.Title + "\n" + "Fonte: " + job.Source + "\n" + job.Link)
		subject := "vagas." + job.Type

		err := app.Nc.Publish(subject, msg)
		if err != nil {
			app.Logger.Error("erro ao tentar enviar vaga: " + err.Error())
			continue
		}

		err = app.JobModel.Insert(&job)
		if err != nil {
			app.Logger.Error(err.Error())
		}
		count++
	}
	result := fmt.Sprintf("Foram encontradas %d vagas", count)
	app.Logger.Info(result)
}
