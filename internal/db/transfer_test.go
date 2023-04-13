package db

//
//import (
//	"context"
//	"github.com/stretchr/testify/require"
//	"testing"
//)
//
//func createTransferT(t *testing.T) Transfer {
//	arg := CreateTransferParams{
//		FromAccountID: 2,
//		ToAccountID:   3,
//		Amount:        10,
//	}
//
//	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
//	require.NoError(t, err)
//	require.NotEmpty(t, transfer)
//
//	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
//	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
//	require.Equal(t, arg.Amount, transfer.Amount)
//
//	require.NotZero(t, transfer.ID)
//	require.NotZero(t, transfer.CreatedAt)
//
//	return transfer
//}
//
//func TestCreateTransfer(t *testing.T) {
//	createTransferT(t)
//}
