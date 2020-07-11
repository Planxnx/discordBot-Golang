package delivery

import (
	"log"
	"os"
	"strings"

	voiceController "github.com/Planxnx/discordBot-Golang/internal/voice/controller"

	"github.com/Planxnx/discordBot-Golang/internal/messages/services"
	"github.com/bwmarrin/discordgo"
)

//Delivery interface
type Delivery interface {
	GetMessageHandler(*discordgo.Session, *discordgo.MessageCreate)
}

type messageDelivery struct{}

//NewMessageDelivery new message delivery
func NewMessageDelivery() Delivery {
	return &messageDelivery{}
}

func (messageDelivery) GetMessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	botPrefix := os.Getenv("BOT_PREFIX")
	if botPrefix == "" {
		botPrefix = "~"
	}
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Println(err)
	}
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		log.Println(err)
	}

	// go messagesController.MessageHandler(s, m, guild)
	if strings.Contains(m.Content, "ควย") || strings.Contains(m.Content, "8;p") {
		voiceController.PlayKuyVoice(s, m, guild)
		services.MessageSender(m.ChannelID, services.GetRandomKuyReplyWord())
	} else if strings.Contains(m.Content, "สัส") || strings.Contains(m.Content, "เหี้ย") || strings.Contains(m.Content, "หี") {
		services.MessageSender(m.ChannelID, services.GetRandomReplyWord())
	}
}
