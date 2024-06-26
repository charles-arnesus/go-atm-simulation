package cmd

import (
	"fmt"
	"go-atm-simulation/service"
	"go-atm-simulation/types"
)

func Start(service service.Service) {
	welcome()
	initData()

	////
	getAccountInfoServiceInput := types.GetAccountInfoServiceInput{
		AccountNumber: "112233",
	}
	getAccountInfoOutput, err := service.GetAccountInfoService(getAccountInfoServiceInput)
	if err != nil {

	}
	fmt.Println(getAccountInfoOutput.AccountInformationDetail)
}
