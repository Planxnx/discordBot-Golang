package main

import (
	"log"
	"os"

	discordBot "github.com/Planxnx/discordBot-Golang/internal/cmd/discod-bot"
)

func main() {
	if err := discordBot.RunServer(); err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
	return
}
