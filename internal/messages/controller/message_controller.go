package controller

import (
	"math/rand"
	"strings"

	"github.com/Planxnx/discordBot-Golang/internal/messages/repository"
	"github.com/Planxnx/discordBot-Golang/internal/messages/services"
	voiceController "github.com/Planxnx/discordBot-Golang/internal/voice/controller"
	"github.com/bwmarrin/discordgo"
)

// MessageHandler handle a only message event.
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild) {
	replyWord := repository.GetRandomReplyWord()
	if strings.Contains(m.Content, "ควย") || strings.Contains(m.Content, "8;p") {
		wordNumber := rand.Intn(len(replyWord.KuyReply))
		voiceController.PlayKuyVoice(s, m, guild)
		services.MessageSender(m.ChannelID, replyWord.KuyReply[wordNumber])
	} else if strings.Contains(m.Content, "สัส") || strings.Contains(m.Content, "เหี้ย") || strings.Contains(m.Content, "หี") {
		wordNumber := rand.Intn(len(replyWord.BadwordReply))
		services.MessageSender(m.ChannelID, replyWord.BadwordReply[wordNumber])
	}
}
