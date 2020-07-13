package delivery

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/Planxnx/discordBot-Golang/internal/discord"
	musicUsecase "github.com/Planxnx/discordBot-Golang/internal/music/usecase"
	voiceUsecase "github.com/Planxnx/discordBot-Golang/internal/voice/usecase"
)

//Delivery interface
type Delivery interface {
	GetCommandsHandler(*discordgo.Session, *discordgo.MessageCreate)
}

type commandsDelivery struct {
	musicUsecase musicUsecase.Usecase
	voiceUsecase voiceUsecase.Usecase
	discord      discord.Discord
}

//NewCommandsDelivery new message delivery
func NewCommandsDelivery(discord discord.Discord, mu musicUsecase.Usecase, vu voiceUsecase.Usecase) Delivery {
	return &commandsDelivery{
		musicUsecase: mu,
		voiceUsecase: vu,
		discord:      discord,
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
		cd.discord.SendMessageToChannel(m.ChannelID, help)
	} else if strings.HasPrefix(m.Content, botPrefix+"join") {
		cd.voiceUsecase.ConnectToVoiceChannel(s, m, guild, true)
	} else if strings.HasPrefix(m.Content, botPrefix+"stop") {
		go cd.voiceUsecase.StopVoice()
		cd.discord.SendMessageToChannel(m.ChannelID, "หยุดเล่นแล้วค้าบ")
	} else if strings.HasPrefix(m.Content, botPrefix+"play") {
		var commandArgs []string = strings.Split(m.Content, " ")
		if len(commandArgs) > 1 {
			cd.musicUsecase.PlayYoutubeURL(commandArgs[1], s, m, guild)
		}
	} else {
		cd.discord.SendMessageToChannel(m.ChannelID, botPrefix+"help เพื่อดูคำสั่งทั้งหมดนะค้าบ")
	}
}
