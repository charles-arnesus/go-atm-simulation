package model

type AccountInformationDetail struct {
	Name    string
	Pin     string
	Balance int64
}

type AccountInformation map[string]AccountInformationDetail
