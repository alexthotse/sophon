package model

import (
	"context"
	"fmt"
	"sophon-core/domain"

	"github.com/firebase/genkit/go/ai"
)

// SophonAgentSwarm is an implementation of a genkit-powered agent swarm.
type SophonAgentSwarm struct {
	model ai.Model
}

func NewSophonAgentSwarm(model ai.Model) *SophonAgentSwarm {
	return &SophonAgentSwarm{
		model: model,
	}
}

func (s *SophonAgentSwarm) Chat(ctx context.Context, messages []domain.ConvoMessage) (domain.ConvoMessage, error) {
	var genkitMessages []*ai.Message
	for _, m := range messages {
		role := ai.RoleUser
		if m.Role == "assistant" {
			role = ai.RoleAssistant
		}
		genkitMessages = append(genkitMessages, &ai.Message{
			Role: role,
			Content: []*ai.Part{
				ai.NewTextPart(m.Message),
			},
		})
	}

	resp, err := s.model.Generate(ctx, &ai.GenerateRequest{
		Messages: genkitMessages,
	}, nil)

	if err != nil {
		return domain.ConvoMessage{}, err
	}

	if len(resp.Candidates) == 0 {
		return domain.ConvoMessage{}, fmt.Errorf("no candidates in response")
	}

	return domain.ConvoMessage{
		Role:    "assistant",
		Message: resp.Candidates[0].Message.Content[0].Text,
	}, nil
}

func (s *SophonAgentSwarm) Stream(ctx context.Context, messages []domain.ConvoMessage) (<-chan domain.ConvoMessage, error) {
	ch := make(chan domain.ConvoMessage)
	// Simplified streaming implementation for brevity
	go func() {
		defer close(ch)
		res, _ := s.Chat(ctx, messages)
		ch <- res
	}()
	return ch, nil
}
