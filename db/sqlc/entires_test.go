package db

import (
	"context"
	"testing"
	"time"

	"github.com/Yadier01/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomEntry(t *testing.T) Entry {

	arg := CreateEntryParams{
		AccountID: 5,
		Amount:    util.RandomMoney(),
	}

	account, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.AccountID, account.AccountID)
	require.Equal(t, arg.Amount, account.Amount)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreateAt)
	return account
}
func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetOneEntry(t *testing.T) {
	entry1 := CreateRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)

	require.WithinDuration(t, entry1.CreateAt, entry2.CreateAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)

	}
}
