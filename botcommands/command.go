package botcommands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// HandleService handle a command event.
func HandleService(s *discordgo.Session, m *discordgo.MessageCreate, botPrefix string) {
	// var commandMsg []string = strings.Split(m.Content, " ")

	if strings.HasPrefix(m.Content, botPrefix+"help") {
		s.ChannelMessageSend(m.ChannelID, "ยังทำไม่เสร็จ กำลังทำอยู่ค้าบ")
	} else {
		s.ChannelMessageSend(m.ChannelID, botPrefix+"help เพื่อดูคำสั่งทั้งหมดนะค้าบ")
	}

}
