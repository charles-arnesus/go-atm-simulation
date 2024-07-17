package model

type AccountInformationDetail struct {
	AccountNumber string
	Name          string
	Pin           string
	Balance       int64
}

type AccountInformation map[string]*AccountInformationDetail

type LoginInput struct {
	AccountNumber string
	Pin           string
}

type LoginOutput struct {
	AccountNumber string
}

type VerifyAccountInput struct {
	AccountNumber string
	Pin           string
}

type VerifyAccountOutput struct {
	AccountNumber string
}

type VerifyAccountInfoRepoInput struct {
	AccountNumber string
	Pin           string
}

type VerifyAccountInfoRepoOutput struct {
	AccountNumber string
}

type ShowAccountBalanceInput struct {
	AccountNumber string
}

type GetAccountInfoInput struct {
	AccountNumber string
}

type GetAccountInfoOutput struct {
	AccountNumber string
	Name          string
	Balance       int64
}

type GetAccountInfoRepoInput struct {
	AccountNumber string
}

type GetAccountInfoRepoOutput struct {
	AccountNumber string
	Name          string
	Balance       int64
}

type WithdrawInput struct {
	AccountNumber string
	Amount        string
}

type WithdrawBalanceInput struct {
	AccountNumber string
	Amount        string
}

type WithdrawBalanceRepoInput struct {
	AccountNumber string
	Amount        string
}
