package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/titusdishon/simple_bank/db/util"
)

func createRandomTransfer(t *testing.T) Transfer {
	accountFrom := CreateRandomAccount(t)
	accountTo := CreateRandomAccount(t)
	arg := CreateTransferParams{
		FromAccountID: accountFrom.ID,
		ToAccountID:   accountTo.ID,
		Amount:        util.RandomInt(1, 4),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, accountFrom.ID)
	require.Equal(t, transfer.ToAccountID, accountTo.ID)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	accountFromArg := UpdateAccountParams{
		ID:      accountFrom.ID,
		Balance: accountFrom.Balance - arg.Amount,
	}
	accountFromUpdate, err := testQueries.UpdateAccount(context.Background(), accountFromArg)
	require.NoError(t, err)
	require.NotEmpty(t, accountFromUpdate)
	accountToArg := UpdateAccountParams{
		ID:      accountTo.ID,
		Balance: accountTo.Balance + arg.Amount,
	}
	accountToUpdate, err := testQueries.UpdateAccount(context.Background(), accountToArg)
	require.NoError(t, err)
	require.NotEmpty(t, accountToUpdate)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
	require.Equal(t, transfer.ID, transfer2.ID)
	require.Equal(t, transfer.Amount, transfer2.Amount)
	require.Equal(t, transfer.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer.FromAccountID, transfer2.FromAccountID)
	require.WithinDuration(t, transfer.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}
	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}
	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
