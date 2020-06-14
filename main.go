package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Planxnx/discordBot-Golang/messages"
	"github.com/bwmarrin/discordgo"
)

var (
	botToken string
)

func init() {
	botToken = os.Getenv("BOT_TOKEN")
	if botToken == "" {
		fmt.Println("BOT_TOKEN not found, Closing Now...")
		os.Exit(0)
	}
}

func main() {

	log.Println("Discord Session is starting with token '", botToken, "'")

	discordSession, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Println("Error: creating Discord session,\nMsg: ", err)
		return
	}

	err = discordSession.Open()
	if err != nil {
		log.Println("Error: opening connection,\nMsg: ", err)
		return
	}

	discordSession.AddHandler(messages.HandleService)

	log.Println("Discord Bot is now running, Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGINT)
	<-sc

	discordSession.Close()
	log.Println("close down the Discord session")
}
