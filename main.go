package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Planxnx/discordBot-Golang/controller"
	"github.com/Planxnx/discordBot-Golang/data"

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
	var err error
	log.Println("Discord Session is starting with token '", botToken, "'")

	data.DiscordSession, err = discordgo.New("Bot " + botToken)
	if err != nil {
		log.Println("Error: creating Discord session,\nMsg: ", err)
		return
	}

	err = data.DiscordSession.Open()
	if err != nil {
		log.Println("Error: opening connection,\nMsg: ", err)
		return
	}

	data.DiscordSession.AddHandler(controller.MessageHandler)

	log.Println("Discord Bot is now running, Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGINT)
	<-sc

	data.DiscordSession.Close()
	log.Println("close down the Discord session")
}
