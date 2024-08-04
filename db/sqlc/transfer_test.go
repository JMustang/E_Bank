package db

import (
	"context"
	"testing"

	"github.com/JMustang/E_Bank/util"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccoubtID: account1.ID,
		ToAccoubtID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	if err != nil {
		t.Error(err)
	}

	return transfer
}
