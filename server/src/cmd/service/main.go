package main

import (
	"contest-influence/internal/service"
)

func main() {
	s := service.New("../config.yaml")
	s.Run()
}
