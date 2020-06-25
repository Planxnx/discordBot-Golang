package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Planxnx/discordBot-Golang/internal/commands/controller"
	"github.com/Planxnx/discordBot-Golang/internal/discord"
	messagesController "github.com/Planxnx/discordBot-Golang/internal/messages/controller"
	"github.com/joho/godotenv"

	"github.com/bwmarrin/discordgo"
)

var (
	botToken string
)

// RunServer runs discord bot server
func RunServer() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Error: can't loading .env file")
	}

	botToken = os.Getenv("BOT_TOKEN")
	if botToken == "" {
		return fmt.Errorf("Error: BOT_TOKEN not found, Closing now")
	}

	log.Println("Discord Session is starting with token '", botToken, "'")
	discord.Session, err = discordgo.New("Bot " + botToken)
	if err != nil {
		return fmt.Errorf("Error: creating Discord session, Message: '%s'", err)
	}

	err = discord.Session.Open()
	if err != nil {
		return fmt.Errorf("Error: opening connection, Message: '%s'", err)
	}

	discord.Session.AddHandler(messageHandler)

	log.Println("Discord Bot is now running, Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGINT)
	<-sc

	discord.Session.Close()
	log.Println("close down the Discord session")
	return nil
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	botPrefix := os.Getenv("BOT_PREFIX")
	if botPrefix == "" {
		botPrefix = "~"
	}

	if strings.HasPrefix(m.Content, botPrefix) {
		go controller.CommandHandler(s, m, botPrefix)
	}
	go messagesController.MessageHandler(s, m)
}
