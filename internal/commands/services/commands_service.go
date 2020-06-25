package services

import (
	"fmt"

	messagesService "github.com/Planxnx/discordBot-Golang/internal/messages/services"
	"github.com/bwmarrin/discordgo"
)

//FindVoiceChannelID find user channelId
func FindVoiceChannelID(guild *discordgo.Guild, m *discordgo.MessageCreate) string {
	for _, voiceState := range guild.VoiceStates {
		if voiceState.UserID == m.Author.ID {
			return voiceState.ChannelID
		}
	}
	return ""
}

//ConnectToVoiceChannel connect to user voice channelId
func ConnectToVoiceChannel(s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild) {
	voiceChannelID := FindVoiceChannelID(guild, m)
	if voiceChannelID == "" {
		messagesService.MessageSender(m.ChannelID, "กรุณาเข้าห้องก่อนนะค้าบ")
	}

	_, err := s.ChannelVoiceJoin(guild.ID, voiceChannelID, false, false)

	if err != nil {
		fmt.Println(err)
	}
}
