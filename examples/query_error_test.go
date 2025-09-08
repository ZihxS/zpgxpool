package examples

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	zpgxpool "github.com/zihxs/zpgxpool"
)

// TestQueryError demonstrates how to test query error handling
func TestQueryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := zpgxpool.NewMockPgxPool(ctrl)

	// Expect the query to be called and return an error
	expectedError := errors.New("database connection failed")
	mockPool.EXPECT().
		Query(gomock.Any(), "SELECT id, name FROM users", gomock.Any()).
		Return(nil, expectedError)

	// Execute the code under test
	_, err := mockPool.Query(context.Background(), "SELECT id, name FROM users")
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}
