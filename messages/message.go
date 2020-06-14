package messages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"

	"github.com/Planxnx/discordBot-Golang/botcommands"
	"github.com/bwmarrin/discordgo"
)

var badWordReply = [10]string{}

type replyWordStruct struct {
	BadwordReply []string `json:"badwordReply"`
	KuyReply     []string `json:"kuyReply"`
}

// HandleService handle all of message event from the given channel.
func HandleService(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	botPrefix := os.Getenv("BOT_PREFIX")
	if botPrefix == "" {
		botPrefix = "~"
	}

	if strings.HasPrefix(m.Content, botPrefix) {
		go botcommands.HandleService(s, m, botPrefix)
		return
	}
	go messageService(s, m)
}

func messageService(s *discordgo.Session, m *discordgo.MessageCreate) {
	messagesFile, err := os.Open("./data/messages.json")
	if err != nil {
		fmt.Println("Error at HandleService: opening messages.json,\nMsg: ", err)
	}
	defer messagesFile.Close()
	replyWordByteValue, _ := ioutil.ReadAll(messagesFile)
	var replyWord replyWordStruct
	json.Unmarshal(replyWordByteValue, &replyWord)

	if strings.Contains(m.Content, "ควย") {
		wordNumber := rand.Intn(len(replyWord.KuyReply))
		s.ChannelMessageSend(m.ChannelID, replyWord.KuyReply[wordNumber])
	} else if strings.Contains(m.Content, "สัส") || strings.Contains(m.Content, "เหี้ย") || strings.Contains(m.Content, "หี") {
		wordNumber := rand.Intn(len(replyWord.BadwordReply))
		s.ChannelMessageSend(m.ChannelID, replyWord.BadwordReply[wordNumber])
	}
}
