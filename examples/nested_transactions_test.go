package examples

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgconn"
	zpgxpool "github.com/zihxs/zpgxpool"
)

// TestNestedTransactions demonstrates how to test nested transactions
func TestNestedTransactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := zpgxpool.NewMockPgxPool(ctrl)
	mockTx1 := zpgxpool.NewMockTx(ctrl)
	mockTx2 := zpgxpool.NewMockTx(ctrl)

	// Expect Begin to be called and return the first mock transaction
	mockPool.EXPECT().
		Begin(gomock.Any()).
		Return(mockTx1, nil)

	// Expect Begin to be called on the first transaction and return the second mock transaction
	mockTx1.EXPECT().
		Begin(gomock.Any()).
		Return(mockTx2, nil)

	// Expect Exec to be called on the nested transaction
	mockTx2.EXPECT().
		Exec(gomock.Any(), "INSERT INTO users (name) VALUES ($1)", "Nested Transaction").
		Return(pgconn.NewCommandTag("INSERT 0 1"), nil)

	// Expect Commit to be called on the nested transaction
	mockTx2.EXPECT().
		Commit(gomock.Any()).
		Return(nil)

	// Expect Commit to be called on the outer transaction
	mockTx1.EXPECT().
		Commit(gomock.Any()).
		Return(nil)

	// Execute the code under test
	tx1, err := mockPool.Begin(context.Background())
	if err != nil {
		t.Fatalf("Failed to begin outer transaction: %v", err)
	}

	tx2, err := tx1.Begin(context.Background())
	if err != nil {
		t.Fatalf("Failed to begin nested transaction: %v", err)
	}

	_, err = tx2.Exec(context.Background(), "INSERT INTO users (name) VALUES ($1)", "Nested Transaction")
	if err != nil {
		t.Fatalf("Failed to execute query in nested transaction: %v", err)
	}

	err = tx2.Commit(context.Background())
	if err != nil {
		t.Fatalf("Failed to commit nested transaction: %v", err)
	}

	err = tx1.Commit(context.Background())
	if err != nil {
		t.Fatalf("Failed to commit outer transaction: %v", err)
	}
}
