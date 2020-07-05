package controller

import (
	"strings"

	"github.com/Planxnx/discordBot-Golang/internal/commands/services"
	messageService "github.com/Planxnx/discordBot-Golang/internal/messages/services"
	musicController "github.com/Planxnx/discordBot-Golang/internal/music/controller"
	voiceController "github.com/Planxnx/discordBot-Golang/internal/voice/controller"
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
		voiceController.PlayOKVoice(s, m, guild)
		go voiceController.StopVoice(m)
	} else if strings.HasPrefix(m.Content, botPrefix+"play") {
		var commandArgs []string = strings.Split(m.Content, " ")
		musicController.PlayYoutubeURL(commandArgs[1], s, m, guild)
	} else {
		go messageService.MessageSender(m.ChannelID, botPrefix+"help เพื่อดูคำสั่งทั้งหมดนะค้าบ")
	}
}
