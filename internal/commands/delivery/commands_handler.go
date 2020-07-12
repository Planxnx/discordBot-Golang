package delivery

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/Planxnx/discordBot-Golang/internal/commands/services"
	messageService "github.com/Planxnx/discordBot-Golang/internal/messages/services"
	musicUsecase "github.com/Planxnx/discordBot-Golang/internal/music/usecase"
	voiceServices "github.com/Planxnx/discordBot-Golang/internal/voice/services"
	voiceUsecase "github.com/Planxnx/discordBot-Golang/internal/voice/usecase"
)

//Delivery interface
type Delivery interface {
	GetCommandsHandler(*discordgo.Session, *discordgo.MessageCreate)
}

type commandsDelivery struct {
	musicUsecase musicUsecase.Usecase
	voiceUsecase voiceUsecase.Usecase
}

//NewCommandsDelivery new message delivery
func NewCommandsDelivery(mu musicUsecase.Usecase, vu voiceUsecase.Usecase) Delivery {
	return &commandsDelivery{
		musicUsecase: mu,
		voiceUsecase: vu,
	}
}

func (cd commandsDelivery) GetCommandsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
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

	if !strings.HasPrefix(m.Content, botPrefix) {
		return
	}

	if strings.HasPrefix(m.Content, botPrefix+"help") {
		help := fmt.Sprintf("**รายชื่อคำสั่งนะจ้า (ยังไม่เสร็จ)**\n==============================\n`%splay [Youtube Link]` : เล่นเพลงจากยูทูป (ตอนนี้เล่นได้แค่ทีล่ะเพลง, ยังเสริชเพลงไม่ได้)\n`%sstop` : สั่งให้หยุดเล่นเพลง\n`%sjoin` : สั่งให้บอทเข้ามาในห้อง\n==============================\nถ้าเจอบัคฝากแจ้งหน่อยนะจ้า", botPrefix, botPrefix, botPrefix)
		go messageService.MessageSender(m.ChannelID, help)
	} else if strings.HasPrefix(m.Content, botPrefix+"join") {
		go services.ConnectVoiceChannel(s, m, guild)
	} else if strings.HasPrefix(m.Content, botPrefix+"stop") {
		voiceServices.StopVoice()
		messageService.MessageSender(m.ChannelID, "หยุดเล่นแล้วค้าบ")
	} else if strings.HasPrefix(m.Content, botPrefix+"play") {
		var commandArgs []string = strings.Split(m.Content, " ")
		if len(commandArgs) > 1 {
			cd.musicUsecase.PlayYoutubeURL(commandArgs[1], s, m, guild)
		}
	} else {
		go messageService.MessageSender(m.ChannelID, botPrefix+"help เพื่อดูคำสั่งทั้งหมดนะค้าบ")
	}
}
