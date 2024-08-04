package db

import (
	"context"
	"testing"
	"time"

	"github.com/JMustang/E_Bank/util"
	"github.com/stretchr/testify/require"
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

func TestCreateTrasnfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccoubtID, transfer2.FromAccoubtID)
	require.Equal(t, transfer1.ToAccoubtID, transfer2.ToAccoubtID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}
