package repository

import (
	"fmt"
	"go-atm-simulation/datasource"
	"go-atm-simulation/model"
	"strconv"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) VerifyAccountInfoRepo(input model.VerifyAccountInfoRepoInput) (output model.VerifyAccountInfoRepoOutput, err error) {
	account, ok := datasource.Accounts[input.AccountNumber]
	if ok {
		if account.Pin != input.Pin {
			err = fmt.Errorf(datasource.InvalidCredential)
			return output, err
		}
		output.AccountNumber = account.AccountNumber

		return output, err
	} else {
		err = fmt.Errorf(datasource.InvalidCredential)
		return output, err
	}
}

func (r *UserRepository) GetAccountInfoRepo(input model.GetAccountInfoRepoInput) (output model.GetAccountInfoRepoOutput, err error) {
	account, ok := datasource.Accounts[input.AccountNumber]
	if ok {
		output.AccountNumber = account.AccountNumber
		output.Name = account.Name
		output.Balance = account.Balance

		return output, err
	} else {
		err = fmt.Errorf(datasource.InvalidAccount)
		return output, err
	}
}

func (r *UserRepository) WithdrawBalanceRepo(input model.WithdrawBalanceRepoInput) (err error) {
	account, ok := datasource.Accounts[input.AccountNumber]
	if ok {
		amount, err := strconv.ParseInt(input.Amount, 10, 64)
		if err != nil {
			return err
		}
		if amount > account.Balance {
			err = fmt.Errorf(datasource.InvalidAmount)
			return err
		}
		account.Balance -= amount

		return err
	} else {
		err = fmt.Errorf(datasource.InvalidAccount)
		return err
	}
}
