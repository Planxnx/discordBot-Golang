package controller

import (
	"fmt"
	"strings"

	"github.com/Planxnx/discordBot-Golang/internal/commands/services"
	messageService "github.com/Planxnx/discordBot-Golang/internal/messages/services"
	musicController "github.com/Planxnx/discordBot-Golang/internal/music/controller"
	voiceServices "github.com/Planxnx/discordBot-Golang/internal/voice/services"

	"github.com/bwmarrin/discordgo"
)

//CommandHandler handle a command event.
func CommandHandler(s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild, botPrefix string) {
	if strings.HasPrefix(m.Content, botPrefix+"help") {
		help := fmt.Sprintf("รายชื่อคำสั่ง\n%splay [Youtube Link] : เล่นเพลงจากยูทูป (ตอนนี้เล่นได้แค่ทีล่ะเพลง, ยังไม่สามารถค้นหาเพลงได้)\n%sstop : สั่งให้หยุดเล่นเพลง", botPrefix, botPrefix)
		go messageService.MessageSender(m.ChannelID, help)
	} else if strings.HasPrefix(m.Content, botPrefix+"join") {
		go services.ConnectVoiceChannel(s, m, guild)
	} else if strings.HasPrefix(m.Content, botPrefix+"stop") {
		voiceServices.StopVoice()
		messageService.MessageSender(m.ChannelID, "หยุดเล่นแล้วค้าบ")
	} else if strings.HasPrefix(m.Content, botPrefix+"play") {
		var commandArgs []string = strings.Split(m.Content, " ")
		if len(commandArgs) > 1 {
			musicController.PlayYoutubeURL(commandArgs[1], s, m, guild)
		}
	} else {
		go messageService.MessageSender(m.ChannelID, botPrefix+"help เพื่อดูคำสั่งทั้งหมดนะค้าบ")
	}
}
