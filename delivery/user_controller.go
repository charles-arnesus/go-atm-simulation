package delivery

import (
	"fmt"
	"go-atm-simulation/model"
	"go-atm-simulation/usecase"
)

type UserController struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserController(userUseCase *usecase.UserUseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (c *UserController) Login(input model.LoginInput) (output model.LoginOutput, err error) {

	verifyAccountInput := model.VerifyAccountInput(input)
	verifyAccountOutput, err := c.UserUseCase.VerifyAccount(verifyAccountInput)
	if err != nil {
		return output, err
	}

	output.AccountNumber = verifyAccountOutput.AccountNumber

	return output, err

}

func (c *UserController) ShowAccountBalance(input model.ShowAccountBalanceInput) (err error) {

	getAccountInfoInput := model.GetAccountInfoInput(input)
	getAccountInfoOutput, err := c.UserUseCase.GetAccountInfo(getAccountInfoInput)
	if err != nil {
		return err
	}

	fmt.Printf("Account Number: %s\n", getAccountInfoOutput.AccountNumber)
	fmt.Printf("Name: %s\n", getAccountInfoOutput.Name)
	fmt.Printf("Balance: %d\n", getAccountInfoOutput.Balance)

	return err
}

func (c *UserController) Withdraw(input model.WithdrawInput) (err error) {

	withdrawBalanceInput := model.WithdrawBalanceInput(input)
	err = c.UserUseCase.WithdrawBalance(withdrawBalanceInput)
	if err != nil {
		return err
	}
	fmt.Println("Withdraw Success")
	return err
}
