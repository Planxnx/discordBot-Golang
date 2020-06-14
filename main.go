package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Planxnx/discordBot-Golang/services"
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

	discordSession.AddHandler(msgHandleService)

	log.Println("Discord Bot is now running, Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGINT)
	<-sc

	discordSession.Close()
	log.Println("close down the Discord session")
}

func msgHandleService(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	botPrefix := os.Getenv("BOT_PREFIX")
	if botPrefix == "" {
		botPrefix = "~"
	}

	if strings.HasPrefix(m.Content, botPrefix) {
		go services.CommandService(s, m, botPrefix)
	}
	go services.MessageService(s, m)
}
