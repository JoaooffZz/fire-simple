module main

go 1.24.7

replace db/conn => ../db/conn

replace db/querys => ../db/querys

replace firesimple/engine => ../src/engine

replace firesimple/services => ../src/services

replace firesimple/config => ../src/config

replace server => ../server

require (
	db/conn v0.0.0-00010101000000-000000000000
	firesimple/config v0.0.0-00010101000000-000000000000
	firesimple/engine v0.0.0-00010101000000-000000000000
	server v0.0.0-00010101000000-000000000000
)

require (
	db/querys v0.0.0-00010101000000-000000000000 // indirect
	firesimple/services v0.0.0-00010101000000-000000000000 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgtype v1.14.0 // indirect
	github.com/jackc/pgx/v4 v4.18.3 // indirect
	golang.org/x/crypto v0.20.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
