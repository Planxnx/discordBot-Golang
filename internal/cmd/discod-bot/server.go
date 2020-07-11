package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Planxnx/discordBot-Golang/internal/commands/controller"
	commandsProvider "github.com/Planxnx/discordBot-Golang/internal/commands/provider"
	"github.com/Planxnx/discordBot-Golang/internal/discord"
	"github.com/Planxnx/discordBot-Golang/internal/logger"
	messagesController "github.com/Planxnx/discordBot-Golang/internal/messages/controller"
	messageProvider "github.com/Planxnx/discordBot-Golang/internal/messages/provider"
	"github.com/Planxnx/discordBot-Golang/internal/routes"
	"github.com/joho/godotenv"
	"go.uber.org/fx"

	"github.com/bwmarrin/discordgo"
)

var (
	botToken string
)

// RunServer runs discord bot server
func RunServer() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("dotEnv: can't loading .env file")
	}

	app := fx.New(
		fx.Provide(logger.NewLogger),
		fx.Provide(discord.NewSession),
		messageProvider.DeliveryModule,
		commandsProvider.DeliveryModule,
		routes.Module,
	)
	app.Run()

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
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println(err)
	}
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println(err)
	}

	if strings.HasPrefix(m.Content, botPrefix) {
		go controller.CommandHandler(s, m, guild, botPrefix)
	}
	go messagesController.MessageHandler(s, m, guild)
}
