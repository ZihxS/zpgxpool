package examples

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	zpgxpool "github.com/zihxs/zpgxpool"
)

// TestNullValues demonstrates how to test NULL values in query results
func TestNullValues(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := zpgxpool.NewMockPgxPool(ctrl)

	// Create mock rows with NULL values
	rows := zpgxpool.NewRows([]string{"id", "name", "email"}).
		AddRow(1, "John Doe", nil) // email is NULL

	// Expect the query to be called and return rows with NULL values
	mockPool.EXPECT().
		Query(gomock.Any(), "SELECT id, name, email FROM users WHERE id = $1", 1).
		Return(rows.ToPgxRows(), nil)

	// Execute the code under test
	resultRows, err := mockPool.Query(context.Background(), "SELECT id, name, email FROM users WHERE id = $1", 1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify the results with NULL values
	if !resultRows.Next() {
		t.Fatal("Expected a row, but got none")
	}

	var id int
	var name string
	var email *string // pointer to handle NULL values

	err = resultRows.Scan(&id, &name, &email)
	if err != nil {
		t.Fatalf("Failed to scan row: %v", err)
	}

	if id != 1 {
		t.Errorf("Expected id 1, got %d", id)
	}

	if name != "John Doe" {
		t.Errorf("Expected name John Doe, got %s", name)
	}

	if email != nil {
		t.Error("Expected email to be NULL")
	}
}
