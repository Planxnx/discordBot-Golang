package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Planxnx/discordBot-Golang/controller"
	"github.com/Planxnx/discordBot-Golang/data"
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
	data.DiscordSession, err = discordgo.New("Bot " + botToken)
	if err != nil {
		return fmt.Errorf("Error: creating Discord session, Message: '%s'", err)
	}

	err = data.DiscordSession.Open()
	if err != nil {
		return fmt.Errorf("Error: opening connection, Message: '%s'", err)
	}

	data.DiscordSession.AddHandler(controller.MessageHandler)

	log.Println("Discord Bot is now running, Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGINT)
	<-sc

	data.DiscordSession.Close()
	log.Println("close down the Discord session")
	return nil
}
