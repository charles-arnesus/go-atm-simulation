package main

import (
	"go-atm-simulation/cmd"
	"go-atm-simulation/datasource"
	"go-atm-simulation/delivery"
	"go-atm-simulation/repository"
	"go-atm-simulation/usecase"
)

func main() {
	userRepository := repository.NewUserRepository()
	userUseCase := usecase.NewUserUseCase(userRepository)
	userDelivery := delivery.NewUserController(userUseCase)
	datasource.Accounts = datasource.New()
	cmd.Start(userDelivery)
}
