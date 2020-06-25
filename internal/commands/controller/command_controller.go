package controller

import (
	"fmt"
	"strings"

	"github.com/Planxnx/discordBot-Golang/internal/commands/services"
	messageService "github.com/Planxnx/discordBot-Golang/internal/messages/services"
	"github.com/bwmarrin/discordgo"
)

//CommandHandler handle a command event.
func CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, botPrefix string) {
	// var commandMsg []string = strings.Split(m.Content, " ")
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println(err)
	}
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		fmt.Println(err)
	}

	if strings.HasPrefix(m.Content, botPrefix+"help") {
		go messageService.MessageSender(m.ChannelID, "ยังทำไม่เสร็จ กำลังทำอยู่ค้าบ\nช่วยผมทำได้นะค้าบ เริ่มขี้เกียจแล้ว\nPull Request มาที่ https://github.com/Planxnx/discordBot-Golang")
	} else if strings.HasPrefix(m.Content, botPrefix+"join") {
		go services.ConnectToVoiceChannel(s, m, guild)
	} else {
		go messageService.MessageSender(m.ChannelID, botPrefix+"help เพื่อดูคำสั่งทั้งหมดนะค้าบ")
	}
}
