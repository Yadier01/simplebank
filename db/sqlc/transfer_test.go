package db

import (
	"context"
	"testing"
	"time"

	"github.com/Yadier01/simplebank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransfer(t *testing.T) Transfer {

	arg := CreateTransferParams{
		FromAccountID: 1,
		ToAccountID:   2,
		Amount:        util.RandomMoney(),
	}

	account, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.FromAccountID, account.FromAccountID)
	require.Equal(t, arg.ToAccountID, account.ToAccountID)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreateAt)
	return account
}

func TestCreateTransfer(t *testing.T) {
	CreateRandomTransfer(t)
}

func TestGetOneTransfer(t *testing.T) {
	tranfer1 := CreateRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), tranfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, tranfer1.ID, transfer2.ID)
	require.Equal(t, tranfer1.Amount, transfer2.Amount)
	require.Equal(t, tranfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, tranfer1.ToAccountID, transfer2.ToAccountID)

	require.WithinDuration(t, tranfer1.CreateAt, transfer2.CreateAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t)
	}

	arg := ListTransferParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)

	}
}
