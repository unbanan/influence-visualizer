package service

import (
	"fmt"
	"net/http"
	"os"

	"contest-influence/server/internal/config"
	"contest-influence/server/internal/context"
	"contest-influence/server/internal/handlers"

	"gopkg.in/yaml.v3"
)

type Service struct {
	Config     config.ServiceConfig
	HttpServer http.Server
	Context    *context.Context
}

func New(config_path string) (s Service) {
	content, err := os.ReadFile(config_path)

	fmt.Println("Initialising service")

	if err != nil {
		fmt.Printf("Cannot read config file: %s", err.Error())
		panic(err)
	}

	config := config.ServiceConfig{}
	err = yaml.Unmarshal(content, &config)

	if err != nil {
		fmt.Printf("Cannot parse config file: %s", err.Error())
		panic(err)
	}

	s.Config = config
	s.Context = context.NewContext(config)

	mux := http.NewServeMux()
	mux.Handle("/api/ping", handlers.NewPingHandler())
	mux.Handle("/api/v1/register", handlers.NewRegisterHandler(s.Context))

	s.HttpServer = http.Server{
		Addr:    fmt.Sprintf(":%d", s.Config.Server.Port),
		Handler: mux,
	}

	return
}

func (s *Service) Run() {
	fmt.Printf("Server addr: %s\n", s.HttpServer.Addr)
	fmt.Println("Running server")
	go func() {
		if err := s.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	select {}
}
