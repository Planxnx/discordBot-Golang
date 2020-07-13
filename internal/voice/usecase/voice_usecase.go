package usecase

import (
	"log"

	"github.com/Planxnx/discordBot-Golang/internal/discord"
	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

var stopChannel chan bool

//Usecase interface
type Usecase interface {
	PlayAudioFile(string, *discordgo.VoiceConnection)
	JoiAndPlayAudioFile(string, *discordgo.Session, *discordgo.MessageCreate, *discordgo.Guild, bool)
	ConnectToVoiceChannel(*discordgo.Session, *discordgo.MessageCreate, *discordgo.Guild, bool) (*discordgo.VoiceConnection, error)
	StopVoice()
}

type voiceUsecase struct {
	discord discord.Discord
}

//NewVoiceUsecase new voice usecase
func NewVoiceUsecase(discord discord.Discord) Usecase {
	return &voiceUsecase{
		discord: discord,
	}
}

//StopVoice stop voice channel
func (voiceUsecase) StopVoice() {
	if discord.GetVoiceStatus() {
		stopChannel <- true
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
func (vu voiceUsecase) JoiAndPlayAudioFile(file string, s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild, isMusicPlaying bool) {
	voiceConnection, err := connectToVoiceChannel(vu.discord, s, m, guild, isMusicPlaying)
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

//ConnectToVoiceChannel connect to user voice channelId
func (vu voiceUsecase) ConnectToVoiceChannel(s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild, isMusicPlaying bool) (*discordgo.VoiceConnection, error) {
	return connectToVoiceChannel(vu.discord, s, m, guild, isMusicPlaying)
}

func findVoiceChannelID(guild *discordgo.Guild, m *discordgo.MessageCreate) string {
	for _, voiceState := range guild.VoiceStates {
		if voiceState.UserID == m.Author.ID {
			return voiceState.ChannelID
		}
	}
	return ""
}
func connectToVoiceChannel(discord discord.Discord, s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild, isMustJoin bool) (voiceConnection *discordgo.VoiceConnection, err error) {
	voiceChannelID := findVoiceChannelID(guild, m)
	if voiceChannelID == "" && isMustJoin {
		if err := discord.SendMessageToChannel(m.ChannelID, "กรุณาเข้าห้องก่อนนะค้าบ"); err != nil {
			return nil, err
		}
	}
	voiceConnection, err = s.ChannelVoiceJoin(guild.ID, voiceChannelID, false, false)
	return
}
