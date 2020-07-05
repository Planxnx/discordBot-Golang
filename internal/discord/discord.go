package discord

import (
	"github.com/bwmarrin/discordgo"
)

var (
	session        *discordgo.Session
	voiceIsPlaying = false
)

//NewSession new Discord session
func NewSession(token string) error {
	var err error
	session, err = discordgo.New("Bot " + token)
	if err != nil {
		return err
	}
	return nil
}

//CreateConnection creates a websocket connection to Discord.
func CreateConnection() error {
	return session.Open()
}

//AddHandler add event handler
func AddHandler(handler interface{}) {
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
