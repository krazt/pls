package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var statusCode int
	defer func() {
		os.Exit(statusCode)
	}()

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	cfg, err := newConfig(os.Args[1:])
	if err != nil {
		log.Println("failed to create config:", err)
		statusCode = 2
		return
	}
	p := newProgram(cfg)

	err = p.run(ctx)
	if err != nil {
		log.Println(err)
		statusCode = 1
		return
	}
}
