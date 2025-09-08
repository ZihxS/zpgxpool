package examples

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	zpgxpool "github.com/zihxs/zpgxpool"
)

// TestSingleRowResult demonstrates how to test single row result
func TestSingleRowResult(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := zpgxpool.NewMockPgxPool(ctrl)

	// Create mock row
	row := zpgxpool.NewRow([]string{"id", "name"}, 1, "John Doe")

	// Expect the query to be called and return a single row
	mockPool.EXPECT().
		QueryRow(gomock.Any(), "SELECT id, name FROM users WHERE id = $1", 1).
		Return(row)

	// Execute the code under test
	resultRow := mockPool.QueryRow(context.Background(), "SELECT id, name FROM users WHERE id = $1", 1)

	var id int
	var name string
	err := resultRow.Scan(&id, &name)
	if err != nil {
		t.Fatalf("Failed to scan row: %v", err)
	}

	if id != 1 {
		t.Errorf("Expected id 1, got %d", id)
	}

	if name != "John Doe" {
		t.Errorf("Expected name John Doe, got %s", name)
	}
}
