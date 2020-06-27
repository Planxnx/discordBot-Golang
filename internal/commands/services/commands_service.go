package services

import (
	"log"

	voiceService "github.com/Planxnx/discordBot-Golang/internal/voice/services"
	"github.com/bwmarrin/discordgo"
)

//ConnectVoiceChannel connect to user voice channelId
func ConnectVoiceChannel(s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild) error {
	_, err := voiceService.ConnectToVoiceChannel(s, m, guild, true)
	if err != nil {
		log.Printf("Error: connect to voice channel, Message: '%s'", err)
	}
	return err
}
