package repository

import (
	"fmt"
	"go-atm-simulation/datasource"
	"go-atm-simulation/model"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetAccountInfoRepo(input model.GetAccountInfoRepoInput) (output model.GetAccountInfoRepoOutput, err error) {
	account, ok := datasource.Accounts[input.AccountNumber]
	if ok {
		output.AccountNumber = account.AccountNumber
		output.Pin = account.Pin
		output.Name = account.Name
		output.Balance = account.Balance

		return output, err
	} else {
		err = fmt.Errorf(model.InvalidAccount)
		return output, err
	}
}

func (r *UserRepository) SubstractBalance(input model.SubstractBalanceInput) (err error) {
	account, ok := datasource.Accounts[input.AccountNumber]
	if ok {
		account.Balance -= input.Amount
		return err
	} else {
		err = fmt.Errorf(model.InvalidAccount)
		return err
	}
}

func (r *UserRepository) AddBalance(input model.AddBalanceInput) (err error) {
	account, ok := datasource.Accounts[input.AccountNumber]
	if ok {
		account.Balance += input.Amount
		return err
	} else {
		err = fmt.Errorf(model.InvalidAccount)
		return err
	}
}
