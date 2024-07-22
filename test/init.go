package test

import (
	"go-atm-simulation/delivery"
	"go-atm-simulation/repository"
	"go-atm-simulation/usecase"
)

var userRepository *repository.UserRepository
var userUseCase *usecase.UserUseCase
var userDelivery *delivery.UserController

func init() {
	userRepository = repository.NewUserRepository()
	userUseCase = usecase.NewUserUseCase(userRepository)
	userDelivery = delivery.NewUserController(userUseCase)
}
