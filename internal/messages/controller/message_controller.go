package controller

import (
	"strings"

	"github.com/Planxnx/discordBot-Golang/internal/messages/services"
	voiceController "github.com/Planxnx/discordBot-Golang/internal/voice/controller"
	"github.com/bwmarrin/discordgo"
)

// MessageHandler handle a only message event.
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild) {
	if strings.Contains(m.Content, "ควย") || strings.Contains(m.Content, "8;p") {
		go voiceController.PlayKuyVoice(s, m, guild)
		services.MessageSender(m.ChannelID, services.GetRandomKuyReplyWord())
	} else if strings.Contains(m.Content, "สัส") || strings.Contains(m.Content, "เหี้ย") || strings.Contains(m.Content, "หี") {
		services.MessageSender(m.ChannelID, services.GetRandomReplyWord())
	}
}
