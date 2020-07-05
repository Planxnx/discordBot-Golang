package controller

import (
	"log"

	messageService "github.com/Planxnx/discordBot-Golang/internal/messages/services"
	"github.com/Planxnx/discordBot-Golang/internal/voice/services"
	"github.com/bwmarrin/discordgo"
)

//StopVoice stop voice
func StopVoice(m *discordgo.MessageCreate) {
	services.StopVoice()
	messageService.MessageSender(m.ChannelID, "หยุดเล่นแล้วค้าบ")
}

//PlayKuyVoice pen-kuy-rai sound
func PlayKuyVoice(s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild) {
	PlayAudioFile("./sound/pen-kuy-rai.mp3", s, m, guild, false)
}

//PlayOKVoice ok sound
func PlayOKVoice(s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild) {
	PlayAudioFile("./sound/ok.mp3", s, m, guild, false)
}

//PlayAudioFile .
func PlayAudioFile(file string, s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild, isMusicPlaying bool) {
	voiceConnection, err := services.ConnectToVoiceChannel(s, m, guild, isMusicPlaying)
	if err != nil {
		log.Printf("Error: connect to voice channel, Message: '%s'", err)
	}
	services.PlayAudioFile(file, voiceConnection)
}
