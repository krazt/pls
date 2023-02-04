package main

import (
	"fmt"
	"os"
)

type config struct {
	openAPIKey string
	prompt     string
}

func newConfig(args []string) (config, error) {
	openAPIKey := os.Getenv("HELP_OPENAI_API_KEY")
	if openAPIKey == "" {
		return config{}, fmt.Errorf("HELP_OPENAI_API_KEY is not set")
	}

	if len(args) == 0 {
		return config{}, fmt.Errorf("a prompt must be provided")
	}
	prompt := args[0]

	return config{
		openAPIKey: openAPIKey,
		prompt:     prompt,
	}, nil
}
