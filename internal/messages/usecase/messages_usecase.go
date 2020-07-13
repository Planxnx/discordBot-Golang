package usecase

import (
	"math/rand"

	"github.com/Planxnx/discordBot-Golang/internal/messages/repository"
)

//Usecase interface
type Usecase interface {
	GetRandomKuyReplyWord() string
	GetRandomReplyWord() string
}

type messagesUsecase struct {
	messagesRepository repository.Repository
}

//NewMessagesUsecase new message delivery
func NewMessagesUsecase(mr repository.Repository) Usecase {
	return &messagesUsecase{
		messagesRepository: mr,
	}
}

// GetRandomKuyReplyWord return bad word kuy
func (mu messagesUsecase) GetRandomKuyReplyWord() string {
	replyWord, err := mu.messagesRepository.GetBadWordList()
	if err != nil {
		return "8;p"
	}
	wordIndex := rand.Intn(len(replyWord.KuyReply))
	return replyWord.KuyReply[wordIndex]
}

// GetRandomReplyWord return bad word
func (mu messagesUsecase) GetRandomReplyWord() string {
	replyWord, err := mu.messagesRepository.GetBadWordList()
	if err != nil {
		return "หยาบคายยย"
	}
	wordIndex := rand.Intn(len(replyWord.BadwordReply))
	return replyWord.KuyReply[wordIndex]
}
