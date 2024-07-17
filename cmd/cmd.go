package cmd

import (
	"bufio"
	"fmt"
	"go-atm-simulation/delivery"
	"go-atm-simulation/model"
	"os"
	"strings"
)

func Start(userDelivery *delivery.UserController) {
	fmt.Println("Welcome to ABC!")

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("")
		fmt.Println("Enter Account Number:")
		fmt.Print("->")

		accountNumber, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		accountNumber = strings.Replace(accountNumber, "\n", "", -1)

		fmt.Println("Enter Password:")
		fmt.Print("->")

		pin, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		pin = strings.Replace(pin, "\n", "", -1)

		fmt.Println(accountNumber)
		fmt.Println(pin)

		loginInput := model.LoginInput{
			AccountNumber: accountNumber,
			Pin:           pin,
		}

		loginOutput, err := userDelivery.Login(loginInput)
		if err != nil {
			fmt.Println(err)
			break
		}

		showMenu()

		menu, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			break
		}
		menu = strings.Replace(menu, "\n", "", -1)

		switch menu {
		case "1":
			//Check Balance
			showAccountBalanceInput := model.ShowAccountBalanceInput(loginOutput)
			err := userDelivery.ShowAccountBalance(showAccountBalanceInput)
			if err != nil {
				fmt.Println(err)
				break
			}
			continue
		case "2":
			//Withdraw
			fmt.Println("Enter Withdraw Amount:")
			fmt.Print("->")

			withdrawAmount, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				break
			}
			withdrawAmount = strings.Replace(withdrawAmount, "\n", "", -1)
			withdrawInput := model.WithdrawInput{
				AccountNumber: loginOutput.AccountNumber,
				Amount:        withdrawAmount,
			}
			err = userDelivery.Withdraw(withdrawInput)
			if err != nil {
				fmt.Println(err)
				break
			}
			continue
		case "3":
			//Transfer
			fmt.Println("Enter Account Number Destination:")
			fmt.Print("->")

			accountNumberDestination, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				break
			}
			accountNumberDestination = strings.Replace(accountNumberDestination, "\n", "", -1)

			fmt.Println("Enter Transfer Amount:")
			fmt.Print("->")

			transferAmount, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				break
			}
			transferAmount = strings.Replace(transferAmount, "\n", "", -1)

			transferInput := model.TransferInput{
				AccountNumberSource:      loginOutput.AccountNumber,
				AccountNumberDestination: accountNumberDestination,
				Amount:                   transferAmount,
			}

			err = userDelivery.Transfer(transferInput)
			if err != nil {
				fmt.Println(err)
				break
			}

			continue
		default:
			fmt.Println("Thank you for using ABC")
		}

		break
	}
}

func showMenu() {
	fmt.Println("Menu:")
	fmt.Println("1.Check Balance")
	fmt.Println("2.Withdraw")
	fmt.Println("3.Transfer")
	fmt.Print("->")
}
