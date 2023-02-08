package main

import (
	"fmt"
	"os"
	"strings"
)

type config struct {
	openAPIKey string
	input      string
}

func newConfig(args []string) (config, error) {
	openAPIKey := os.Getenv("PLS_OPENAI_API_KEY")
	if openAPIKey == "" {
		return config{}, fmt.Errorf("PLS_OPENAI_API_KEY is not set")
	}

	input := strings.Join(args, " ")
	if input == "" {
		return config{}, fmt.Errorf("an input must be provided")
	}

	return config{
		openAPIKey: openAPIKey,
		input:      input,
	}, nil
}
