package cmd

import "go-atm-simulation/datasource"

func initData() {
	datasource.Accounts = datasource.New()
}
