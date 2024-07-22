package test

import (
	"go-atm-simulation/datasource"
	"go-atm-simulation/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransfer(t *testing.T) {
	datasource.Accounts = datasource.New()
	transferInput := model.TransferInput{
		AccountNumberSource:      "112233",
		AccountNumberDestination: "112244",
		Amount:                   "30",
	}

	err := userDelivery.Transfer(transferInput)
	assert.Nil(t, err)
	assert.Equal(t, int64(70), datasource.Accounts["112233"].Balance)
	assert.Equal(t, int64(60), datasource.Accounts["112244"].Balance)
}
