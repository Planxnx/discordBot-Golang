package controller

import (
	"math/rand"
	"strings"

	"github.com/Planxnx/discordBot-Golang/internal/messages/repository"
	"github.com/Planxnx/discordBot-Golang/internal/messages/services"
	"github.com/bwmarrin/discordgo"
)

// MessageService handle a only message event.
func MessageService(s *discordgo.Session, m *discordgo.MessageCreate) {
	replyWord := repository.GetRandomReplyWord()
	if strings.Contains(m.Content, "ควย") || strings.Contains(m.Content, "8;p") {
		wordNumber := rand.Intn(len(replyWord.KuyReply))
		services.MessageSender(m.ChannelID, replyWord.KuyReply[wordNumber])
	} else if strings.Contains(m.Content, "สัส") || strings.Contains(m.Content, "เหี้ย") || strings.Contains(m.Content, "หี") {
		wordNumber := rand.Intn(len(replyWord.BadwordReply))
		services.MessageSender(m.ChannelID, replyWord.BadwordReply[wordNumber])
	}
}
