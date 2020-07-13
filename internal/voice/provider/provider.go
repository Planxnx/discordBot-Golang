package provider

import (
	"github.com/Planxnx/discordBot-Golang/internal/voice/usecase"
	"go.uber.org/fx"
)

//UsecaseModule .
var UsecaseModule = fx.Options(
	fx.Provide(usecase.NewVoiceUsecase),
)
