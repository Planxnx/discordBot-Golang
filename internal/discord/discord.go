package discord

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	session        *discordgo.Session
	voiceIsPlaying = false
)

//Discord interface
type Discord interface {
	AddHandler(handler interface{})
}

type discordSession struct{}

//NewSession new Discord session
func NewSession(logger *log.Logger) (Discord, error) {
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		return nil, fmt.Errorf("Error: BOT_TOKEN not found, Closing now")
	}

	logger.Println("Discord Session is starting with token '", botToken, "'")
	var err error
	session, err = discordgo.New("Bot " + botToken)
	if err != nil {
		return nil, err
	}
	if err := session.Open(); err != nil {
		return nil, err
	}
	return &discordSession{}, nil
}

//AddHandler add event handler
func (discordSession) AddHandler(handler interface{}) {
	session.AddHandler(handler)
}

//CloseConnection closes a websocket and stops all listening/heartbeat goroutines.
func CloseConnection() error {
	return session.Close()
}

//SendMessageToChannel send message to the given channel id
func SendMessageToChannel(channelID string, message string) error {
	_, err := session.ChannelMessageSend(channelID, message)
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
