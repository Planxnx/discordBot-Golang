package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	botToken string
)

func init() {
	botToken = os.Getenv("BOT_TOKEN")
	if botToken == "" {
		botToken = "TOKEN NOT FOUND!!"
	}
}

func main() {

	fmt.Println("Discord Session is starting with token '", botToken, "'")

	discordSession, err := discordgo.New("Bot " + botToken)
	if err != nil {
		fmt.Println("Error: creating Discord session,\nMsg: ", err)
		return
	}

	err = discordSession.Open()
	if err != nil {
		fmt.Println("Error: opening connection,\nMsg: ", err)
		return
	}

	fmt.Println("Discord Bot is now running, Press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGINT)
	<-sc

	discordSession.Close()
	fmt.Println("close down the Discord session")
}
