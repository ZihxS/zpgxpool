# Examples

This folder contains various examples of unit tests using the zpgxpool package. These examples demonstrate how to test different scenarios with database operations, including successful queries, error handling, transactions, and edge cases.

## Prerequisites

Before running these examples, make sure you have:

1. Go installed (version 1.19 or later)
2. The necessary dependencies installed

## Running the Examples

To run all the examples, use the following command:

```bash
go test ./...
```

To run a specific example, use:

```bash
go test -v ./<example_file_name>.go
```

For example:

```bash
go test -v ./successful_query_test.go
```

## Example Descriptions

1. **successful_query_test.go** - Demonstrates how to test a successful query execution that returns multiple rows.

2. **successful_transaction_test.go** - Shows how to test a successful database transaction with commit.

3. **query_error_test.go** - Illustrates how to test error handling when a query fails.

4. **transaction_rollback_test.go** - Demonstrates how to test transaction rollback when an error occurs.

5. **empty_query_results_test.go** - Shows how to test queries that return no results.

6. **null_values_test.go** - Illustrates how to handle and test NULL values in query results.

7. **single_row_result_test.go** - Demonstrates how to test queries that return a single row.

8. **nested_transactions_test.go** - Shows how to test nested transactions.

9. **batch_operations_test.go** - Illustrates how to test batch database operations.

## Best Practices Demonstrated

These examples demonstrate several best practices for unit testing database operations:

- Using mocks to isolate the code under test from the actual database
- Verifying interactions with mock expectations
- Testing both success and failure paths
- Handling edge cases like empty results and NULL values
- Proper resource cleanup with `defer` statements
- Clear and descriptive test names
- Comprehensive error handling

## Customization

You can modify these examples to test your own database operations by:

1. Changing the SQL queries to match your use cases
2. Modifying the expected results to match your data structures
3. Adding new test cases for additional scenarios
4. Adjusting the mock expectations to match your code's behavior

Each example is self-contained and can be run independently, making it easy to understand and adapt for your specific needs.