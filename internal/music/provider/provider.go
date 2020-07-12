package provider

import (
	"github.com/Planxnx/discordBot-Golang/internal/music/usecase"
	"go.uber.org/fx"
)

//UsecaseModule .
var UsecaseModule = fx.Options(
	fx.Provide(usecase.NewMusicUsecase),
)
