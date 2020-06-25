package services

import (
	"math/rand"
	"strings"

	"github.com/Planxnx/discordBot-Golang/internal/discord"
	"github.com/Planxnx/discordBot-Golang/internal/messages/repository"
	"github.com/bwmarrin/discordgo"
)

type replyWordStruct struct {
	BadwordReply []string `json:"badwordReply"`
	KuyReply     []string `json:"kuyReply"`
}

// MessageService handle a only message event.
func MessageService(s *discordgo.Session, m *discordgo.MessageCreate) {
	replyWord := repository.GetRandomReplyWord()
	if strings.Contains(m.Content, "ควย") {
		wordNumber := rand.Intn(len(replyWord.KuyReply))
		MessageSender(m.ChannelID, replyWord.KuyReply[wordNumber])
	} else if strings.Contains(m.Content, "สัส") || strings.Contains(m.Content, "เหี้ย") || strings.Contains(m.Content, "หี") {
		wordNumber := rand.Intn(len(replyWord.BadwordReply))
		MessageSender(m.ChannelID, replyWord.BadwordReply[wordNumber])
	}
}

// MessageSender .
func MessageSender(channelID string, msg string) {
	discord.Session.ChannelMessageSend(channelID, msg)
}
