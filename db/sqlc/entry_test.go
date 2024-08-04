package db

import (
	"context"
	"testing"

	"github.com/JMustang/E_Bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccoubtID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccoubtID, entry.AccoubtID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}
