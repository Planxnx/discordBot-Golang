package cmd

import (
	"context"
	"log"

	commandsProvider "github.com/Planxnx/discordBot-Golang/internal/commands/provider"
	"github.com/Planxnx/discordBot-Golang/internal/discord"
	"github.com/Planxnx/discordBot-Golang/internal/logger"
	messageProvider "github.com/Planxnx/discordBot-Golang/internal/messages/provider"
	musicProvider "github.com/Planxnx/discordBot-Golang/internal/music/provider"
	voiceProvider "github.com/Planxnx/discordBot-Golang/internal/voice/provider"

	"github.com/Planxnx/discordBot-Golang/internal/routes"
	youtubeProvider "github.com/Planxnx/discordBot-Golang/internal/youtube/provider"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var (
	botToken string
)

func registerHooks(lifecycle fx.Lifecycle, discord discord.Discord) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				log.Print("Starting server.")
				if err := discord.OpenConnection(); err != nil {
					log.Printf("%v\n", err)
				}
				return nil
			},
			OnStop: func(context.Context) error {
				log.Print("Stopping server.")
				if err := discord.CloseConnection(); err != nil {
					log.Printf("%v\n", err)
				}
				return nil
			},
		},
	)
}

// RunServer runs discord bot server
func RunServer() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("dotEnv: can't loading .env file")
	}

	app := fx.New(
		fx.Provide(logger.NewLogger),
		fx.Provide(discord.NewSession),
		fx.Invoke(registerHooks),
		messageProvider.RepositoryModule,
		messageProvider.UsecaseModule,
		voiceProvider.UsecaseModule,
		youtubeProvider.UsecaseModule,
		musicProvider.UsecaseModule,
		messageProvider.DeliveryModule,
		commandsProvider.DeliveryModule,
		routes.Module,
	)
	app.Run()

	return nil
}
