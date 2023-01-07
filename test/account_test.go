package test

import (
	"context"
	"database/sql"
	db "github.com/quandat10/banking/db/sqlc"
	"github.com/quandat10/banking/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) db.Account {
	arg := db.CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomConcurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	newAccount := createRandomAccount(t)

	testAccount, err := testQueries.GetAccount(context.Background(), newAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, testAccount)

	require.Equal(t, newAccount.ID, testAccount.ID)
	require.Equal(t, newAccount.Owner, testAccount.Owner)
	require.Equal(t, newAccount.Balance, testAccount.Balance)
	require.Equal(t, newAccount.Currency, testAccount.Currency)
	require.WithinDuration(t, newAccount.CreatedAt, testAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	newAccount := createRandomAccount(t)

	arg := db.UpdateAccountParams{
		ID:      newAccount.ID,
		Balance: util.RandomMoney(),
	}

	testAccount, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, testAccount)

	require.Equal(t, newAccount.ID, testAccount.ID)
	require.Equal(t, newAccount.Owner, testAccount.Owner)
	require.Equal(t, arg.Balance, testAccount.Balance)
	require.Equal(t, newAccount.Currency, testAccount.Currency)
	require.WithinDuration(t, newAccount.CreatedAt, testAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	newAccount := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), newAccount.ID)

	require.NoError(t, err)

	testAccount, err := testQueries.GetAccount(context.Background(), newAccount.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, testAccount)
}

func TestListAccount(t *testing.T) {
	for i := 0; i <= 10; i++ {
		createRandomAccount(t)
	}

	arg := db.ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
