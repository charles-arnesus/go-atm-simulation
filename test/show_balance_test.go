package test

import (
	"go-atm-simulation/datasource"
	"go-atm-simulation/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowBalance(t *testing.T) {
	datasource.Accounts = datasource.New()
	showAccountBalanceInput := model.ShowAccountBalanceInput{
		AccountNumber: "112233",
	}
	err := userDelivery.ShowAccountBalance(showAccountBalanceInput)
	assert.Nil(t, err)
}
