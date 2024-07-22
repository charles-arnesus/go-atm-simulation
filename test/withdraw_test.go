package test

import (
	"go-atm-simulation/datasource"
	"go-atm-simulation/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithdraw(t *testing.T) {
	datasource.Accounts = datasource.New()
	withdrawInput := model.WithdrawInput{
		AccountNumber: "112233",
		Amount:        "30",
	}
	err := userDelivery.Withdraw(withdrawInput)
	assert.Nil(t, err)
	assert.Equal(t, int64(70), datasource.Accounts["112233"].Balance)
}
