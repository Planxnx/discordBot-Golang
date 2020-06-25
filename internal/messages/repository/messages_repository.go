package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//ReplyWordStruct structure
type ReplyWordStruct struct {
	BadwordReply []string `json:"badwordReply"`
	KuyReply     []string `json:"kuyReply"`
}

//GetRandomReplyWord return strings of bad word
func GetRandomReplyWord() ReplyWordStruct {
	messagesFile, err := os.Open("./data/messages.json")
	if err != nil {
		fmt.Println("Error at HandleService: opening messages.json,\nMsg: ", err)
	}
	defer messagesFile.Close()
	replyWordByteValue, _ := ioutil.ReadAll(messagesFile)
	var replyWord ReplyWordStruct
	json.Unmarshal(replyWordByteValue, &replyWord)
	return replyWord
}
