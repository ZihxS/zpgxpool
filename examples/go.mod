module examples

go 1.24

require (
	github.com/golang/mock v1.6.0
	github.com/jackc/pgx/v5 v5.7.5
	github.com/zihxs/zpgxpool v0.0.0
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/text v0.26.0 // indirect
)

replace github.com/zihxs/zpgxpool => ../
