package main

import (
	"go-atm-simulation/cmd"
	"go-atm-simulation/repository"
	"go-atm-simulation/service"
)

func main() {
	repository := repository.New()
	service := service.New(repository)
	cmd.Start(service)
}
