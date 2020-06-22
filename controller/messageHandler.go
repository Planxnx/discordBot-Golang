package controller

import (
	"os"
	"strings"

	"github.com/Planxnx/discordBot-Golang/services"
	"github.com/bwmarrin/discordgo"
)

//MessageHandler Controller
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	botPrefix := os.Getenv("BOT_PREFIX")
	if botPrefix == "" {
		botPrefix = "~"
	}

	if strings.HasPrefix(m.Content, botPrefix) {
		go CommandService(s, m, botPrefix)
	}
	go services.MessageService(s, m)
}
