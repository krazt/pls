package main

import (
	"fmt"
	"os"
)

type config struct {
	openAPIKey string
	input      string
}

func newConfig(args []string) (config, error) {
	openAPIKey := os.Getenv("HELP_OPENAI_API_KEY")
	if openAPIKey == "" {
		return config{}, fmt.Errorf("HELP_OPENAI_API_KEY is not set")
	}

	if len(args) == 0 {
		return config{}, fmt.Errorf("an input must be provided")
	}
	input := args[0]

	return config{
		openAPIKey: openAPIKey,
		input:      input,
	}, nil
}