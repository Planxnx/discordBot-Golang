package services

import (
	"github.com/Planxnx/discordBot-Golang/internal/discord"
	messagesService "github.com/Planxnx/discordBot-Golang/internal/messages/services"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

var stopChannel chan bool

//InitVoiceChannel .
func InitVoiceChannel() {
	stopChannel = make(chan bool)
}

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
func ConnectToVoiceChannel(s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild, isMustJoin bool) (voiceConnection *discordgo.VoiceConnection, err error) {
	voiceChannelID := FindVoiceChannelID(guild, m)
	if voiceChannelID == "" && isMustJoin {
		messagesService.MessageSender(m.ChannelID, "กรุณาเข้าห้องก่อนนะค้าบ")
	}

	voiceConnection, err = s.ChannelVoiceJoin(guild.ID, voiceChannelID, false, false)
	return
}

//PlayAudioFile .
func PlayAudioFile(file string, voiceConnection *discordgo.VoiceConnection) {
	defer discord.UpdateVoiceStatus(false) 
	if !discord.GetVoiceStatus() {
		discord.UpdateVoiceStatus(true)
		dgvoice.PlayAudioFile(voiceConnection, file, stopChannel)
	}
}

//StopVoice stop voice channel
func StopVoice() {
	if discord.GetVoiceStatus() {
		stopChannel <- true
	}
}
