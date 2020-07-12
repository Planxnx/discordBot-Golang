package cmd

import (
	"log"

	commandsProvider "github.com/Planxnx/discordBot-Golang/internal/commands/provider"
	"github.com/Planxnx/discordBot-Golang/internal/discord"
	"github.com/Planxnx/discordBot-Golang/internal/logger"
	messageProvider "github.com/Planxnx/discordBot-Golang/internal/messages/provider"
	musicProvider "github.com/Planxnx/discordBot-Golang/internal/music/provider"
	"github.com/Planxnx/discordBot-Golang/internal/routes"
	youtubeProvider "github.com/Planxnx/discordBot-Golang/internal/youtube/provider"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
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
		youtubeProvider.UsecaseModule,
		musicProvider.UsecaseModule,
		messageProvider.DeliveryModule,
		commandsProvider.DeliveryModule,
		routes.Module,
	)
	app.Run()

	return nil
}
