package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
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

//go:embed assets/templates/prompt.txt
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

	err = p.handleResponse(response)
	if err != nil {
		return fmt.Errorf("failed to handle response: %w", err)
	}

	return nil
}

func (p program) handleResponse(response gpt.CompletionResponse) error {
	if len(response.Choices) == 0 {
		return fmt.Errorf("completion response has no choices")
	}

	responseText := response.Choices[0].Text
	responseText = strings.TrimSpace(responseText)

	if len(responseText) < 4 {
		return fmt.Errorf("completion response is too short")
	}

	switch {
	case strings.HasPrefix(responseText, "E:"):
		responseText = strings.TrimPrefix(responseText, "E:")
		responseText = strings.TrimSpace(responseText)

		fmt.Fprintln(os.Stderr, responseText)
	case strings.HasPrefix(responseText, "O:"):
		responseText = strings.TrimPrefix(responseText, "O:")
		responseText = strings.TrimSpace(responseText)

		err := p.handlePrediction(responseText)
		if err != nil {
			return fmt.Errorf("failed to handle prediction: %w", err)
		}
	default:
		return fmt.Errorf("completion response has unknown prefix")
	}

	return nil
}

func (p program) handlePrediction(prediction string) error {
	fmt.Printf("%s\n\n%s\n", prediction, "Run the command? [y/N]")

	var option string
	_, err := fmt.Scanln(&option)
	if err != nil {
		return fmt.Errorf("failed to scan input: %w", err)
	}
	option = strings.ToLower(option)

	if option == "y" {
		err := p.runCommand(prediction)
		if err != nil {
			return fmt.Errorf("failed to run command: %w", err)
		}
	} else {
		fmt.Println("The command was not run")
	}

	return nil
}

func (p program) runCommand(cmd string) error {
	c := exec.Command("sh", "-c", cmd)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin

	fmt.Println()

	err := c.Run()
	if err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}

	return nil
}

func newCompletionRequest(prompt string, stop []string) gpt.CompletionRequest {
	return gpt.CompletionRequest{
		Model:            gpt.GPT3TextDavinci003,
		Prompt:           prompt,
		Suffix:           "",
		MaxTokens:        128,
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
