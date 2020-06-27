package services

import (
	"github.com/Planxnx/discordBot-Golang/internal/discord"
)

type replyWordStruct struct {
	BadwordReply []string `json:"badwordReply"`
	KuyReply     []string `json:"kuyReply"`
}

// MessageSender send message by given channel id
func MessageSender(channelID string, msg string) {
	discord.SendMessageToChannel(channelID, msg)
}
