package data

import (
	"github.com/Planxnx/discordBot-Golang/model"
	"github.com/bwmarrin/discordgo"
)

var (
	DiscordSession *discordgo.Session
	BotStatus      bool
	NowPlaying     model.Song
	PlayList       []model.Song
)
