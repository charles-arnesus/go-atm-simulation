package test

import (
	"go-atm-simulation/datasource"
	"go-atm-simulation/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	datasource.Accounts = datasource.New()
	loginInput := model.LoginInput{
		AccountNumber: "112233",
		Pin:           "012108",
	}
	loginOutput, err := userDelivery.Login(loginInput)
	assert.Nil(t, err)
	assert.Equal(t, loginInput.AccountNumber, loginOutput.AccountNumber)
}
