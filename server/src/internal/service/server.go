package service

import (
	"fmt"
	"net/http"
	"os"

	"contest-influence/internal/handlers"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port uint16 `yaml:"port"`
	}

	InfluenceDB struct {
		DSN string `yaml:"influencedb.dsn"`
	}
}

type server struct {
	Config     Config
	HttpServer http.Server
}

func New(config_path string) (s server) {
	content, err := os.ReadFile(config_path)

	fmt.Println("Initialising service")

	if err != nil {
		fmt.Printf("Cannot read config file: %s", err.Error())
		panic(err)
	}

	err = yaml.Unmarshal(content, &s.Config)

	if err != nil {
		fmt.Printf("Cannot parse config file: %s", err.Error())
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/ping", &handlers.PingHandler{})
	s.HttpServer = http.Server{
		Addr:    fmt.Sprintf(":%d", s.Config.Server.Port),
		Handler: mux,
	}

	return
}

func (s *server) Run() {
	fmt.Printf("Server addr: %s\n", s.HttpServer.Addr)
	fmt.Println("Running server")
	go func() {
		if err := s.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	select {}
}
