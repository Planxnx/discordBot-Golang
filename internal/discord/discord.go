package discord

import (
	"github.com/bwmarrin/discordgo"
)

//Session is discordgo.Session
var Session *discordgo.Session

//NewSession new Discord session
func NewSession(token string) error {
	var err error
	Session, err = discordgo.New("Bot " + token)
	if err != nil {
		return err
	}
	return nil
}

//CreateConnection creates a websocket connection to Discord.
func CreateConnection() error {
	return Session.Open()
}

//AddHandler add event handler
func AddHandler(handler interface{}) {
	Session.AddHandler(handler)
}

//CloseConnection closes a websocket and stops all listening/heartbeat goroutines.
func CloseConnection() error {
	return Session.Close()
}

//SendMessageToChannel send message to the given channel id
func SendMessageToChannel(channelID string, message string) error {
	_, err := Session.ChannelMessageSend(channelID, message)
	return err
}
