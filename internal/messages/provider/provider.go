package provider

import (
	"github.com/Planxnx/discordBot-Golang/internal/messages/delivery"
	"go.uber.org/fx"
)

//DeliveryModule .
var DeliveryModule = fx.Options(
	fx.Provide(delivery.NewMessageDelivery),
)
