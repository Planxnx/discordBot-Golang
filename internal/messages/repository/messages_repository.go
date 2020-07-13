package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Planxnx/discordBot-Golang/internal/discord"
)

//Repository interface
type Repository interface {
	GetBadWordList() (ReplyWordStruct, error)
}

type messageRepository struct {
	discord discord.Discord
}

//NewMessageRepository new message repository
func NewMessageRepository(discord discord.Discord) Repository {
	return &messageRepository{
		discord: discord,
	}
}

//ReplyWordStruct structure
type ReplyWordStruct struct {
	BadwordReply []string `json:"badwordReply"`
	KuyReply     []string `json:"kuyReply"`
}

//GetBadWordList return list of bad word
func (messageRepository) GetBadWordList() (ReplyWordStruct, error) {
	//need to injection config
	messagesFile, err := os.Open("./data/messages.json")
	if err != nil {
		fmt.Println("Error at HandleService: opening messages.json,\nMsg: ", err)
		return ReplyWordStruct{}, err
	}
	defer messagesFile.Close()
	replyWordByteValue, _ := ioutil.ReadAll(messagesFile)
	var replyWord ReplyWordStruct
	json.Unmarshal(replyWordByteValue, &replyWord)
	return replyWord, nil
}
