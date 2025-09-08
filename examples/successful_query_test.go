package examples

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	zpgxpool "github.com/zihxs/zpgxpool"
)

// User represents a user in the system
type User struct {
	ID    int
	Name  string
	Email string
}

// TestSuccessfulQuery demonstrates how to test a successful query execution
func TestSuccessfulQuery(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := zpgxpool.NewMockPgxPool(ctrl)

	// Create mock rows with expected data
	rows := zpgxpool.NewRows([]string{"id", "name", "email"}).
		AddRow(1, "John Doe", "john@example.com").
		AddRow(2, "Jane Smith", "jane@example.com")

	// Expect the query to be called and return the mock rows
	mockPool.EXPECT().
		Query(gomock.Any(), "SELECT id, name, email FROM users", gomock.Any()).
		Return(rows.ToPgxRows(), nil)

	// Execute the code under test
	resultRows, err := mockPool.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify the results
	var users []User
	for resultRows.Next() {
		var user User
		err := resultRows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			t.Fatalf("Failed to scan row: %v", err)
		}
		users = append(users, user)
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}

	if users[0].Name != "John Doe" {
		t.Errorf("Expected first user to be John Doe, got %s", users[0].Name)
	}
}
