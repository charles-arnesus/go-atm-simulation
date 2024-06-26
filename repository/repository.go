package repository

import (
	"go-atm-simulation/datasource"
	"go-atm-simulation/types"
)

type Repository interface {
	GetAccountInfoRepo(input types.GetAccountInfoRepoInput) (output types.GetAccountInfoRepoOutput, err error)
}

type repository struct {
}

func New() *repository {
	return &repository{}
}

func (r *repository) GetAccountInfoRepo(input types.GetAccountInfoRepoInput) (output types.GetAccountInfoRepoOutput, err error) {
	output.AccountInformationDetail = datasource.Accounts[input.AccountNumber]
	return output, err
}
