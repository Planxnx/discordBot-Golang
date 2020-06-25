package services

import (
	"github.com/Planxnx/discordBot-Golang/internal/discord"
)

type replyWordStruct struct {
	BadwordReply []string `json:"badwordReply"`
	KuyReply     []string `json:"kuyReply"`
}

// MessageSender .
func MessageSender(channelID string, msg string) {
	discord.Session.ChannelMessageSend(channelID, msg)
}
