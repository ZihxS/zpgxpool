package examples

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgconn"
	zpgxpool "github.com/zihxs/zpgxpool"
)

// TestSuccessfulTransaction demonstrates how to test a successful transaction
func TestSuccessfulTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := zpgxpool.NewMockPgxPool(ctrl)
	mockTx := zpgxpool.NewMockTx(ctrl)

	// Expect Begin to be called and return a mock transaction
	mockPool.EXPECT().
		Begin(gomock.Any()).
		Return(zpgxpool.Tx(mockTx), nil)

	// Expect Exec to be called on the transaction
	mockTx.EXPECT().
		Exec(gomock.Any(), "INSERT INTO users (name, email) VALUES ($1, $2)", "John", "john@example.com").
		Return(pgconn.NewCommandTag("INSERT 0 1"), nil)

	// Expect Commit to be called
	mockTx.EXPECT().
		Commit(gomock.Any()).
		Return(nil)

	// Execute the code under test
	tx, err := mockPool.Begin(context.Background())
	if err != nil {
		t.Fatalf("Failed to begin transaction: %v", err)
	}

	_, err = tx.Exec(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2)", "John", "john@example.com")
	if err != nil {
		t.Fatalf("Failed to execute query: %v", err)
	}

	err = tx.Commit(context.Background())
	if err != nil {
		t.Fatalf("Failed to commit transaction: %v", err)
	}
}
