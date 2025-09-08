# 📚 zpgxpool Documentation

## Table of Contents
1. [Introduction](#introduction)
2. [PgxPool Interface](#pgxpool-interface)
   - [Methods](#pgxpool-methods)
3. [Tx Interface](#tx-interface)
   - [Methods](#tx-methods)
4. [Mocking Functions](#mocking-functions)
   - [Row](#row)
   - [Rows](#rows)
5. [Usage Examples](#usage-examples)
   - [Basic Usage](#basic-usage)
   - [Transaction Usage](#transaction-usage)
   - [Mocking in Tests](#mocking-in-tests)
6. [Comprehensive Unit Testing Guide](#comprehensive-unit-testing-guide)
   - [Positive Test Cases](#positive-test-cases)
   - [Negative Test Cases](#negative-test-cases)
   - [Edge Case Testing](#edge-case-testing)
   - [SQL Transaction Testing](#sql-transaction-testing)
   - [Advanced Mocking Scenarios](#advanced-mocking-scenarios)
   - [Best Practices for Testing](#best-practices-for-testing)

---

# 📚 Dokumentasi zpgxpool

## Daftar Isi
1. [Pengantar](#pengantar)
2. [Interface PgxPool](#interface-pgxpool)
   - [Method-method](#method-pgxpool)
3. [Interface Tx](#interface-tx)
   - [Method-method](#method-tx)
4. [Fungsi Mocking](#fungsi-mocking)
   - [Row](#row)
   - [Rows](#rows)
5. [Contoh Penggunaan](#contoh-penggunaan)
   - [Penggunaan Dasar](#penggunaan-dasar)
   - [Penggunaan Transaksi](#penggunaan-transaksi)
   - [Mocking dalam Test](#mocking-dalam-test)
6. [Panduan Komprehensif Unit Testing](#panduan-komprehensif-unit-testing)
   - [Kasus Uji Positif](#kasus-uji-positif)
   - [Kasus Uji Negatif](#kasus-uji-negatif)
   - [Menguji Kasus Tepi](#menguji-kasus-tepi)
   - [Menguji Transaksi SQL](#menguji-transaksi-sql)
   - [Skenario Mocking Lanjutan](#skenario-mocking-lanjutan)
   - [Praktik Terbaik untuk Testing](#praktik-terbaik-untuk-testing)

---

## 📖 Introduction / Pengantar

The `zpgxpool` package provides a wrapper around the `pgx/v5/pgxpool` package with additional mocking capabilities for testing. It defines interfaces that make it easier to mock database operations in tests.

Package `zpgxpool` menyediakan wrapper untuk package `pgx/v5/pgxpool` dengan kemampuan mocking tambahan untuk testing. Package ini mendefinisikan interface-interface yang memudahkan mocking operasi database dalam test.

---

## 🔄 PgxPool Interface

The `PgxPool` interface wraps the functionality of `pgx/v5/pgxpool.Pool` and provides methods for database operations.

Interface `PgxPool` membungkus fungsionalitas dari `pgx/v5/pgxpool.Pool` dan menyediakan method-method untuk operasi database.

### 📋 PgxPool Methods / Method PgxPool

#### Acquire
```go
Acquire(ctx context.Context) (*pgxpool.Conn, error)
```
Acquires a connection from the pool. It returns an error if the pool is closed or the context is canceled.

Mendapatkan koneksi dari pool. Mengembalikan error jika pool ditutup atau context dibatalkan.

#### AcquireAllIdle
```go
AcquireAllIdle(ctx context.Context) []*pgxpool.Conn
```
Acquires all idle connections from the pool. This can be used to empty the pool.

Mendapatkan semua koneksi yang tidak aktif dari pool. Dapat digunakan untuk mengosongkan pool.

#### AcquireFunc
```go
AcquireFunc(ctx context.Context, f func(*pgxpool.Conn) error) error
```
Acquires a connection from the pool, executes the provided function, and releases the connection back to the pool.

Mendapatkan koneksi dari pool, menjalankan fungsi yang diberikan, dan melepaskan koneksi kembali ke pool.

#### Begin
```go
Begin(ctx context.Context) (Tx, error)
```
Begins a transaction. It returns a `Tx` interface that can be used to perform transaction operations.

Memulai transaksi. Mengembalikan interface `Tx` yang dapat digunakan untuk melakukan operasi transaksi.

#### BeginTx
```go
BeginTx(ctx context.Context, txOptions pgx.TxOptions) (Tx, error)
```
Begins a transaction with the specified transaction options.

Memulai transaksi dengan opsi transaksi yang ditentukan.

#### Close
```go
Close()
```
Closes all connections in the pool and prevents new connections from being acquired.

Menutup semua koneksi dalam pool dan mencegah koneksi baru dari diambil.

#### Config
```go
Config() *pgxpool.Config
```
Returns the pool configuration.

Mengembalikan konfigurasi pool.

#### CopyFrom
```go
CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
```
Executes a COPY FROM command to efficiently insert large amounts of data.

Menjalankan perintah COPY FROM untuk memasukkan data dalam jumlah besar secara efisien.

#### Exec
```go
Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
```
Executes a SQL command that does not return rows.

Menjalankan perintah SQL yang tidak mengembalikan baris.

#### Ping
```go
Ping(ctx context.Context) error
```
Checks the connection to the database.

Memeriksa koneksi ke database.

#### Query
```go
Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
```
Executes a SQL query that returns rows.

Menjalankan query SQL yang mengembalikan baris.

#### QueryRow
```go
QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
```
Executes a SQL query that is expected to return at most one row.

Menjalankan query SQL yang diharapkan mengembalikan paling banyak satu baris.

#### Reset
```go
Reset()
```
Resets the pool by closing all connections and creating new ones according to the configuration.

Mereset pool dengan menutup semua koneksi dan membuat koneksi baru sesuai konfigurasi.

#### SendBatch
```go
SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
```
Sends all queued queries to the server at once.

Mengirim semua query yang antri ke server sekaligus.

#### Stat
```go
Stat() *pgxpool.Stat
```
Returns pool statistics.

Mengembalikan statistik pool.

---

## 🔄 Tx Interface

The `Tx` interface represents a database transaction and provides methods for transaction operations.

Interface `Tx` merepresentasikan transaksi database dan menyediakan method-method untuk operasi transaksi.

### 📋 Tx Methods / Method Tx

#### Begin
```go
Begin(ctx context.Context) (Tx, error)
```
Begins a nested transaction.

Memulai transaksi bersarang.

#### Commit
```go
Commit(ctx context.Context) error
```
Commits the transaction.

Melakukan commit transaksi.

#### Rollback
```go
Rollback(ctx context.Context) error
```
Rolls back the transaction.

Membatalkan transaksi.

#### CopyFrom
```go
CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
```
Executes a COPY FROM command within the transaction.

Menjalankan perintah COPY FROM dalam transaksi.

#### SendBatch
```go
SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
```
Sends all queued queries to the server at once within the transaction.

Mengirim semua query yang antri ke server sekaligus dalam transaksi.

#### LargeObjects
```go
LargeObjects() pgx.LargeObjects
```
Returns a handle to the large objects API.

Mengembalikan handle ke API large objects.

#### Prepare
```go
Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
```
Prepares a statement for use within the transaction.

Mempersiapkan statement untuk digunakan dalam transaksi.

#### Exec
```go
Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
```
Executes a SQL command that does not return rows within the transaction.

Menjalankan perintah SQL yang tidak mengembalikan baris dalam transaksi.

#### Query
```go
Query(ctx context.Context, sql string, args ...any) (Rows, error)
```
Executes a SQL query that returns rows within the transaction.

Menjalankan query SQL yang mengembalikan baris dalam transaksi.

#### QueryRow
```go
QueryRow(ctx context.Context, sql string, args ...any) Row
```
Executes a SQL query that is expected to return at most one row within the transaction.

Menjalankan query SQL yang diharapkan mengembalikan paling banyak satu baris dalam transaksi.

#### Conn
```go
Conn() *pgx.Conn
```
Returns the underlying connection.

Mengembalikan koneksi yang mendasarinya.

---

## 🎭 Mocking Functions / Fungsi Mocking

The package provides several functions for creating mock database objects for testing.

Package ini menyediakan beberapa fungsi untuk membuat objek database mock untuk testing.

### 🟢 Row

#### NewRow
```go
func NewRow(columns []string, values ...any) *Row
```
Creates a new Row with the specified columns and values.

Membuat Row baru dengan kolom dan nilai yang ditentukan.

#### Scan
```go
func (r *Row) Scan(dest ...any) error
```
Scans the row values into the destination variables.

Memindai nilai baris ke variabel tujuan.

#### 🟢 RowError
```go
func (r *Row) RowError(row int, err error) *Row
```
Sets an error that will be returned when the row is read.

Menetapkan error yang akan dikembalikan saat baris dibaca.

### 🟢 Rows

#### NewRows
```go
func NewRows(columns []string) *Rows
```
Creates a new Rows with the specified columns.

Membuat Rows baru dengan kolom yang ditentukan.

#### NewRowsWithColumnDefinition
```go
func NewRowsWithColumnDefinition(columns ...pgconn.FieldDescription) *Rows
```
Creates a new Rows with the specified column definitions.

Membuat Rows baru dengan definisi kolom yang ditentukan.

#### AddRow
```go
func (r *Rows) AddRow(values ...any) *Rows
```
Adds a row with the specified values.

Menambahkan baris dengan nilai yang ditentukan.

#### FromCSVString
```go
func (r *Rows) FromCSVString(s string) *Rows
```
Builds rows from a CSV string.

Membangun baris dari string CSV.

#### CloseError
```go
func (r *Rows) CloseError(err error) *Rows
```
Sets an error that will be returned by rows.Close.

Menetapkan error yang akan dikembalikan oleh rows.Close.

#### 🟢 RowError
```go
func (r *Rows) RowError(row int, err error) *Rows
```
Sets an error that will be returned when a given row number is read.

Menetapkan error yang akan dikembalikan saat nomor baris tertentu dibaca.

#### ToPgxRows
```go
func (r *Rows) ToPgxRows() pgx.Rows
```
Converts the mock rows to pgx.Rows.

Mengonversi baris mock ke pgx.Rows.

---

## 💡 Usage Examples / Contoh Penggunaan

### 🚀 Basic Usage / Penggunaan Dasar

```go
// Create a connection pool
// Membuat connection pool
pool, err := pgxpool.New(context.Background(), "your-connection-string")
if err != nil {
    log.Fatal(err)
}
defer pool.Close()

// Use the pool through the interface
// Menggunakan pool melalui interface
var db zpgxpool.PgxPool = pool

// Execute a query
// Menjalankan query
row := db.QueryRow(context.Background(), "SELECT id, name FROM users WHERE id = $1", 1)
var id int
var name string
err = row.Scan(&id, &name)
if err != nil {
    log.Fatal(err)
}
```

### 🔄 Transaction Usage / Penggunaan Transaksi

```go
// Begin a transaction
// Memulai transaksi
tx, err := db.Begin(context.Background())
if err != nil {
    log.Fatal(err)
}
defer tx.Rollback(context.Background())

// Execute queries within the transaction
// Menjalankan query dalam transaksi
_, err = tx.Exec(context.Background(), "INSERT INTO users (name) VALUES ($1)", "John")
if err != nil {
    log.Fatal(err)
}

// Commit the transaction
// Melakukan commit transaksi
err = tx.Commit(context.Background())
if err != nil {
    log.Fatal(err)
}
```

### 🎭 Mocking in Tests / Mocking dalam Test

```go
func TestUserService(t *testing.T) {
    // Create a mock pool
    // Membuat mock pool
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPool := zpgxpool.NewMockPgxPool(ctrl)

    // Set up expectations
    // Menyiapkan ekspektasi
    rows := zpgxpool.NewRows([]string{"id", "name"}).
        AddRow(1, "John").
        AddRow(2, "Jane")

    mockPool.EXPECT().
        Query(gomock.Any(), "SELECT id, name FROM users", gomock.Any()).
        Return(rows.ToPgxRows(), nil)

    // Use the mock in your code
    // Menggunakan mock dalam kode Anda
    // ... your test code ...
```

---

## 🧪 Comprehensive Unit Testing Guide / Panduan Komprehensif Unit Testing

### ✅ Positive Test Cases / Kasus Uji Positif

Positive test cases verify that the system behaves correctly when given valid inputs.

Kasus uji positif memverifikasi bahwa sistem berperilaku dengan benar ketika diberikan input yang valid.

#### Testing Successful Query Execution / Menguji Eksekusi Query yang Berhasil

```go
func TestSuccessfulQuery(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPool := zpgxpool.NewMockPgxPool(ctrl)

    // Create mock rows with expected data
    // Membuat baris mock dengan data yang diharapkan
    rows := zpgxpool.NewRows([]string{"id", "name", "email"}).
        AddRow(1, "John Doe", "john@example.com").
        AddRow(2, "Jane Smith", "jane@example.com")

    // Expect the query to be called and return the mock rows
    // Mengharapkan query dipanggil dan mengembalikan baris mock
    mockPool.EXPECT().
        Query(gomock.Any(), "SELECT id, name, email FROM users", gomock.Any()).
        Return(rows.ToPgxRows(), nil)

    // Execute the code under test
    // Menjalankan kode yang diuji
    resultRows, err := mockPool.Query(context.Background(), "SELECT id, name, email FROM users")
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    // Verify the results
    // Memverifikasi hasil
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
```

#### Testing Successful Transaction / Menguji Transaksi yang Berhasil

```go
func TestSuccessfulTransaction(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPool := zpgxpool.NewMockPgxPool(ctrl)
    mockTx := zpgxpool.NewMockTx(ctrl)

    // Expect Begin to be called and return a mock transaction
    // Mengharapkan Begin dipanggil dan mengembalikan transaksi mock
    mockPool.EXPECT().
        Begin(gomock.Any()).
        Return(mockTx, nil)

    // Expect Exec to be called on the transaction
    // Mengharapkan Exec dipanggil pada transaksi
    mockTx.EXPECT().
        Exec(gomock.Any(), "INSERT INTO users (name, email) VALUES ($1, $2)", "John", "john@example.com").
        Return(pgconn.NewCommandTag("INSERT 0 1"), nil)

    // Expect Commit to be called
    // Mengharapkan Commit dipanggil
    mockTx.EXPECT().
        Commit(gomock.Any()).
        Return(nil)

    // Execute the code under test
    // Menjalankan kode yang diuji
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
```

### ❌ Negative Test Cases / Kasus Uji Negatif

Negative test cases verify that the system handles errors and invalid inputs gracefully.

Kasus uji negatif memverifikasi bahwa sistem menangani error dan input yang tidak valid dengan baik.

#### Testing Query Error / Menguji Error Query

```go
func TestQueryError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPool := zpgxpool.NewMockPgxPool(ctrl)

    // Expect the query to be called and return an error
    // Mengharapkan query dipanggil dan mengembalikan error
    expectedError := errors.New("database connection failed")
    mockPool.EXPECT().
        Query(gomock.Any(), "SELECT id, name FROM users", gomock.Any()).
        Return(nil, expectedError)

    // Execute the code under test
    // Menjalankan kode yang diuji
    _, err := mockPool.Query(context.Background(), "SELECT id, name FROM users")
    if err == nil {
        t.Fatal("Expected an error, got nil")
    }

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, got %v", expectedError, err)
    }
}
```

#### Testing Transaction Rollback on Error / Menguji Rollback Transaksi saat Error

```go
func TestTransactionRollbackOnError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPool := zpgxpool.NewMockPgxPool(ctrl)
    mockTx := zpgxpool.NewMockTx(ctrl)

    // Expect Begin to be called and return a mock transaction
    // Mengharapkan Begin dipanggil dan mengembalikan transaksi mock
    mockPool.EXPECT().
        Begin(gomock.Any()).
        Return(mockTx, nil)

    // Expect Exec to be called and return an error
    // Mengharapkan Exec dipanggil dan mengembalikan error
    expectedError := errors.New("insert failed")
    mockTx.EXPECT().
        Exec(gomock.Any(), "INSERT INTO users (name) VALUES ($1)", "John").
        Return(pgconn.NewCommandTag(""), expectedError)

    // Expect Rollback to be called
    // Mengharapkan Rollback dipanggil
    mockTx.EXPECT().
        Rollback(gomock.Any()).
        Return(nil)

    // Execute the code under test
    // Menjalankan kode yang diuji
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
```

### ⚠️ Edge Case Testing / Menguji Kasus Tepi

Edge case testing verifies behavior at the boundaries of expected inputs.

Menguji kasus tepi memverifikasi perilaku pada batas input yang diharapkan.

#### Testing Empty Query Results / Menguji Hasil Query Kosong

```go
func TestEmptyQueryResults(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPool := zpgxpool.NewMockPgxPool(ctrl)

    // Create empty mock rows
    // Membuat baris mock kosong
    rows := zpgxpool.NewRows([]string{"id", "name"})

    // Expect the query to be called and return empty rows
    // Mengharapkan query dipanggil dan mengembalikan baris kosong
    mockPool.EXPECT().
        Query(gomock.Any(), "SELECT id, name FROM users WHERE id = $1", 999).
        Return(rows.ToPgxRows(), nil)

    // Execute the code under test
    // Menjalankan kode yang diuji
    resultRows, err := mockPool.Query(context.Background(), "SELECT id, name FROM users WHERE id = $1", 999)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    // Verify that no rows are returned
    // Memverifikasi bahwa tidak ada baris yang dikembalikan
    if resultRows.Next() {
        t.Error("Expected no rows, but got at least one")
    }
}
```

#### Testing NULL Values / Menguji Nilai NULL

```go
func TestNullValues(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPool := zpgxpool.NewMockPgxPool(ctrl)

    // Create mock rows with NULL values
    // Membuat baris mock dengan nilai NULL
    rows := zpgxpool.NewRows([]string{"id", "name", "email"}).
        AddRow(1, "John Doe", nil) // email is NULL

    // Expect the query to be called and return rows with NULL values
    // Mengharapkan query dipanggil dan mengembalikan baris dengan nilai NULL
    mockPool.EXPECT().
        Query(gomock.Any(), "SELECT id, name, email FROM users WHERE id = $1", 1).
        Return(rows.ToPgxRows(), nil)

    // Execute the code under test
    // Menjalankan kode yang diuji
    resultRows, err := mockPool.Query(context.Background(), "SELECT id, name, email FROM users WHERE id = $1", 1)
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    // Verify the results with NULL values
    // Memverifikasi hasil dengan nilai NULL
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
```

#### Testing Single Row Result / Menguji Hasil Baris Tunggal

```go
func TestSingleRowResult(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPool := zpgxpool.NewMockPgxPool(ctrl)

    // Create mock row
    // Membuat baris mock
    row := zpgxpool.NewRow([]string{"id", "name"}, 1, "John Doe")

    // Expect the query to be called and return a single row
    // Mengharapkan query dipanggil dan mengembalikan baris tunggal
    mockPool.EXPECT().
        QueryRow(gomock.Any(), "SELECT id, name FROM users WHERE id = $1", 1).
        Return(row)

    // Execute the code under test
    // Menjalankan kode yang diuji
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
```

### 🔄 SQL Transaction Testing / Menguji Transaksi SQL

Testing SQL transactions requires careful attention to commit/rollback behavior and nested transactions.

Menguji transaksi SQL memerlukan perhatian khusus pada perilaku commit/rollback dan transaksi bersarang.

#### Testing Nested Transactions / Menguji Transaksi Bersarang

```go
func TestNestedTransactions(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPool := zpgxpool.NewMockPgxPool(ctrl)
    mockTx1 := zpgxpool.NewMockTx(ctrl)
    mockTx2 := zpgxpool.NewMockTx(ctrl)

    // Expect Begin to be called and return the first mock transaction
    // Mengharapkan Begin dipanggil dan mengembalikan transaksi mock pertama
    mockPool.EXPECT().
        Begin(gomock.Any()).
        Return(mockTx1, nil)

    // Expect Begin to be called on the first transaction and return the second mock transaction
    // Mengharapkan Begin dipanggil pada transaksi pertama dan mengembalikan transaksi mock kedua
    mockTx1.EXPECT().
        Begin(gomock.Any()).
        Return(mockTx2, nil)

    // Expect Exec to be called on the nested transaction
    // Mengharapkan Exec dipanggil pada transaksi bersarang
    mockTx2.EXPECT().
        Exec(gomock.Any(), "INSERT INTO users (name) VALUES ($1)", "Nested Transaction").
        Return(pgconn.NewCommandTag("INSERT 0 1"), nil)

    // Expect Commit to be called on the nested transaction
    // Mengharapkan Commit dipanggil pada transaksi bersarang
    mockTx2.EXPECT().
        Commit(gomock.Any()).
        Return(nil)

    // Expect Commit to be called on the outer transaction
    // Mengharapkan Commit dipanggil pada transaksi luar
    mockTx1.EXPECT().
        Commit(gomock.Any()).
        Return(nil)

    // Execute the code under test
    // Menjalankan kode yang diuji
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
```

#### Testing Transaction Rollback with Savepoints / Menguji Rollback Transaksi dengan Savepoint

```go
func TestTransactionRollbackWithSavepoints(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPool := zpgxpool.NewMockPgxPool(ctrl)
    mockTx := zpgxpool.NewMockTx(ctrl)

    // Expect Begin to be called and return a mock transaction
    // Mengharapkan Begin dipanggil dan mengembalikan transaksi mock
    mockPool.EXPECT().
        Begin(gomock.Any()).
        Return(mockTx, nil)

    // Expect Exec calls for multiple operations
    // Mengharapkan panggilan Exec untuk beberapa operasi
    mockTx.EXPECT().
        Exec(gomock.Any(), "INSERT INTO users (name) VALUES ($1)", "User 1").
        Return(pgconn.NewCommandTag("INSERT 0 1"), nil)

    // Second operation fails
    // Operasi kedua gagal
    expectedError := errors.New("constraint violation")
    mockTx.EXPECT().
        Exec(gomock.Any(), "INSERT INTO users (name) VALUES ($1)", "User 2").
        Return(pgconn.NewCommandTag(""), expectedError)

    // Expect Rollback to be called
    // Mengharapkan Rollback dipanggil
    mockTx.EXPECT().
        Rollback(gomock.Any()).
        Return(nil)

    // Execute the code under test
    // Menjalankan kode yang diuji
    tx, err := mockPool.Begin(context.Background())
    if err != nil {
        t.Fatalf("Failed to begin transaction: %v", err)
    }
    defer tx.Rollback(context.Background())

    // First operation succeeds
    // Operasi pertama berhasil
    _, err = tx.Exec(context.Background(), "INSERT INTO users (name) VALUES ($1)", "User 1")
    if err != nil {
        t.Fatalf("First operation failed: %v", err)
    }

    // Second operation fails
    // Operasi kedua gagal
    _, err = tx.Exec(context.Background(), "INSERT INTO users (name) VALUES ($1)", "User 2")
    if err == nil {
        t.Fatal("Expected second operation to fail, but it succeeded")
    }

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, got %v", expectedError, err)
    }
}
```

### 🎭 Advanced Mocking Scenarios / Skenario Mocking Lanjutan

#### Testing Batch Operations / Menguji Operasi Batch

```go
func TestBatchOperations(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPool := zpgxpool.NewMockPgxPool(ctrl)
    mockBatchResults := NewMockBatchResults(ctrl) // Assuming you have a mock for BatchResults

    // Create a batch
    // Membuat batch
    batch := &pgx.Batch{}
    batch.Queue("INSERT INTO users (name) VALUES ($1)", "User 1")
    batch.Queue("INSERT INTO users (name) VALUES ($1)", "User 2")

    // Expect SendBatch to be called and return mock results
    // Mengharapkan SendBatch dipanggil dan mengembalikan hasil mock
    mockPool.EXPECT().
        SendBatch(gomock.Any(), batch).
        Return(mockBatchResults)

    // Expect Exec to be called on the batch results for each query
    // Mengharapkan Exec dipanggil pada hasil batch untuk setiap query
    mockBatchResults.EXPECT().
        Exec().
        Return(pgconn.NewCommandTag("INSERT 0 1"), nil).
        Times(2)

    // Expect Close to be called on the batch results
    // Mengharapkan Close dipanggil pada hasil batch
    mockBatchResults.EXPECT().
        Close().
        Return(nil)

    // Execute the code under test
    // Menjalankan kode yang diuji
    batchResults := mockPool.SendBatch(context.Background(), batch)
    defer batchResults.Close()

    for i := 0; i < 2; i++ {
        _, err := batchResults.Exec()
        if err != nil {
            t.Fatalf("Batch operation %d failed: %v", i, err)
        }
    }
}
```

#### Testing Connection Pool Behavior / Menguji Perilaku Connection Pool

```go
func TestConnectionPoolBehavior(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPool := zpgxpool.NewMockPgxPool(ctrl)
    mockConn := &pgxpool.Conn{} // Using a real connection object for this test

    // Expect Acquire to be called and return a connection
    // Mengharapkan Acquire dipanggil dan mengembalikan koneksi
    mockPool.EXPECT().
        Acquire(gomock.Any()).
        Return(mockConn, nil)

    // Expect AcquireFunc to be called with a function
    // Mengharapkan AcquireFunc dipanggil dengan fungsi
    mockPool.EXPECT().
        AcquireFunc(gomock.Any(), gomock.Any()).
        DoAndReturn(func(ctx context.Context, f func(*pgxpool.Conn) error) error {
            return f(mockConn)
        })

    // Expect Config to be called and return a config
    // Mengharapkan Config dipanggil dan mengembalikan konfigurasi
    config := &pgxpool.Config{}
    mockPool.EXPECT().
        Config().
        Return(config)

    // Expect Stat to be called and return statistics
    // Mengharapkan Stat dipanggil dan mengembalikan statistik
    stat := &pgxpool.Stat{}
    mockPool.EXPECT().
        Stat().
        Return(stat)

    // Execute the code under test
    // Menjalankan kode yang diuji
    conn, err := mockPool.Acquire(context.Background())
    if err != nil {
        t.Fatalf("Failed to acquire connection: %v", err)
    }
    conn.Release()

    // Test AcquireFunc
    // Menguji AcquireFunc
    err = mockPool.AcquireFunc(context.Background(), func(c *pgxpool.Conn) error {
        // Do something with the connection
        // Melakukan sesuatu dengan koneksi
        return nil
    })
    if err != nil {
        t.Fatalf("AcquireFunc failed: %v", err)
    }

    // Test Config
    // Menguji Config
    cfg := mockPool.Config()
    if cfg == nil {
        t.Error("Expected config, got nil")
    }

    // Test Stat
    // Menguji Stat
    stats := mockPool.Stat()
    if stats == nil {
        t.Error("Expected stats, got nil")
    }
}
```

### 📋 Best Practices for Testing / Praktik Terbaik untuk Testing

1. **Always clean up resources**: Use `defer` statements to ensure mocks and controllers are properly cleaned up.

   **Selalu bersihkan sumber daya**: Gunakan pernyataan `defer` untuk memastikan mock dan controller dibersihkan dengan benar.

2. **Test both success and failure paths**: Ensure your tests cover both positive and negative scenarios.

   **Uji jalur sukses dan kegagalan**: Pastikan test Anda mencakup skenario positif dan negatif.

3. **Use descriptive test names**: Test names should clearly indicate what is being tested and what the expected outcome is.

   **Gunakan nama test yang deskriptif**: Nama test harus dengan jelas menunjukkan apa yang diuji dan apa hasil yang diharapkan.

4. **Mock only what you need**: Don't over-mock. Only mock the dependencies that are directly relevant to the code being tested.

   **Mock hanya apa yang Anda butuhkan**: Jangan terlalu banyak mocking. Hanya mock dependensi yang langsung relevan dengan kode yang diuji.

5. **Verify interactions**: Use mock expectations to verify that the correct methods are called with the correct parameters.

   **Verifikasi interaksi**: Gunakan ekspektasi mock untuk memverifikasi bahwa method yang benar dipanggil dengan parameter yang benar.

6. **Test edge cases**: Don't forget to test boundary conditions, empty results, NULL values, and error conditions.

   **Uji kasus tepi**: Jangan lupa menguji kondisi batas, hasil kosong, nilai NULL, dan kondisi error.

7. **Keep tests independent**: Each test should be able to run independently of others.

   **Jaga test tetap independen**: Setiap test harus dapat dijalankan secara independen dari yang lain.

8. **Use table-driven tests for similar scenarios**: When testing multiple similar cases, use table-driven tests to reduce code duplication.

   **Gunakan test berbasis tabel untuk skenario serupa**: Saat menguji beberapa kasus serupa, gunakan test berbasis tabel untuk mengurangi duplikasi kode.