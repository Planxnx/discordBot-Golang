package discord

import (
	"github.com/bwmarrin/discordgo"
)

var session *discordgo.Session

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
