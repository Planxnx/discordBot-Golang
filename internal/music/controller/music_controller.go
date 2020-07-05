package controller

import (
	"log"

	"github.com/Planxnx/discordBot-Golang/internal/music/services"
	voiceServices "github.com/Planxnx/discordBot-Golang/internal/voice/services"
	"github.com/bwmarrin/discordgo"
)

//PlayYoutubeURL .
func PlayYoutubeURL(s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild) {
	url, err := services.GetYoutubeDownloadURL("https://youtu.be/D4bRL3n88kA")
	if err != nil {
		log.Printf("Error: can't get youtube download url, Message: '%s'", err)
	}

	voiceConnection, err := voiceServices.ConnectToVoiceChannel(s, m, guild, true)
	if err != nil {
		log.Printf("Error: connect to voice channel, Message: '%s'", err)
	}
	go voiceServices.PlayAudioFile(url, voiceConnection)
}
