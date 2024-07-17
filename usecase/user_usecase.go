package usecase

import (
	"go-atm-simulation/model"
	"go-atm-simulation/repository"
)

type UserUseCase struct {
	UserRepository *repository.UserRepository
}

func NewUserUseCase(userRepository *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (c *UserUseCase) VerifyAccount(input model.VerifyAccountInput) (output model.VerifyAccountOutput, err error) {
	verifyAccountInfoRepoInput := model.VerifyAccountInfoRepoInput{
		AccountNumber: input.AccountNumber,
		Pin:           input.Pin,
	}
	verifyAccountInfoRepoOutput, err := c.UserRepository.VerifyAccountInfoRepo(verifyAccountInfoRepoInput)
	if err != nil {
		return output, err
	}
	output = model.VerifyAccountOutput{
		AccountNumber: verifyAccountInfoRepoOutput.AccountNumber,
	}
	return output, err
}

func (c *UserUseCase) GetAccountInfo(input model.GetAccountInfoInput) (output model.GetAccountInfoOutput, err error) {
	getAccountInfoRepoInput := model.GetAccountInfoRepoInput{
		AccountNumber: input.AccountNumber,
	}
	getAccountInfoRepoOutput, err := c.UserRepository.GetAccountInfoRepo(getAccountInfoRepoInput)
	if err != nil {
		return output, err
	}

	output.AccountNumber = getAccountInfoRepoOutput.AccountNumber
	output.Name = getAccountInfoRepoOutput.Name
	output.Balance = getAccountInfoRepoOutput.Balance

	return output, err
}

func (c *UserUseCase) WithdrawBalance(input model.WithdrawBalanceInput) (err error) {

	withdrawBalanceRepoInput := model.WithdrawBalanceRepoInput(input)
	err = c.UserRepository.WithdrawBalanceRepo(withdrawBalanceRepoInput)
	return err
}
