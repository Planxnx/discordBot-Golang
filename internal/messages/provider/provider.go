package provider

import (
	"github.com/Planxnx/discordBot-Golang/internal/messages/delivery"
	"github.com/Planxnx/discordBot-Golang/internal/messages/repository"
	"github.com/Planxnx/discordBot-Golang/internal/messages/usecase"
	"go.uber.org/fx"
)

//DeliveryModule .
var DeliveryModule = fx.Options(
	fx.Provide(delivery.NewMessageDelivery),
)

//RepositoryModule .
var RepositoryModule = fx.Options(
	fx.Provide(repository.NewMessageRepository),
)

//UsecaseModule .
var UsecaseModule = fx.Options(
	fx.Provide(usecase.NewMessagesUsecase),
)
