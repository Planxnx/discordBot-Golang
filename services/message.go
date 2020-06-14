package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var badWordReply = [10]string{}

type replyWordStruct struct {
	BadwordReply []string `json:"badwordReply"`
	KuyReply     []string `json:"kuyReply"`
}

func MessageService(s *discordgo.Session, m *discordgo.MessageCreate) {
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
		MessageSender(s, m.ChannelID, replyWord.KuyReply[wordNumber])
	} else if strings.Contains(m.Content, "สัส") || strings.Contains(m.Content, "เหี้ย") || strings.Contains(m.Content, "หี") {
		wordNumber := rand.Intn(len(replyWord.BadwordReply))
		MessageSender(s, m.ChannelID, replyWord.BadwordReply[wordNumber])
	}
}

func MessageSender(s *discordgo.Session, channelID string, msg string) {
	s.ChannelMessageSend(channelID, msg)
}
