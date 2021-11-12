package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	acc1 := CreateRandomAccount(t)
	acc2 := CreateRandomAccount(t)
	fmt.Println(">>before:", acc1.Balance, acc2.Balance)

	// run n concurrent transfers transactions.
	n := 10
	amount := int64(100)

	errs := make(chan error)

	for i := 0; i < n; i++ {

		FromAccountID := acc1.ID
		ToAccountID := acc2.ID

		if i%2 == 0 {
			FromAccountID = acc2.ID
			ToAccountID = acc1.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: FromAccountID,
				ToAccountID:   ToAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	// Check results

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// check the final updated balances
	updatedAcc1, err := store.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	updatedAcc2, err := store.GetAccount(context.Background(), acc2.ID)
	require.NoError(t, err)

	fmt.Println(">>after:", updatedAcc1.Balance, updatedAcc2.Balance)
	require.Equal(t, acc1.Balance, updatedAcc1.Balance)
	require.Equal(t, acc2.Balance, updatedAcc2.Balance)
}
