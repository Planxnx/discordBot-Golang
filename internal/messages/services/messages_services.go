package services

import (
	"math/rand"

	"github.com/Planxnx/discordBot-Golang/internal/discord"
	"github.com/Planxnx/discordBot-Golang/internal/messages/repository"
)

type replyWordStruct struct {
	BadwordReply []string `json:"badwordReply"`
	KuyReply     []string `json:"kuyReply"`
}

// MessageSender send message by given channel id
func MessageSender(channelID string, msg string) {
	discord.SendMessageToChannel(channelID, msg)
}

// GetRandomReplyWord return bad word
func GetRandomReplyWord() string {
	replyWord := repository.GetBadWordList()
	wordIndex := rand.Intn(len(replyWord.BadwordReply))
	return replyWord.KuyReply[wordIndex]
}

// GetRandomKuyReplyWord return bad word kuy
func GetRandomKuyReplyWord() string {
	replyWord := repository.GetBadWordList()
	wordIndex := rand.Intn(len(replyWord.KuyReply))
	return replyWord.KuyReply[wordIndex]
}
