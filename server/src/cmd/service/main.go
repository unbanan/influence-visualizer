package main

import (
	"context"
	"os/signal"
	"syscall"

	"contest-influence/server/internal/config"
	"contest-influence/server/internal/service"
)

func main() {
	rootCtx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	var err error

	config, err := config.FromYaml("../config.yaml")

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
