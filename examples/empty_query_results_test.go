package examples

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	zpgxpool "github.com/zihxs/zpgxpool"
)

// TestEmptyQueryResults demonstrates how to test empty query results
func TestEmptyQueryResults(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := zpgxpool.NewMockPgxPool(ctrl)

	// Create empty mock rows
	rows := zpgxpool.NewRows([]string{"id", "name"})

	// Expect the query to be called and return empty rows
	mockPool.EXPECT().
		Query(gomock.Any(), "SELECT id, name FROM users WHERE id = $1", 999).
		Return(rows.ToPgxRows(), nil)

	// Execute the code under test
	resultRows, err := mockPool.Query(context.Background(), "SELECT id, name FROM users WHERE id = $1", 999)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify that no rows are returned
	if resultRows.Next() {
		t.Error("Expected no rows, but got at least one")
	}
}
