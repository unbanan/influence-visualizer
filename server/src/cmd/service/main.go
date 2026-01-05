package main

import (
	"contest-influence/server/internal/service"
)

func main() {
	s := service.New("../config.yaml")
	s.Run()
}
