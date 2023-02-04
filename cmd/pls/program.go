package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"os"
	"runtime"
	"strings"
	"text/template"

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

//go:embed prompt_template.md
var promptTemplateStr string

type promptData struct {
	Input string
	OS    string
	Arch  string
	Dir   string
}

func newPromptData(input string) promptData {
	dir, err := os.Getwd()
	if err != nil {
		dir = "?"
	}

	return promptData{
		Input: input,
		OS:    runtime.GOOS,
		Arch:  runtime.GOARCH,
		Dir:   dir,
	}
}

func (p program) run(ctx context.Context) error {
	s := strings.TrimSpace(promptTemplateStr)
	promptTemplate, err := template.New("").Parse(s)
	if err != nil {
		return fmt.Errorf("failed to parse prompt template: %w", err)
	}

	data := newPromptData(p.cfg.input)

	var buf bytes.Buffer
	err = promptTemplate.Execute(&buf, data)
	if err != nil {
		return fmt.Errorf("failed to execute prompt template: %w", err)
	}
	prompt := buf.String()

	request := newCompletionRequest(prompt, []string{`\n---\n`})

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

func newCompletionRequest(prompt string, stop []string) gpt.CompletionRequest {
	return gpt.CompletionRequest{
		Model:            gpt.GPT3TextDavinci003,
		Prompt:           prompt,
		Suffix:           "",
		MaxTokens:        512,
		Temperature:      0,
		TopP:             1,
		N:                1,
		Stream:           false,
		LogProbs:         0,
		Echo:             false,
		Stop:             stop,
		PresencePenalty:  0,
		FrequencyPenalty: 0,
		BestOf:           1,
		LogitBias:        nil,
		User:             "",
	}
}
