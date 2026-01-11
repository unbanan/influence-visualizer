package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"contest-influence/server/internal/config"
	"contest-influence/server/internal/service"

	"github.com/joho/godotenv"
)

func main() {
	rootCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var err error

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		panic("config path is required")
	}

	config, err := config.FromYaml(os.Args[1])
	if err != nil {
		panic(err)
	}

	service, err := service.New(config)
	if err != nil {
		panic(err)
	}

	err = service.Run()
	if err != nil {
		panic(err)
	}

	<-rootCtx.Done()
	cancel()

	err = service.Shutdown()
	if err != nil {
		panic(err)
	}
}
