package usecase

import (
	"fmt"
	"log"

	"github.com/Planxnx/discordBot-Golang/internal/discord"
	messageService "github.com/Planxnx/discordBot-Golang/internal/messages/services"
	voiceUsecase "github.com/Planxnx/discordBot-Golang/internal/voice/usecase"
	youtubeUsecase "github.com/Planxnx/discordBot-Golang/internal/youtube/usecase"
	"github.com/bwmarrin/discordgo"
)

//Usecase interface
type Usecase interface {
	PlayYoutubeURL(string, *discordgo.Session, *discordgo.MessageCreate, *discordgo.Guild)
}

type musicUsecase struct {
	youtubeUsecase youtubeUsecase.Usecase
	voiceUsecase   voiceUsecase.Usecase
}

//NewMusicUsecase new message delivery
func NewMusicUsecase(yu youtubeUsecase.Usecase, vu voiceUsecase.Usecase) Usecase {
	return &musicUsecase{
		youtubeUsecase: yu,
		voiceUsecase:   vu,
	}
}

func (mu musicUsecase) PlayYoutubeURL(url string, s *discordgo.Session, m *discordgo.MessageCreate, guild *discordgo.Guild) {
	voiceConnection, err := mu.voiceUsecase.ConnectToVoiceChannel(s, m, guild, true)
	if err != nil {
		log.Printf("Error: connect to voice channel, Message: '%s'", err)
		messageService.MessageSender(m.ChannelID, "มีปัญหาเข้าห้องไม่ได้ หรือ พูดไม่ได้จ้า")
		return
	}

	if discord.GetVoiceStatus() {
		messageService.MessageSender(m.ChannelID, "รอเพลงเล่นเสร็จก่อนแปปนึงนะค้าบ")
		return
	}
	youtubeInfo, err := mu.youtubeUsecase.GetYoutubeDownloadURL(url)
	if err != nil {
		log.Printf("Error: can't get youtube download url, Message: '%s'", err)
		messageService.MessageSender(m.ChannelID, "หาเพลงไม่เจอค้าบ")
		return
	}
	msg := fmt.Sprintf("กำลังจะเล่น '%s' นะค้าบ", youtubeInfo.Title)
	messageService.MessageSender(m.ChannelID, msg)
	mu.voiceUsecase.PlayAudioFile(youtubeInfo.DownloadLink, voiceConnection)
}
