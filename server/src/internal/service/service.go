package service

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"contest-influence/server/internal/config"
	"contest-influence/server/internal/handlers"
	"contest-influence/server/internal/repos"
)

type Service interface {
	Run() error
	Shutdown() error
}

type ServiceImpl struct {
	config          config.ServiceConfig
	influenceDBRepo repos.InfluenceDBRepo
	server          *http.Server
}

func New(config config.ServiceConfig) (ServiceImpl, error) {
	fmt.Println("Initialising service")

	influencedb, err := repos.NewInfluenceDBRepo(config.InfluenceDB)

	if err != nil {
		return ServiceImpl{}, err
	}

	return ServiceImpl{
		config:          config,
		influenceDBRepo: influencedb,
	}, nil
}

func (s *ServiceImpl) RunHTTPService() {
	s.server = &http.Server{
		Addr:    s.makeListenerAddress(),
		Handler: s.makeRouter(),
	}

	fmt.Printf("Server addr: %s\n", s.server.Addr)
	fmt.Println("Running server")

	err := s.server.ListenAndServe()

	if err != nil {
		fmt.Printf("failed to server HTTP server %e", err)
	}
}

func (s *ServiceImpl) Run() error {
	go s.RunHTTPService()
	return nil
}

func (s *ServiceImpl) Shutdown() error {
	ctx, close := context.WithTimeout(context.Background(), 15*time.Second)
	defer close()

	return s.server.Shutdown(ctx)
}

func (s *ServiceImpl) makeListenerAddress() string {
	return fmt.Sprintf(":%d", s.config.Server.Port)
}

func (s *ServiceImpl) makeRouter() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/api/ping", handlers.NewPingHandler())
	mux.Handle("/api/v1/register", handlers.NewRegisterHandler(
		regexp.MustCompile(s.config.Common.PlayerNameRegex),
		s.influenceDBRepo,
	))
	return mux
}
