package controller

import (
	"fmt"
	"strings"

	"github.com/Planxnx/discordBot-Golang/services"
	"github.com/bwmarrin/discordgo"
)

// CommandService handle a command event.
func CommandService(s *discordgo.Session, m *discordgo.MessageCreate, botPrefix string) {
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
		go services.MessageSender(m.ChannelID, "ยังทำไม่เสร็จ กำลังทำอยู่ค้าบ\nช่วยผมทำได้นะค้าบ เริ่มขี้เกียจแล้ว\nPull Request มาที่ https://github.com/Planxnx/discordBot-Golang")
	} else if strings.HasPrefix(m.Content, botPrefix+"join") {
		go connectToVoiceChannel(s, m, guild)
	} else {
		go services.MessageSender(m.ChannelID, botPrefix+"help เพื่อดูคำสั่งทั้งหมดนะค้าบ")
	}
}

func findVoiceChannelID(guild *discordgo.Guild, m *discordgo.MessageCreate) string {
	for _, voiceState := range guild.VoiceStates {
		if voiceState.UserID == m.Author.ID {
			return voiceState.ChannelID
		}
	}
	return ""
}

func connectToVoiceChannel(s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild) {
	voiceChannelID := findVoiceChannelID(guild, m)
	if voiceChannelID == "" {
		services.MessageSender(m.ChannelID, "กรุณาเข้าห้องก่อนนะค้าบ")
	}

	_, err := s.ChannelVoiceJoin(guild.ID, voiceChannelID, false, false)

	if err != nil {
		fmt.Println(err)
	}
}
