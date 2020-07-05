package controller

import (
	"fmt"
	"log"

	"github.com/Planxnx/discordBot-Golang/internal/discord"
	messageService "github.com/Planxnx/discordBot-Golang/internal/messages/services"
	"github.com/Planxnx/discordBot-Golang/internal/music/services"
	voiceServices "github.com/Planxnx/discordBot-Golang/internal/voice/services"
	"github.com/bwmarrin/discordgo"
)

//PlayYoutubeURL .
func PlayYoutubeURL(url string, s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild) {
	voiceConnection, err := voiceServices.ConnectToVoiceChannel(s, m, guild, true)
	if err != nil {
		log.Printf("Error: connect to voice channel, Message: '%s'", err)
		return
	}

	if discord.GetVoiceStatus() {
		messageService.MessageSender(m.ChannelID, "รอเพลงเล่นเสร็จก่อนแปปนึงนะค้าบ")
		return
	}

	youtubeInfo, err := services.GetYoutubeDownloadURL(url)
	if err != nil {
		log.Printf("Error: can't get youtube download url, Message: '%s'", err)
		messageService.MessageSender(m.ChannelID, "หาเพลงไม่เจอค้าบ")
		return
	}
	msg := fmt.Sprintf("กำลังจะเล่น '%s' นะค้าบ", youtubeInfo.Title)
	messageService.MessageSender(m.ChannelID, msg)
	discord.UpdateVoiceStatus(true)
	voiceServices.PlayAudioFile(youtubeInfo.DownloadLink, voiceConnection)
	discord.UpdateVoiceStatus(false)
}
