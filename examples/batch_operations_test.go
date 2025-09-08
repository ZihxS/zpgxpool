package examples

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	zpgxpool "github.com/zihxs/zpgxpool"
)

// MockBatchResults is a mock implementation of pgx.BatchResults
type MockBatchResults struct {
	ctrl *gomock.Controller
}

func NewMockBatchResults(ctrl *gomock.Controller) *MockBatchResults {
	return &MockBatchResults{ctrl: ctrl}
}

func (m *MockBatchResults) Exec() (pgconn.CommandTag, error) {
	return pgconn.NewCommandTag("INSERT 0 1"), nil
}

func (m *MockBatchResults) Query() (pgx.Rows, error) {
	return nil, nil
}

func (m *MockBatchResults) QueryRow() pgx.Row {
	return nil
}

func (m *MockBatchResults) Close() error {
	return nil
}

// TestBatchOperations demonstrates how to test batch operations
func TestBatchOperations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPool := zpgxpool.NewMockPgxPool(ctrl)
	mockBatchResults := NewMockBatchResults(ctrl)

	// Create a batch
	batch := &pgx.Batch{}
	batch.Queue("INSERT INTO users (name) VALUES ($1)", "User 1")
	batch.Queue("INSERT INTO users (name) VALUES ($1)", "User 2")

	// Expect SendBatch to be called and return mock results
	mockPool.EXPECT().
		SendBatch(gomock.Any(), batch).
		Return(mockBatchResults)

	// Execute the code under test
	batchResults := mockPool.SendBatch(context.Background(), batch)
	defer batchResults.Close()

	for i := 0; i < 2; i++ {
		_, err := batchResults.Exec()
		if err != nil {
			t.Fatalf("Batch operation %d failed: %v", i, err)
		}
	}
}
