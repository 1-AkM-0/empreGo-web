package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session *discordgo.Session
}

func (b *Bot) SendMessage(channelID, content string) (*discordgo.Message, error) {
	return b.session.ChannelMessageSend(channelID, content)
}

func NewBot(token string) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("discord NewBot: %w", err)
	}
	session.Identify.Intents = discordgo.IntentGuildMessages

	err = session.Open()
	if err != nil {
		return nil, fmt.Errorf("discord Open: %w", err)
	}

	return &Bot{session: session}, nil

}
