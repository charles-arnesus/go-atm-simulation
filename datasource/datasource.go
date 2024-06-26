package datasource

import "go-atm-simulation/model"

var Accounts model.AccountInformation

func New() (output model.AccountInformation) {
	output = make(model.AccountInformation)

	output["112233"] = model.AccountInformationDetail{
		Name:    "Joh Doe",
		Pin:     "012108",
		Balance: 100,
	}
	output["112244"] = model.AccountInformationDetail{
		Name:    "Jane Doe",
		Pin:     "932012",
		Balance: 30,
	}

	return output
}
