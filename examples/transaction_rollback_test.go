package examples

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgconn"
	zpgxpool "github.com/zihxs/zpgxpool"
)

// TestTransactionRollbackOnError demonstrates how to test transaction rollback on error
func TestTransactionRollbackOnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := zpgxpool.NewMockPgxPool(ctrl)
	mockTx := zpgxpool.NewMockTx(ctrl)

	// Expect Begin to be called and return a mock transaction
	mockPool.EXPECT().
		Begin(gomock.Any()).
		Return(mockTx, nil)

	// Expect Exec to be called and return an error
	expectedError := errors.New("insert failed")
	mockTx.EXPECT().
		Exec(gomock.Any(), "INSERT INTO users (name) VALUES ($1)", "John").
		Return(pgconn.NewCommandTag(""), expectedError)

	// Expect Rollback to be called
	mockTx.EXPECT().
		Rollback(gomock.Any()).
		Return(nil)

	// Execute the code under test
	tx, err := mockPool.Begin(context.Background())
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(context.Background(), "INSERT INTO users (name) VALUES ($1)", "John")
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}
