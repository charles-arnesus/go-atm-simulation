package usecase

import (
	"fmt"
	"go-atm-simulation/model"
	"go-atm-simulation/repository"
	"strconv"
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
	getAccountInfoRepoInput := model.GetAccountInfoRepoInput{
		AccountNumber: input.AccountNumber,
	}
	getAccountInfoRepoOutput, err := c.UserRepository.GetAccountInfoRepo(getAccountInfoRepoInput)
	if err != nil {
		return output, err
	}

	if getAccountInfoRepoOutput.Pin != input.Pin {
		err = fmt.Errorf(model.InvalidCredential)
		return output, err
	}

	output = model.VerifyAccountOutput{
		AccountNumber: getAccountInfoRepoOutput.AccountNumber,
	}
	return output, err
}

func (c *UserUseCase) GetAccountInfo(input model.GetAccountInfoInput) (output model.GetAccountInfoOutput, err error) {
	getAccountInfoRepoInput := model.GetAccountInfoRepoInput(input)
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
	//convert transfer amount into int64
	transferAmount, err := strconv.ParseInt(input.Amount, 10, 64)
	if err != nil {
		return err
	}

	//get source account info
	getAccountInfoRepoInput := model.GetAccountInfoRepoInput{
		AccountNumber: input.AccountNumber,
	}
	getAccountInfoRepoOutput, err := c.UserRepository.GetAccountInfoRepo(getAccountInfoRepoInput)
	if err != nil {
		return err
	}

	//check if source account balance is enough to make a transfer
	if transferAmount > getAccountInfoRepoOutput.Balance {
		err = fmt.Errorf(model.InsufficientAmount)
		return err
	}

	//substract source account balance
	substractBalanceInput := model.SubstractBalanceInput{
		AccountNumber: getAccountInfoRepoOutput.AccountNumber,
		Amount:        transferAmount,
	}
	err = c.UserRepository.SubstractBalance(substractBalanceInput)
	if err != nil {
		return err
	}

	return err
}

func (c *UserUseCase) TransferBalance(input model.TransferBalanceInput) (err error) {
	//convert transfer amount into int64
	transferAmount, err := strconv.ParseInt(input.Amount, 10, 64)
	if err != nil {
		return err
	}

	//get destination account info
	getDestinationAccountInfoRepoInput := model.GetAccountInfoRepoInput{
		AccountNumber: input.AccountNumberDestination,
	}
	getDestinationAccountInfoRepoOutput, err := c.UserRepository.GetAccountInfoRepo(getDestinationAccountInfoRepoInput)
	if err != nil {
		return err
	}

	//get source account info
	getSourceAccountInfoRepoInput := model.GetAccountInfoRepoInput{
		AccountNumber: input.AccountNumberSource,
	}
	getSourceAccountInfoRepoOutput, err := c.UserRepository.GetAccountInfoRepo(getSourceAccountInfoRepoInput)
	if err != nil {
		return err
	}

	//check if source account balance is enough to make a transfer
	if transferAmount > getSourceAccountInfoRepoOutput.Balance {
		err = fmt.Errorf(model.InsufficientAmount)
		return err
	}

	//substract source account balance
	substractBalanceInput := model.SubstractBalanceInput{
		AccountNumber: getSourceAccountInfoRepoOutput.AccountNumber,
		Amount:        transferAmount,
	}
	err = c.UserRepository.SubstractBalance(substractBalanceInput)
	if err != nil {
		return err
	}

	//add destination account balance
	addBalanceInput := model.AddBalanceInput{
		AccountNumber: getDestinationAccountInfoRepoOutput.AccountNumber,
		Amount:        transferAmount,
	}
	err = c.UserRepository.AddBalance(addBalanceInput)
	if err != nil {
		return err
	}

	return err

}
