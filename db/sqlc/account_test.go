package db

import (
	"avancedGo/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	argCreateAccountParams := CreateAccountParams{
		Owner:    util.RandomOwnerName(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), argCreateAccountParams)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, argCreateAccountParams.Owner, account.Owner)
	require.Equal(t, argCreateAccountParams.Balance, account.Balance)
	require.Equal(t, argCreateAccountParams.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccoutn(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {

	randomAccount1 := createRandomAccount(t)
	sameRandomAccount1, err := testQueries.GetAccount(context.Background(), randomAccount1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, sameRandomAccount1)

	require.Equal(t, randomAccount1.ID, sameRandomAccount1.ID)
	require.Equal(t, randomAccount1.Balance, sameRandomAccount1.Balance)
	require.Equal(t, randomAccount1.Currency, sameRandomAccount1.Currency)
	require.Equal(t, randomAccount1.Owner, sameRandomAccount1.Owner)
	require.Equal(t, randomAccount1.Currency, sameRandomAccount1.Currency)

	require.WithinDuration(t, randomAccount1.CreatedAt, sameRandomAccount1.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	randomAccuount1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      randomAccuount1.ID,
		Balance: util.RandomMoney(),
	}

	err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)

	accountUpdated, err := testQueries.GetAccount(context.Background(), randomAccuount1.ID)

	require.NoError(t, err)
	require.Equal(t, arg.ID, accountUpdated.ID)
	require.Equal(t, arg.Balance, accountUpdated.Balance)
	require.Equal(t, randomAccuount1.Currency, accountUpdated.Currency)
	require.Equal(t, randomAccuount1.Owner, accountUpdated.Owner)
	require.NotEmpty(t, accountUpdated)
}

func TestDeleteAccount(t *testing.T) {
	randomAccount1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), randomAccount1.ID)

	require.NoError(t, err)

	accountInexistent, err := testQueries.GetAccount(context.Background(), randomAccount1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, accountInexistent)

}
