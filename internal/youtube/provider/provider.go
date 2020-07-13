package provider

import (
	"github.com/Planxnx/discordBot-Golang/internal/youtube/usecase"
	"go.uber.org/fx"
)

//UsecaseModule .
var UsecaseModule = fx.Options(
	fx.Provide(usecase.NewYoutubeUsecase),
)
