package routes

import (
	commandsDelivery "github.com/Planxnx/discordBot-Golang/internal/commands/delivery"
	"github.com/Planxnx/discordBot-Golang/internal/discord"
	messageDelivery "github.com/Planxnx/discordBot-Golang/internal/messages/delivery"
	"go.uber.org/fx"
)

//NewRoutes new Routes Handler
func NewRoutes(discord discord.Discord, messageDelivery messageDelivery.Delivery, commandsDelivery commandsDelivery.Delivery) {
	discord.AddHandler(messageDelivery.GetMessageHandler)
	discord.AddHandler(commandsDelivery.GetCommandsHandler)
}

//Module .
var Module = fx.Options(
	fx.Invoke(NewRoutes),
)
