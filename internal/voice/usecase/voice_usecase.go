package usecase

import (
	"log"

	"github.com/Planxnx/discordBot-Golang/internal/discord"
	"github.com/Planxnx/discordBot-Golang/internal/voice/services"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

var stopChannel chan bool

//Usecase interface
type Usecase interface {
	PlayAudioFile(string, *discordgo.VoiceConnection)
	JoiAndPlayAudioFile(string, *discordgo.Session, *discordgo.MessageCreate, *discordgo.Guild, bool)
}

type voiceUsecase struct {
	discord discord.Discord
}

//NewVoiceUsecase new message delivery
func NewVoiceUsecase(discord discord.Discord) Usecase {
	return &voiceUsecase{
		discord: discord,
	}
}

//PlayAudioFile return youtube download url
func (voiceUsecase) PlayAudioFile(file string, voiceConnection *discordgo.VoiceConnection) {
	if !discord.GetVoiceStatus() {
		stopChannel = make(chan bool)
		discord.UpdateVoiceStatus(true)
		dgvoice.PlayAudioFile(voiceConnection, file, stopChannel)
		close(stopChannel)
		discord.UpdateVoiceStatus(false)
	}
}

//JoiAndPlayAudioFile return youtube download url
func (voiceUsecase) JoiAndPlayAudioFile(file string, s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild, isMusicPlaying bool) {
	voiceConnection, err := services.ConnectToVoiceChannel(s, m, guild, isMusicPlaying)
	if err != nil {
		log.Printf("Error: connect to voice channel, Message: '%s'", err)
	}
	if !discord.GetVoiceStatus() {
		stopChannel = make(chan bool)
		discord.UpdateVoiceStatus(true)
		dgvoice.PlayAudioFile(voiceConnection, file, stopChannel)
		close(stopChannel)
		discord.UpdateVoiceStatus(false)
	}
}
