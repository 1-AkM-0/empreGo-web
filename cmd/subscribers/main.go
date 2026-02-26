package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/1-AkM-0/empreGo-web/internal/discord"
	"github.com/nats-io/nats.go"
)

func main() {
	url := nats.DefaultURL

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	bot, err := discord.NewBot(os.Getenv("BOT_TOKEN"))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer func() {
		if closeErr := bot.Close(); closeErr != nil {
			logger.Error("erro ao fechar o bot", "detalhes", closeErr)
		}
	}()

	nc, err := nats.Connect(url)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer func() {
		if drainErr := nc.Drain(); drainErr != nil {
			logger.Error("erro ao fazer drain das vagas", "detalhes", drainErr)
		}
	}()

	discordChannels := map[string]string{
		"vagas.fullstack": "1475567878240473129",
		"vagas.backend":   "1475567813149200394",
		"vagas.frontend":  "1475567850541289595",
		"vagas.geral":     "1476258634047688826"}

	for topic, channelID := range discordChannels {
		_, err = nc.Subscribe(topic, func(msg *nats.Msg) {
			_, err := bot.SendMessage(channelID, string(msg.Data))
			if err != nil {
				logger.Error("erro ao enviar vaga", "error", err.Error())
			}
		})
		if err != nil {
			logger.Error("erro ao inscrever-se no canal", "error", err.Error())
		}
	}

	fmt.Println("Listening on [vagas.*]")

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan

	fmt.Println("desligando subscriber")
}
