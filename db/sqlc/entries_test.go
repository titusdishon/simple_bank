package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/titusdishon/simple_bank/util"
)

func CreateRandomEntry(t *testing.T) Entry {
	account := CreateRandomAccount(t)
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomAmount(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, account.ID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	CreateRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry := CreateRandomEntry(t)
	entry1, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry1)
	require.Equal(t, entry.ID, entry1.ID)
	require.Equal(t, entry.Amount, entry1.Amount)
	require.WithinDuration(t, entry.CreatedAt, entry1.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t)
	}
	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestUpdateEntry(t *testing.T) {
	entry := CreateRandomEntry(t)
	arg := UpdateEntryParams{
		ID:     entry.ID,
		Amount: util.RandomAmount(),
	}
	entry1, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry1)
	require.Equal(t, entry.ID, entry1.ID)
	require.Equal(t, arg.Amount, entry1.Amount)
	require.Equal(t, entry.AccountID, entry1.AccountID)
	require.WithinDuration(t, entry.CreatedAt, entry1.CreatedAt, time.Second)
}
