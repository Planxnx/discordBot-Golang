package discord

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	instance       *discordSession
	voiceIsPlaying = false
)

//Discord interface
type Discord interface {
	AddHandler(handler interface{})
	SendMessageToChannel(string, string) error
	CloseConnection() error
	OpenConnection() error
}

type discordSession struct {
	session *discordgo.Session
}

//NewSession new Discord session
func NewSession(logger *log.Logger) (Discord, error) {
	if session != nil {
		return &discordSession{}, nil
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		return nil, fmt.Errorf("Error: BOT_TOKEN not found, Closing now")
	}

	session, err := discordgo.New("Bot " + botToken)
	if err != nil {
		return nil, err
	}

	instance = &discordSession{
		session: session,
	}

	return instance, nil
}

//AddHandler add event handler
func (ds discordSession) AddHandler(handler interface{}) {
	ds.session.AddHandler(handler)
}

//OpenConnection open a websocket and listening goroutines.
func (ds discordSession) OpenConnection() error {
	log.Println("Discord Session is starting with token '", ds.session.Token, "'")
	return ds.session.Open()
}

//CloseConnection closes a websocket and stops all listening/heartbeat goroutines.
func (ds discordSession) CloseConnection() error {
	return ds.session.Close()
}

//SendMessageToChannel send message to the given channel id
func (ds discordSession) SendMessageToChannel(channelID string, message string) error {
	_, err := ds.session.ChannelMessageSend(channelID, message)
	return err
}

//VoiceStatusSwitch switch on/off voice channel
func VoiceStatusSwitch() {
	voiceIsPlaying = !voiceIsPlaying
}

//UpdateVoiceStatus update voice channel status
func UpdateVoiceStatus(status bool) {
	voiceIsPlaying = status
}

//GetVoiceStatus update voice channel status
func GetVoiceStatus() bool {
	return voiceIsPlaying
}
