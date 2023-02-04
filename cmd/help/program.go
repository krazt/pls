package main

import (
	"context"
	"fmt"
	"strings"

	gpt "github.com/sashabaranov/go-gpt3"
)

type program struct {
	cfg       config
	gptClient *gpt.Client
}

func newProgram(cfg config) *program {
	gptClient := gpt.NewClient(cfg.openAPIKey)

	return &program{
		cfg:       cfg,
		gptClient: gptClient,
	}
}

func (p program) run(ctx context.Context) error {
	request := gpt.CompletionRequest{
		Model:            gpt.GPT3TextDavinci003,
		Prompt:           p.cfg.prompt,
		Suffix:           "",
		MaxTokens:        512,
		Temperature:      0,
		TopP:             1,
		N:                1,
		Stream:           false,
		LogProbs:         0,
		Echo:             false,
		Stop:             nil,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		BestOf:           1,
		LogitBias:        nil,
		User:             "",
	}

	response, err := p.gptClient.CreateCompletion(ctx, request)
	if err != nil {
		return fmt.Errorf("failed to create completion: %w", err)
	}

	if len(response.Choices) == 0 {
		return fmt.Errorf("completion response has no choices")
	}

	responseText := response.Choices[0].Text
	responseText = strings.TrimSpace(responseText)
	fmt.Println(responseText)

	return nil
}
