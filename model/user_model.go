package model

const (
	InvalidCredential  string = "invalid account number / pin"
	InvalidAccount     string = "invalid account"
	InvalidAmount      string = "invalid amount"
	InsufficientAmount string = "insufficient amount"
)

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
	Pin           string
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

type TransferInput struct {
	AccountNumberSource      string
	AccountNumberDestination string
	Amount                   string
}

type TransferBalanceInput struct {
	AccountNumberSource      string
	AccountNumberDestination string
	Amount                   string
}

type SubstractBalanceInput struct {
	AccountNumber string
	Amount        int64
}

type AddBalanceInput struct {
	AccountNumber string
	Amount        int64
}
