package viewmodel

import (
	"sophon-core/domain"
	"sophon-core/ports"
)

type ChatViewModel struct {
	Messages []domain.ConvoMessage
	Status   string
	ai       ports.AIProvider
}

func NewChatViewModel(ai ports.AIProvider) *ChatViewModel {
	return &ChatViewModel{
		ai: ai,
	}
}

func (vm *ChatViewModel) SendMessage(text string) error {
	msg := domain.ConvoMessage{
		Role:    "user",
		Message: text,
	}
	vm.Messages = append(vm.Messages, msg)
	vm.Status = "Thinking..."

	// In a real app, we would call vm.ai.Chat and update vm.Messages
	return nil
}
