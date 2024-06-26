package types

import "go-atm-simulation/model"

type GetAccountInfoRepoInput struct {
	AccountNumber string
}

type GetAccountInfoRepoOutput struct {
	AccountInformationDetail model.AccountInformationDetail
}

type GetAccountInfoServiceInput struct {
	AccountNumber string
}

type GetAccountInfoServiceOutput struct {
	AccountInformationDetail model.AccountInformationDetail
}
