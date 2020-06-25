package main

import (
	"log"
	"os"

	cmd "github.com/Planxnx/discordBot-Golang/internal/cmd/discod-bot"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
	return
}
